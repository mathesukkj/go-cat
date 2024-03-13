package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("check if filename exists", func(t *testing.T) {
		args := []string{}
		err := Cli(args)

		if err == nil {
			t.Errorf("Expected to get an error, but didn't get one")
		}
	})
}
