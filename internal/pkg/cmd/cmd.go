package cmd

import (
	"errors"
	"os"
	"strings"
)

// Name returns the name string comprised from the os.Args
func Name() (name string, err error) {
	name = strings.TrimSpace(strings.Join(os.Args, " "))
	if name == "" {
		err = errors.New("No game name provided")
	}
	return
}
