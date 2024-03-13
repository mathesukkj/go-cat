package cli

import (
	"fmt"
	"log"
	"os"
)

func Cli(args []string) {
	if len(args) == 0 {
		log.Panic("No filename given!")
	}

	ReadFile(args[0])
}

func ReadFile(filepath string) error {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fmt.Println(err)
	fmt.Println(bs)
	return nil
}
