package test

import (
	"runtime"
	"testing"
)


func ErrorFormat(tb testing.TB, skip int, msg string, args ...any) {
	_, file, line, _ := runtime.Caller(skip + 1)
	tb.Errorf("%s:%d %s %v", file, line, msg, args)
}
