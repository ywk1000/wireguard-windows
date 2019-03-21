/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package service

import (
	"encoding/gob"
	"errors"
	"net/rpc"
	"os"
	"time"

	"golang.zx2c4.com/wireguard/windows/conf"
)

type Tunnel struct {
	Name string
}

type TunnelState int

const (
	TunnelUnknown TunnelState = iota
	TunnelStarted
	TunnelStopped
	TunnelStarting
	TunnelStopping
)

type NotificationType int

const (
	TunnelChangeNotificationType NotificationType = iota
	TunnelsChangeNotificationType
)

var rpcClient *rpc.Client

type TunnelChangeCallback struct {
	cb func(tunnel *Tunnel, state TunnelState, err error)
}

var tunnelChangeCallbacks = make(map[*TunnelChangeCallback]bool)

type TunnelsChangeCallback struct {
	cb func()
}

var tunnelsChangeCallbacks = make(map[*TunnelsChangeCallback]bool)

type mockEvent struct {
	Type   NotificationType
	tunnel *mockTunnel
	error  error
}

var mockEvents = make(chan mockEvent, 1024)

func InitializeIPCClient(reader *os.File, writer *os.File, events *os.File) {
	go func() {
		for event := range mockEvents {
			switch event.Type {
			case TunnelChangeNotificationType:
				for cb := range tunnelChangeCallbacks {
					cb.cb(&Tunnel{event.tunnel.name}, event.tunnel.state, event.error)
				}
			case TunnelsChangeNotificationType:
				for cb := range tunnelsChangeCallbacks {
					cb.cb()
				}
			}
		}
	}()
	return

	rpcClient = rpc.NewClient(&pipeRWC{reader, writer})
	go func() {
		decoder := gob.NewDecoder(events)
		for {
			var notificationType NotificationType
			err := decoder.Decode(&notificationType)
			if err != nil {
				return
			}
			switch notificationType {
			case TunnelChangeNotificationType:
				var tunnel string
				err := decoder.Decode(&tunnel)
				if err != nil || len(tunnel) == 0 {
					continue
				}
				var state TunnelState
				err = decoder.Decode(&state)
				if err != nil {
					continue
				}
				var errStr string
				err = decoder.Decode(&errStr)
				if err != nil {
					continue
				}
				var retErr error
				if len(errStr) > 0 {
					retErr = errors.New(errStr)
				}
				if state == TunnelUnknown {
					continue
				}
				t := &Tunnel{tunnel}
				for cb := range tunnelChangeCallbacks {
					cb.cb(t, state, retErr)
				}
			case TunnelsChangeNotificationType:
				for cb := range tunnelsChangeCallbacks {
					cb.cb()
				}
			}
		}
	}()
}

func mustFromWgQuick(raw, name string) *conf.Config {
	c, err := conf.FromWgQuick(raw, name)
	if err != nil {
		panic(err)
	}

	return c
}

func newMockTunnel(raw, name string) *mockTunnel {
	return &mockTunnel{
		name:          name,
		storedConfig:  mustFromWgQuick(raw, name),
		runtimeConfig: mustFromWgQuick(raw, name),
		state:         TunnelStopped,
	}
}

type mockTunnel struct {
	name          string
	storedConfig  *conf.Config
	runtimeConfig *conf.Config
	state         TunnelState
}

var (
	tunnels = map[string]*mockTunnel{
		"test": newMockTunnel(`[Interface]
PrivateKey = yMQR1/vVL6BYj+Giq5vLKX27GiE0F5C0KlTrIpDMuFs=
Address = 10.0.0.0/24
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = iUm/UxiVOBxfidu6F2VkIn3YvPb6I+tWzzJrQaCYBGc=
Endpoint = fake.endpoint.com:10000
AllowedIPs = 0.0.0.0/0
`, "test"),
		"test2": newMockTunnel(`[Interface]
PrivateKey = QOnJEK3XyAMVtog519Gi3I91mjbVX3o3w6GKX/CdrWE=
Address = 10.0.1.0/24
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = aZY4oX7rMosln4mIrO/lUH8+LV+5k4JDMiSiN1ftZTQ=
Endpoint = fake.target.com:10001
AllowedIPs = 0.0.0.0/0
`, "test2"),
		"test3": newMockTunnel(`[Interface]
PrivateKey = 2AgaEpf/PFFCoRaA/w+B3lzjh2k86ozwJgQKfe7gAW4=
Address = 10.0.2.0/24
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = gThUZ7eV7iyG25Yb9P7B0EXrnSnA5c8D4/Hx4F9JGgY=
Endpoint = fake.endpoint.com:10002
AllowedIPs = 0.0.0.0/0
`, "test3"),
		"test4": newMockTunnel(`[Interface]
PrivateKey = eCj1ppYn3GaNyUsjey1aKqsef1U5fwX9nCCke0ZIiFM=
Address = 10.0.3.0/24
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = VslzyrCemDaY1APDGKk20NyXBWCLgRmCaWve1BEJ6HY=
Endpoint = fake.endpoint.com:10003
AllowedIPs = 0.0.0.0/0
`, "test4"),
	}
)

func (t *Tunnel) StoredConfig() (c conf.Config, err error) {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return conf.Config{}, errors.New("unknown config")
	}
	return *(tunnel.storedConfig), nil

	err = rpcClient.Call("ManagerService.StoredConfig", t.Name, &c)
	return
}

func (t *Tunnel) RuntimeConfig() (c conf.Config, err error) {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return conf.Config{}, errors.New("unknown config")
	}
	return *(tunnel.runtimeConfig), nil

	err = rpcClient.Call("ManagerService.RuntimeConfig", t.Name, &c)
	return
}

func (t *Tunnel) Start() error {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return errors.New("unknown config")
	}

	tunnel.state = TunnelStarting
	mockEvents <- mockEvent{TunnelChangeNotificationType, tunnel, nil}
	time.AfterFunc(3*time.Second, func() {
		tunnel.state = TunnelStarted
		mockEvents <- mockEvent{TunnelChangeNotificationType, tunnel, nil}
	})
	return nil

	return rpcClient.Call("ManagerService.Start", t.Name, nil)
}

func (t *Tunnel) Stop() error {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return errors.New("unknown config")
	}
	tunnel.state = TunnelStopping
	mockEvents <- mockEvent{TunnelChangeNotificationType, tunnel, nil}
	time.AfterFunc(3*time.Second, func() {
		tunnel.state = TunnelStopped
		mockEvents <- mockEvent{TunnelChangeNotificationType, tunnel, nil}
	})
	return nil

	return rpcClient.Call("ManagerService.Stop", t.Name, nil)
}

func (t *Tunnel) WaitForStop() error {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return errors.New("unknown config")
	}
	for {
		if tunnel.state == TunnelStopped {
			break
		}
		time.Sleep(time.Second)
	}
	return nil

	return rpcClient.Call("ManagerService.WaitForStop", t.Name, nil)
}

func (t *Tunnel) Delete() error {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return errors.New("unknown config")
	}
	// is stopping enough? probably have to actually remove the tunnel
	tunnel.state = TunnelStopped
	mockEvents <- mockEvent{TunnelChangeNotificationType, tunnel, nil}
	return nil

	return rpcClient.Call("ManagerService.Delete", t.Name, nil)
}

func (t *Tunnel) State() (TunnelState, error) {
	tunnel, ok := tunnels[t.Name]
	if !ok {
		return TunnelUnknown, errors.New("unknown config")
	}
	return tunnel.state, nil

	var state TunnelState
	return state, rpcClient.Call("ManagerService.State", t.Name, &state)
}

func IPCClientNewTunnel(conf *conf.Config) (Tunnel, error) {
	tunnels[conf.Name] = newMockTunnel(conf.ToWgQuick(), conf.Name)
	// Send also the TunnelChangeNotificationType ?
	mockEvents <- mockEvent{TunnelsChangeNotificationType, nil, nil}
	return Tunnel{conf.Name}, nil

	var tunnel Tunnel
	return tunnel, rpcClient.Call("ManagerService.Create", *conf, &tunnel)
}

func IPCClientTunnels() ([]Tunnel, error) {
	var mockTunnels []Tunnel
	for name, _ := range tunnels {
		mockTunnels = append(mockTunnels, Tunnel{name})
	}
	return mockTunnels, nil

	var tunnels []Tunnel
	return tunnels, rpcClient.Call("ManagerService.Tunnels", uintptr(0), &tunnels)
}

func IPCClientQuit(stopTunnelsOnQuit bool) (bool, error) {
	return true, nil

	var alreadyQuit bool
	return alreadyQuit, rpcClient.Call("ManagerService.Quit", stopTunnelsOnQuit, &alreadyQuit)
}

func IPCClientLogFilePath() (string, error) {
	return "TODO.bin", nil

	var path string
	return path, rpcClient.Call("ManagerService.LogFilePath", uintptr(0), &path)
}

func IPCClientRegisterTunnelChange(cb func(tunnel *Tunnel, state TunnelState, err error)) *TunnelChangeCallback {
	s := &TunnelChangeCallback{cb}
	tunnelChangeCallbacks[s] = true
	return s
}
func (cb *TunnelChangeCallback) Unregister() {
	delete(tunnelChangeCallbacks, cb)
}
func IPCClientRegisterTunnelsChange(cb func()) *TunnelsChangeCallback {
	s := &TunnelsChangeCallback{cb}
	tunnelsChangeCallbacks[s] = true
	return s
}
func (cb *TunnelsChangeCallback) Unregister() {
	delete(tunnelsChangeCallbacks, cb)
}
