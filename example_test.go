package testenv_test

import (
	"github.com/whitecaleb/go-test-env"
)

const aTestKey = "A_TEST_KEY"
const anotherTestKey = "ANOTHER_TEST_KEY"

// An example of environment variable keys to be used in the test.
var exampleTestKeys = []string{aTestKey, anotherTestKey}

func ExampleNewEnvironment() {
	// Instantiate an environment once per test.
	env := testenv.NewEnvironment(exampleTestKeys)

	_ = env
}

func ExampleEnvironment_RestoreOriginalVars() {
	env := testenv.NewEnvironment(exampleTestKeys)

	// To ensure the original variables are restored defer the restore
	// immediately after creating the test environment.
	defer env.RestoreOriginalVars()

	// Do tests...
}

func ExampleEnvironment_SetVar() {
	env := testenv.NewEnvironment(exampleTestKeys)
	defer env.RestoreOriginalVars()

	// Set a variable to use in the test.
	env.SetVar(aTestKey, "a test value")
}

func ExampleEnvironment_SetVars() {
	env := testenv.NewEnvironment(exampleTestKeys)
	defer env.RestoreOriginalVars()

	// Create a collection of variables to use in the test.
	myTestVars := make(map[string]string)
	myTestVars[aTestKey] = "a test value"
	myTestVars[anotherTestKey] = "another test value"

	// Set a variable to use in the test.
	env.SetVars(myTestVars)
}
