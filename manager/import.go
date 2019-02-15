package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
)

func (mw *MainWindowModel) importTunnels() error {
	dlg := new(walk.FileDialog)
	dlg.Filter = "WireGuard tunnel files (*.zip;*.conf)|*.zip;*.conf"
	dlg.Title = "Import tunnel file..."

	if ok, err := dlg.ShowOpen(mw.MainWindow); err != nil || !ok {
		return err
	}
	ext := filepath.Ext(dlg.FilePath)
	switch ext {
	case ".conf":
		tunnel, err := importConf(dlg.FilePath)
		if err != nil {
			return err
		}
		mw.model.items = append(mw.model.items, tunnel)
	case ".zip":
		tunnels, err := importZip(dlg.FilePath)
		if err != nil {
			return err
		}
		for _, tunnel := range tunnels {
			mw.model.items = append(mw.model.items, tunnel)
		}
	default:
		return fmt.Errorf("Unrecognized file %v", dlg.FilePath)
	}
	mw.model.PublishItemsReset()
	return nil
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func importConf(path string) (*conf.Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	tunnel, err := conf.FromWgQuick(string(b), filenameWithoutExtension(filepath.Base(path)))
	if err != nil {
		return nil, err
	}
	return tunnel, nil
}

func importZip(path string) ([]*conf.Config, error) {
	z, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	tunnels := make([]*conf.Config, 0)
	for _, f := range z.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		rc.Close()

		tunnel, err := conf.FromWgQuick(string(b), filenameWithoutExtension(f.Name))
		if err != nil {
			return nil, err
		}
		tunnels = append(tunnels, tunnel)
	}
	return tunnels, nil
}
