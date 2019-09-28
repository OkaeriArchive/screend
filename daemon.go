/*
 * screend
 * Copyright (C) 2019 OkaeriPoland
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
package main

import (
	"github.com/go-ini/ini"
)

type ScreenInfo struct {
	Name string `ini:"name"`
	User string `ini:"user"`
}

type ExecutionInfo struct {
	Command      string `ini:"command"`
	StartHook    string `ini:"start_hook"`
	RunDirectory string `ini:"run_directory"`
}

type DaemonInfo struct {
	Enabled bool `ini:"enabled"`
	Logging bool `ini:"logging"`
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
