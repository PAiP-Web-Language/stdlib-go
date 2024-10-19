package todo

import (
	"fmt"
	"strings"
)

// TodoStruct is an empty struct meant to be used as placeholder for temporary no value return
type TodoStruct struct {}

// TodoInterface is an empty interface
type TodoInterface interface {}

// TodoNotUsed is a function that takes in any number of arguments
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant for catch all to ignore errors during development
// This function is strictly a placeholder. Its not meant to be used to ignore errors for valid reason.
// If you want to ignore errors for valid reason you can use dev.*UsedInDebugging or dev.*UsedInSpecificScenarios functions
// Example of usage:
// var _ = TodoNotUsed(fmt.Sprint, test)
func TodoNotUsed(x ...any) (r TodoStruct) { return }

// TodoPanic is a function that notes that this path of code is not done yet
// This function is wrapper around panic
// You can add argument to specify why this is not implemented if needed
// This function should be used for temporary notes of not implemented code paths
// If you want to specify this code path to be meant this way use dev.NotImplemented or dev.Unreachable instead
func Todo(args ...string) {
	if len(args) == 0 {
		panic("not yet implemented")
	} else if len(args) == 1 {
		panic(fmt.Sprintf("not yet implemented: %s", args[0]))
	} else {
		panic(fmt.Sprintf("not yet implemented: %s", strings.Join(args, " ")))
	}
}

