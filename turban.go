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

// Package turban offers an shell experience.
package turban

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/readline.v1"
)

//#################//
//### Variables ###//
//#################//

var (
	rl        *readline.Instance
	promptStr = "TURBAN » "

	// Readline Auto Completion.
	completer = readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("exit"),
	)

	printASCIIArt = func() {
		fmt.Println("____ _ _  ___  ___   __  _ _ ")
		fmt.Println(" ))  ))`) ))_) ))_) /_`) )\\`)")
		fmt.Println("((  ((_( ((`\\ ((__)(( ( ((`( ")
		fmt.Println("")
	}
)

//##############//
//### Public ###//
//##############//

// Run the shell.
func Run(withColor bool) error {
	defer Stop()

	// Disables colorized output if required.
	color.NoColor = !withColor

	// Print the ASCII art as welcome message.
	color.Set(color.FgWhite)
	printASCIIArt()

	// Start the readline loop.
	return readlineLoop()
}

// Stop the shell run loop.
func Stop() {
	// Close the readline instance if present.
	if rl != nil {
		rl.Close()
	}

	// Unset the color again.
	color.Unset()
}

// SetPrompt sets the prompt string.
func SetPrompt(p string) {
	promptStr = p
}

// SetPrintASCIIArtFunc sets the function which is called when the ASCII art
// logo should be printed.
func SetPrintASCIIArtFunc(f func()) {
	printASCIIArt = f
}

//###############//
//### Private ###//
//###############//

func readlineLoop() error {
	var err error
	promptPrint := color.New(color.FgRed, color.Bold).SprintFunc()

	// Prepare readline.
	rl, err = readline.NewEx(&readline.Config{
		Prompt:       promptPrint(promptStr),
		AutoComplete: completer,
	})
	if err != nil {
		return err
	}
	defer rl.Close()

	// Read loop.
	for {
		// Read a line.
		line, err := rl.Readline()
		if err != nil {
			// Log the error.
			if err == io.EOF || err == readline.ErrInterrupt {
				return nil
			}

			return fmt.Errorf("readline error: %v", err)
		}

		// Trim spaces.
		line = strings.TrimSpace(line)

		// Set the color.
		color.Unset()
		color.Set(color.FgWhite)

		// Exit if required.
		if line == "exit" {
			return nil
		}

		// Split the input line.
		args := strings.Fields(line)

		// Skip if empty.
		if len(args) == 0 {
			continue
		}

		// Get the command key.
		key := args[0]

		// Remove the first command key from the slice.
		args = args[1:]

		// Try to find the command in the commands map.
		cmd, ok := commands[key]
		if !ok {
			fmt.Println("error: invalid command")
			continue
		}

		if cmd.Run == nil {
			fmt.Println("error: no run function defined for this shell command")
			continue
		}

		// Run the command.
		err = cmd.Run(args)
		if err != nil {
			fmt.Printf("error: %v\n", err)

			// Print the usage if this is an invalid usage error.
			if err == ErrInvalidUsage {
				fmt.Println("Usage:", cmd.Usage)
			}
		}
	}
}
