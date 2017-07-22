# testenv
--
    import "github.com/whitecaleb/go-test-env"

Package testenv is used to test reading from environment variables using the
local environment. Simply backs up any variables the test intends to use and
allows them to be restored afterwards. Also ensures nothing outside the intended
test values are set.

*NOTE: This is intended for use in tests only and will panic upon error.*

## Usage

#### type Environment

```go
type Environment struct {
}
```

Environment is the key value pairs to be used in place of `os` when dealing with
environment variables.

#### func  NewEnvironment

```go
func NewEnvironment(testKeys []string) Environment
```
NewEnvironment backs up environment variables used by the test and returns the
environment to safely set test variables against.

testKeys are the keys the test uses and are backed up.

It is the user's responsibility to RestoreOriginalVars after the test.

#### func (Environment) RestoreOriginalVars

```go
func (e Environment) RestoreOriginalVars()
```
RestoreOriginalVars restores the local environment returning all variables set
at the time of NewEnvironment.

#### func (Environment) SetVar

```go
func (e Environment) SetVar(key, value string)
```
SetVar safely sets a local environment variable that can be restored later.

#### func (Environment) SetVars

```go
func (e Environment) SetVars(env map[string]string)
```
SetVars safely sets multiple local environment variables that can be restored
later.
