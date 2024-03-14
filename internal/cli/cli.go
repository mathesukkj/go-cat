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

	bs := ReadFile(args)

	PrintFile(os.Stdout, bs)
}

func ReadFile(filepaths []string) []byte {
	fileContents := []byte{}

	for _, v := range filepaths {
		bs, err := os.ReadFile(v)
		if err != nil {
			errorMessage := fmt.Sprintf("gocat: %s\n", err)
			fileContents = append(fileContents, []byte(errorMessage)...)
		} else {
			fileContents = append(fileContents, bs...)
		}
	}

	return fileContents
}

func PrintFile(w io.Writer, bs []byte) {
	w.Write(bs)
}
