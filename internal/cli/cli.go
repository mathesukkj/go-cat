package cli

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Cli(args []string) {
	if len(args) == 0 {
		log.Panic("No filename given!")
	}

	bs, err := ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	PrintFile(os.Stdout, bs)
}

func ReadFile(filepath string) ([]byte, error) {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	fmt.Println(err)
	fmt.Println(bs)
	return bs, nil
}

func PrintFile(w io.Writer, bs []byte) {
	w.Write(bs)
}
