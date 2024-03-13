package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := Cli(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func Cli(args []string) error {
	if len(args) == 0 {
		return errors.New("Write the filename!!!")
	}

	fmt.Println(args)
	return nil
}
