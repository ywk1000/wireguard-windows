package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
)

func (mw *MainWindowModel) exportCurrentTunnel() error {
	if mw.model.Current == nil {
		walk.MsgBox(mw, "No tunnel selected", "You need to select a tunnel first", walk.MsgBoxIconInformation)
		return nil
	}
	dlg := new(walk.FileDialog)
	dlg.Filter = "WireGuard tunnel file (*.conf)|*.conf"
	dlg.Title = "Save to..."
	dlg.FilePath = mw.model.Current.Name + ".conf"
	if ok, err := dlg.ShowSave(mw); err != nil || !ok {
		return err
	}

	filepath := dlg.FilePath
	if path.Ext(filepath) != ".zip" {
		filepath += ".zip"
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, mw.model.Current.ToWgQuick())
	return nil
}

func (mw *MainWindowModel) exportTunnels() error {
	dlg := new(walk.FileDialog)
	dlg.Filter = "WireGuard tunnels (*.zip)|*.zip"
	dlg.Title = "Save to..."

	if ok, err := dlg.ShowSave(mw); err != nil || !ok {
		return err
	}

	filepath := dlg.FilePath
	if path.Ext(filepath) != ".zip" {
		filepath += ".zip"
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, tunnel := range mw.model.items {
		if err = writeTunnelToZip(zipWriter, tunnel); err != nil {
			return err
		}
	}
	return nil
}

func writeTunnelToZip(zipWriter *zip.Writer, tunnel *conf.Config) error {
	inner, err := zipWriter.Create(tunnel.Name + ".conf")
	if err != nil {
		return err
	}
	if _, err = io.WriteString(inner, tunnel.ToWgQuick()); err != nil {
		return err
	}
	return nil
}

func (mw *MainWindowModel) exportLog() {
	dlg := new(walk.FileDialog)
	dlg.Filter = "Log files (*.log)|*.log"
	dlg.Title = "Save to..."

	if ok, err := dlg.ShowSave(mw); err != nil {
		log.Print(err)
	} else if !ok {
		return
	}
	//todo: save log file here

}
