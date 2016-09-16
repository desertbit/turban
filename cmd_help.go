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
	"fmt"
	"sort"
)

//#################//
//### Variables ###//
//#################//

var (
	helpHeader = ""
)

//##############//
//### Public ###//
//##############//

// SetHelpHeader sets the header displayed in the help.
func SetHelpHeader(h string) {
	helpHeader = h
}

//###############//
//### Private ###//
//###############//

func init() {
	AddCommand("help", &Command{
		Help:  "Print this help text.",
		Usage: "help",
		Run: func(args []string) error {
			// Check if any arguments are passed.
			if len(args) > 0 {
				return ErrInvalidUsage
			}

			// Get all command map keys.
			i := 0
			keys := make([]string, len(commands))
			for k := range commands {
				keys[i] = k
				i++
			}

			// Sort the keys slice.
			sort.Strings(keys)

			// Header.
			printASCIIArt()

			if len(helpHeader) > 0 {
				fmt.Printf("\n%s\n", helpHeader)
			}

			fmt.Print("\nAvailable commands:\n\n")

			// Print all available commands with a description.
			for _, key := range keys {
				printc("  "+key, commands[key].Help)
			}

			// Flush the output.
			flush()

			// Print a new empty line.
			fmt.Println()

			return nil
		},
	})
}
