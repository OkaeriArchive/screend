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
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {

	files, err := filepath.Glob("daemons/*.ini")
	if err != nil {
		log.Fatal("Failed to find daemons: ", err)
		return
	}

	var daemons []DaemonConfig
	for _, cfgPath := range files {

		cfg, err := loadDaemon(cfgPath)
		if err != nil {
			log.Print("Failed load ", cfgPath, err)
			continue
		}

		daemons = append(daemons, *cfg)
	}

	firstRun := true

	for {
		for _, daemon := range daemons {

			if !daemon.Enabled {
				continue
			}

			command := daemon.Command
			runDirectory := daemon.RunDirectory
			name := daemon.Name
			userName := daemon.User
			logging := daemon.Logging

			screen, err := runScreen(userName, name, runDirectory, command, logging)
			if err != nil {
				if err.Error() == "SCREEN_ALREADY_EXISTS" {
					if firstRun {
						log.Print("(", name, " on ", userName, ") WARNING: skipped, screen already exists")
					}
				} else {
					log.Print("(", name, " on ", userName, ") FAILED: ", command, " [", err, "]")
				}
				continue
			}

			log.Print("("+strconv.Itoa(screen.id)+", "+screen.name, " on ", userName, ") STARTED: "+command)

			startHook := daemon.StartHook
			if len(startHook) > 0 {

				userInfo, err := getUserInfoByName(userName)
				if err != nil {
					log.Print("(", name, " on ", userName, ") START HOOK FAILED: cannot get user info")
				}

				output, err := executeCommand(userInfo, false, runDirectory, strings.Split(startHook, " ")...)
				if err != nil {
					log.Print("(", name, " on ", userName, ") START HOOK FAILED: ", output, " [", err, "]")
				}
			}
		}

		firstRun = false
		time.Sleep(5 * time.Second)
	}
}
