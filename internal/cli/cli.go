package cli

import (
	"errors"
	"fmt"
	"os"
)

func Cli(args []string) error {
	if len(args) == 0 {
		return errors.New("Write the filename!!!")
	}

	filepath := args[0]
	bs, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fmt.Println(err)
	fmt.Println(bs)
	return nil
}
