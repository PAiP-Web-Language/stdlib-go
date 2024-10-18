package dev

// EnsureTypesMatch checks if provided argumet is of type T
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it in init function of modules to have your LSP suggest
// if your type matches some interface or still needs some methods
func EnsureTypesMatch[T any](a T) {
}

// VoidStruct is an empty struct
type VoidStruct struct {}

// FunctionsUsedInDebugging is a function that takes in any number of functions
// That are used in debugging
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = FunctionsUsedInDebugging(fmt.Sprint)
func FunctionsUsedInDebugging(x ...any) (r VoidStruct) { return }

// FunctionsUsedInSpecificScenarios is a function that takes in any number of functions
// That are used in some specific scenarios that LSP can not pickup. For example preprocessor based code
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = FunctionsUsedInSpecificScenarios(fmt.Sprint)
func FunctionsUsedInSpecificScenarios(x ...any) (r VoidStruct) { return }

// PackagesUsedInDebugging is a function that takes in any number of packages
// (more specifically any exported function from those packages just to note package, because go would not allow you to use package as arg)
// That are used in debugging
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = PackagesUsedInDebugging(fmt.Sprint)
func PackagesUsedInDebugging(x ...any) (r VoidStruct) { return }

// PackagesUsedInSpecificScenarios is a function that takes in any number of packages
// (more specifically any exported function from those packages just to note package, because go would not allow you to use package as arg)
// That are used in some specific scenarios that LSP can not pickup. For example preprocessor based code
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = PackagesUsedInSpecificScenarios(fmt.Sprint)
func PackagesUsedInSpecificScenarios(x ...any) (r VoidStruct) { return }

// VariablesUsedInDebugging is a function that takes in any number of variables
// That are used in debugging
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = VariablesUsedInDebugging(test)
func VariablesUsedInDebugging(x ...any) (r VoidStruct) { return }

// VariablesUsedInSpecificScenarios is a function that takes in any number of variables
// That are used in some specific scenarios that LSP can not pickup. For example preprocessor based code
// This function is NOOP so it should be pretty fast
// Idea for that function is to use it to ignore errors for imported and not used / declared and not used
// This function is meant to note that specifed arguments are being used in debugging
// This function is not meant for catch all to ignore errors
// If you need catch all temporarly you can use todo.TodoNotUsed function
// Example of usage:
// var _ = VariablesUsedInSpecificScenarios(test)
func VariablesUsedInSpecificScenarios(x ...any) (r VoidStruct) { return }

// NotImplemented is a function that notes that this path of code is not imlemented
// This function is wrapper around panic
// You can add argument to specify why this is not implemented if needed
func NotImplemented(args ...string) {
	if len(args) == 0 {
		panic("not implemented")
	} else if len(args) == 1 {
		panic(fmt.Sprintf("not implemented: %s", args[0]))
	} else {
		panic(fmt.Sprintf("not implemented: %s", strings.Join(args, " ")))
	}
}

// Unreachable is a function that notes that this path of code is meant to be unreachable
// This function is wrapper around panic
// You can add argument to specify why this is not reachable if needed
func Unreachable(args ...string) {
	if len(args) == 0 {
		panic("internal error: entered unreachable code")
	} else if len(args) == 1 {
		panic(fmt.Sprintf("internal error: entered unreachable code: %s", args[0]))
	} else {
		panic(fmt.Sprintf("internal error: entered unreachable code: %s", strings.Join(args, " ")))
	}
}

