package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mathesukkj/gocat/internal/cli"
)

type SpyWriter struct {
	called bool
}

func (s *SpyWriter) Write(p []byte) (int, error) {
	fmt.Println(p)
	s.called = true
	return 0, nil
}

func TestCli(t *testing.T) {
	t.Run("check if filename exists", func(t *testing.T) {
		defer func() { _ = recover() }()

		args := []string{}
		cli.Cli(args)

		t.Errorf("code didnt panic but it should")
	})
}

func TestReadFile(t *testing.T) {
	t.Run("check if file is correctly read", func(t *testing.T) {
		file := createTempFile(t, "arquivo de teste")
		bs := cli.ReadFile([]string{file.Name()})

		if string(bs) != "arquivo de teste\n" {
			t.Error("file not being read correctly")
		}
	})
	t.Run("read more than one file", func(t *testing.T) {
		file1 := createTempFile(t, "test1")
		file2 := createTempFile(t, "test2")
		files := []string{file1.Name(), file2.Name()}

		bs := cli.ReadFile(files)

		if string(bs) != "test1\ntest2\n" {
			t.Error("file not being read correctly")
		}
	})
}

func TestPrintFile(t *testing.T) {
	t.Run("check if file is correctly logged", func(t *testing.T) {
		content := []byte("string")

		spy := &SpyWriter{}
		cli.PrintFile(spy, content)

		if !spy.called {
			t.Errorf("writer wasn't called when it should have been")
		}
	})
}

func createTempFile(t testing.TB, content string) *os.File {
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
