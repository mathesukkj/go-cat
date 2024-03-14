package main

import (
	"os"

	"mathesukkj/gocat/internal/cli"
)

func main() {
	cli.Cli(os.Args[1:])
}
