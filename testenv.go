// Package testenv is used to test reading from environment variables using the
// local environment. Simply backs up any variables the test intends to use and
// allows them to be restored afterwards. Also ensures nothing outside the
// intended test values are set.
//
// *NOTE: This is intended for use in tests only and will panic upon error.*
package testenv

import (
	"fmt"
	"os"
)

// Environment is the key value pairs to be used in place of `os` when dealing
// with environment variables.
type Environment struct {
	backup map[string]string
}

// NewEnvironment backs up environment variables used by the test and returns
// the environment to safely set test variables against.
//
// testKeys are the keys the test uses and are backed up.
//
// It is the user's responsibility to RestoreOriginalVars after the test.
func NewEnvironment(testKeys []string) Environment {
	vars := make(map[string]string)
	for _, key := range testKeys {
		vars[key] = os.Getenv(key)
	}
	return Environment{backup: vars}
}

// RestoreOriginalVars restores the local environment returning all variables
// set at the time of NewEnvironment.
func (e Environment) RestoreOriginalVars() {
	for key, value := range e.backup {
		err := os.Setenv(key, value)
		if err != nil {
			panic(fmt.Sprintf("Failed to restore env %s", key))
		}
	}
}

// SetVar safely sets a local environment variable that can be restored later.
func (e Environment) SetVar(key, value string) {
	// Check that we've backed up the variable before overwriting.
	if _, ok := e.backup[key]; !ok {
		panic(fmt.Sprintf("Haven't backed up env variable %s", key))
	}
	// Set env and panic on error.
	err := os.Setenv(key, value)
	if err != nil {
		panic(fmt.Sprintf("Could not set env %s", key))
	}
}

// SetVars safely sets multiple local environment variables that can be
// restored later.
func (e Environment) SetVars(env map[string]string) {
	for k, v := range env {
		e.SetVar(k, v)
	}
}
