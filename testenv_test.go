package testenv

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestRestoreOriginalVars(t *testing.T) {
	// create a unique key that will not be set in the local environment
	unsetUniqueKey := uuid.New().String()
	checkGetEnv(t, unsetUniqueKey, "")
	// create a unique key that will be set in the local environment
	setUniqueKey := uuid.New().String()
	originalValue := "myOriginalTestValue"
	os.Setenv(setUniqueKey, originalValue)
	checkGetEnv(t, setUniqueKey, originalValue)

	env := NewEnvironment([]string{setUniqueKey, unsetUniqueKey})

	// some test values
	val1 := "test value 1"
	val2 := "test value 2"

	// create test variables
	vars := make(map[string]string)
	vars[unsetUniqueKey] = val1
	vars[setUniqueKey] = val2

	// set the local env with test vars
	env.SetVars(vars)
	checkGetEnv(t, unsetUniqueKey, val1)
	checkGetEnv(t, setUniqueKey, val2)

	// restore original local env
	env.RestoreOriginalVars()
	checkGetEnv(t, unsetUniqueKey, "")
	checkGetEnv(t, setUniqueKey, originalValue)
}

func checkGetEnv(t *testing.T, key, value string) {
	v := os.Getenv(key)
	if v != value {
		t.Errorf("%s expected var %s to equal %s, got %s", t.Name(), key, value, v)
	}
}

func TestSetVarShouldPanic(t *testing.T) {
	// check for panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should not have set variable.")
		}
	}()

	// create env with no backups
	env := Environment{}

	// attempt to set var
	env.SetVar("A_KEY", "any value")
}
