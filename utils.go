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
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

//#################//
//### Variables ###//
//#################//

var (
	// Tab writer
	tabWriterStdout *tabwriter.Writer
)

//##############//
//### Public ###//
//##############//

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

//###############//
//### Private ###//
//###############//

func init() {
	// Create a new tab writer to format the output.
	tabWriterStdout = new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	tabWriterStdout.Init(os.Stdout, 2, 8, 2, '\t', 0)
}

func print(a ...interface{}) {
	fmt.Fprint(tabWriterStdout, a...)
}

func println(a ...interface{}) {
	fmt.Fprintln(tabWriterStdout, a...)
}

func printf(format string, a ...interface{}) {
	fmt.Fprintf(tabWriterStdout, format, a...)
}

// printc prints all the variadic arguments as columns.
func printc(a ...interface{}) {
	// Create the format string
	var f string
	for range a {
		f += "%v\t"
	}
	f += "\n"

	// Print the columns.
	printf(f, a...)
}

func flush() {
	tabWriterStdout.Flush()
}
