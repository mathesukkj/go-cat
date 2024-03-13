package main

import (
	"os"

	"mathesukkj/go-cat/internal/cli"
)

func main() {
	cli.Cli(os.Args[1:])

}
