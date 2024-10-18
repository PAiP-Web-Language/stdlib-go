package dev

import (
	"math"
	"os"
	"strings"

	"github.com/gookit/goutil/dump"
)

// defaultSkip is the default number of stack frames to skip
// This constant is meant to be directly copied from the source code line below
// ~/go/pkg/mod/github.com/gookit/goutil@<ver>/dump/dump.go:24
const defaultSkip = 3

// This constant is meant to be tested manually and changed so it shows correct information
const DevDumpSkip = defaultSkip + 0

// Dump is a function that takes in any number of arguments and prints them to the console
// This function is meant to be used for debugging purposes
// This function is printing to stderr
func Dump(vs ...any) {
	dump.NewWithOptions(
		dump.WithCallerSkip(DevDumpSkip),
	).Fprint(os.Stderr, vs...)
}

// DumpDepth is a function that takes in any number of arguments and prints them to the console with specified depth
// This function is meant to be used for debugging purposes
// This function is printing to stderr
func DumpDepth(depth int, vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = depth
		},
	).Fprint(os.Stderr, vs...)
}

// DumpAll is a function that takes in any number of arguments and prints them to the console with unlimited depth
// This function is meant to be used for debugging purposes
// This function is printing to stderr
// Depth is techically unlimited, because it is set to math.MaxInt
func DumpAll(vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = math.MaxInt
		},
	).Fprint(os.Stderr, vs...)
}

// StdoutDump is a function that takes in any number of arguments and prints them to the console
// This function is meant to be used for debugging purposes
// This function is printing to stdout
func StdoutDump(vs ...any) {
	dump.NewWithOptions(
		dump.WithCallerSkip(DevDumpSkip),
	).Print(vs...)
}

// StdoutDumpDepth is a function that takes in any number of arguments and prints them to the console with specified depth
// This function is meant to be used for debugging purposes
// This function is printing to stdout
func StdoutDumpDepth(depth int, vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = depth
		},
	).Print(vs...)
}

// StdoutDumpAll is a function that takes in any number of arguments and prints them to the console with unlimited depth
// This function is meant to be used for debugging purposes
// This function is printing to stdout
// Depth is techically unlimited, because it is set to math.MaxInt
func StdoutDumpAll(vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = math.MaxInt
		},
	).Print(vs...)
}

// StringDump is a function that takes in any number of arguments and returns string of them
// This function is meant to be used for debugging purposes
// This function is printing to string and returning it
func StringDump(vs ...any) string {
	buf := new(strings.Builder)
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.NoColor = true
			opts.CallerSkip = DevDumpSkip
		},
	).Fprint(buf, vs...)
	return buf.String()
}

// StringDumpDepth is a function that takes in any number of arguments and returns string of them with specified depth
// This function is meant to be used for debugging purposes
// This function is printing to string and returning it
func SdumpDepth(depth int, vs ...any) string {
	buf := new(strings.Builder)
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.NoColor = true
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = depth
		},
	).Fprint(buf, vs...)
	return buf.String()
}

// StringDumpAll is a function that takes in any number of arguments and returns string of them with unlimited depth
// This function is meant to be used for debugging purposes
// This function is printing to string and returning it
// Depth is techically unlimited, because it is set to math.MaxInt
func SdumpAll(vs ...any) string {
	buf := new(strings.Builder)
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.NoColor = true
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = math.MaxInt
		},
	).Fprint(buf, vs...)
	return buf.String()
}

// SingleStringDump is a function that takes in one argument and returns string of them with unlimited depth
// This function is meant to be used for debugging purposes
// This function is printing to string and returning it
// Depth is techically unlimited, because it is set to math.MaxInt
// This function is meant to be used for dumping single value to string for more processing like debug slog
func SingleStringDump(vs any) string {
	buf := new(strings.Builder)
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.NoColor = true
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = math.MaxInt
			opts.ShowFlag = dump.Fnopos
		},
	).Fprint(buf, vs)
	return buf.String()
}
