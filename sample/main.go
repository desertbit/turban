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
	"fmt"

	"github.com/desertbit/turban"
)

func main() {
	turban.SetPrompt("TURBAN » ")
	turban.SetHelpHeader("TURBAN - A simple shell experience")
	turban.SetPrintASCIIArtFunc(printASCIIArt)

	turban.AddCommand("foo", &turban.Command{
		Help:  "Print foo help text",
		Usage: "foo [BAR]",
		Run: func(args []string) error {
			// ...
			return nil
		},
	})

	turban.Run(true)
}

func printASCIIArt() {
	fmt.Println("____ _ _  ___  ___   __  _ _ ")
	fmt.Println(" ))  ))`) ))_) ))_) /_`) )\\`)")
	fmt.Println("((  ((_( ((`\\ ((__)(( ( ((`( ")
	fmt.Println("")
}
