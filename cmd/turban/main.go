/*
 *  Turban Shell
 *  Copyright (C) 2016  Roland Singer <roland.singer[at]desertbit.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/desertbit/turban"
)

const (
	turbanFileName = "TURBAN.toml"
)

type command struct {
	Name  string
	Help  string
	Usage string
	Exec  string
}

type turbanFile struct {
	Cmd []command
}

func main() {
	// Get the current working dir.
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to obtain the current working directory: %v", err)
	}

	dirPath := flag.String("d", pwd, "set the source directory path which contains the TURBAN.toml file")
	promptStr := flag.String("p", "» ", "set the prompt string")
	helpHeaderStr := flag.String("h", "", "set the helper header string")

	flag.Parse()

	turban.SetPrompt(*promptStr)
	turban.SetHelpHeader(*helpHeaderStr)

	err = loadTurbanFile(*dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	turban.Run(true)
}

func loadTurbanFile(dirPath string) error {
	turbanFilePath := filepath.Join(dirPath, turbanFileName)

	var t turbanFile
	_, err := toml.DecodeFile(turbanFilePath, &t)
	if err != nil {
		return fmt.Errorf("failed to decode turban file: %v", err)
	}

	for _, cmd := range t.Cmd {
		if len(cmd.Name) == 0 {
			return fmt.Errorf("invalid turban file: no command name specified")
		}

		execStr := cmd.Exec

		turban.AddCommand(cmd.Name, &turban.Command{
			Help:  cmd.Help,
			Usage: cmd.Usage,
			Run: func(args []string) (err error) {
				cmdSplit := strings.Fields(execStr)
				if len(cmdSplit) < 1 {
					return turban.ErrInvalidUsage
				}

				commandName := cmdSplit[0]
				cmdSplit = cmdSplit[1:]

				// Replace all %1, %2, ... %n values with the argument values.
				maxArg := 0
				for i := 0; i < len(cmdSplit); i++ {
					p := cmdSplit[i]
					if !strings.HasPrefix(p, "%") {
						continue
					}
					p = p[1:]
					n, err := strconv.Atoi(p)
					if err != nil {
						continue
					}
					if n > maxArg {
						maxArg = n
					}
					n--
					if n < 0 || n >= len(args) {
						return turban.ErrInvalidUsage
					}
					cmdSplit[i] = args[n]
				}

				// Check if too many arguments where passed.
				if len(args) > maxArg {
					return turban.ErrInvalidUsage
				}

				// Run the command.
				cmd := exec.Command(commandName, cmdSplit...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Dir = dirPath

				err = cmd.Run()
				if err != nil {
					return fmt.Errorf("command failed: %v", err)
				}

				return nil
			},
		})
	}

	return nil
}
