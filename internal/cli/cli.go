package cli

import (
	"errors"
	"fmt"
)

func Cli(args []string) error {
	if len(args) == 0 {
		return errors.New("Write the filename!!!")
	}

	fmt.Println(args)
	return nil
}
