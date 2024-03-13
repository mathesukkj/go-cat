package test

import (
	"fmt"
	"os"
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
		_, err := cli.ReadFile(filepath)

		if err == nil {
			t.Errorf("Expected to get an error, but didn't get one")
		}
	})
	t.Run("check if file is correctly read", func(t *testing.T) {
		content := "arquivo de teste"
		file := CreateTempFile(t, content)
		bs, err := cli.ReadFile(file.Name())

		if err != nil {
			fmt.Println(err)
			t.Errorf("got an error but it shouldnt have")
		}

		if string(bs) != content {
			t.Error("file not being read correctly")
		}
	})
}

func CreateTempFile(t testing.TB, content string) *os.File {
	t.Helper()

	file, err := os.CreateTemp("/tmp/", "test")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte(content)); err != nil {
		panic(err)
	}

	return file
}
