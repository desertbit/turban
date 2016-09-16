# Turban Shell

[![GoDoc](https://godoc.org/github.com/desertbit/turban?status.svg)](https://godoc.org/github.com/desertbit/turban)
[![Go Report Card](https://goreportcard.com/badge/github.com/desertbit/turban)](https://goreportcard.com/report/github.com/desertbit/turban)

A simple shell experience.

![Turban Preview](preview.jpg "Turban")

## Turban Application

There is a turban shell application available in the cmd directory.
It reads a TURBAN.toml file and provides a simple shell.

```
	go get -u github.com/desertbit/turban
	go install github.com/desertbit/turban/cmd/turban
	turban -p "prompt text" -d "path to directory"
```

## Library Sample

```go
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
```
