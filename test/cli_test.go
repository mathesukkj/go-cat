package test

import (
	"testing"

	"mathesukkj/go-cat/internal/cli"
)

func TestCli(t *testing.T) {
	t.Run("check if filename exists", func(t *testing.T) {
		args := []string{}
		err := cli.Cli(args)

		if err == nil {
			t.Errorf("Expected to get an error, but didn't get one")
		}
	})
	t.Run("check if file exists", func(t *testing.T) {
		args := []string{"/tmp/arquivo-naoexistentekkk"}
		err := cli.Cli(args)

		if err == nil {
			t.Errorf("Expected to get an error, but didn't get one")
		}
	})
}
