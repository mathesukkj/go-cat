package test

import (
	"testing"

	"mathesukkj/go-cat/internal/cli"
)

func TestCli(t *testing.T) {
	t.Run("check if filename exists", func(t *testing.T) {
		defer func() { _ = recover() }()

		args := []string{}
		cli.Cli(args)

		t.Errorf("code didnt panic but it should")
	})
}

func TestReadFile(t *testing.T) {
	t.Run("check if file exists", func(t *testing.T) {
		filepath := "/tmp/arquivo-naoexistentekkk"
		err := cli.ReadFile(filepath)

		if err == nil {
			t.Errorf("Expected to get an error, but didn't get one")
		}
	})

}
