package main

import (
	"github.com/go-ini/ini"
)

type ScreenInfo struct {
	Name string `ini:"name"`
}

type ExecutionInfo struct {
	Command string `ini:"command"`
}

type DaemonInfo struct {
	Enabled bool `ini:"enabled"`
}

type DaemonConfig struct {
	ScreenInfo    `ini:"Screen"`
	ExecutionInfo `ini:"Execution"`
	DaemonInfo    `ini:"Daemon"`
}

func loadDaemon(cfgPath string) (*DaemonConfig, error) {

	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:         true,
		UnescapeValueCommentSymbols: true,
	}, cfgPath)

	daemonConfig := new(DaemonConfig)
	err = cfg.MapTo(daemonConfig)
	if err != nil {
		return daemonConfig, err
	}

	return daemonConfig, nil
}
