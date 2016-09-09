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

package turban

import (
	"errors"

	"gopkg.in/readline.v1"
)

//#################//
//### Variables ###//
//#################//

var (
	// ErrInvalidUsage when returned by the Run Command function,
	// then usage is printed.
	ErrInvalidUsage = errors.New("invalid usage")
)

var (
	commands = make(map[string]*Command)
)

//###############//
//### Command ###//
//###############//

// RunFunc specifies the function which is run during a command.
type RunFunc func(args []string) error

// A Command specifies a turban shell command.
// PrefixCompleter is optional and will be set to a default if not set.
type Command struct {
	Help  string
	Usage string
	Run   RunFunc

	PrefixCompleter *readline.PrefixCompleter
}

//##############//
//### Public ###//
//##############//

// AddCommand register a new shell command.
func AddCommand(key string, cmd *Command) {
	if cmd.PrefixCompleter == nil {
		cmd.PrefixCompleter = readline.PcItem(key)
	}

	commands[key] = cmd

	// Add the completer.
	completer.Children = append(completer.Children, cmd.PrefixCompleter)
}
