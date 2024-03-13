package main

import (
	"os"

	"mathesukkj/go-cat/internal/cli"
)

func main() {
	err := cli.Cli(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
