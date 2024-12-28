package test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/os/shell"
)

type CliTestStringCmdResult shell.StringCmdResult

func CliRun(t testing.TB, cmd []string) CliTestStringCmdResult {
	t.Helper()
	cmdr := shell.ExecString(cmd)
	Assert(t, cmdr.IsOk())
	return CliTestStringCmdResult(cmdr.Unwrap())
}

func (r CliTestStringCmdResult) ExpectStdout(t testing.TB, expected string) CliTestStringCmdResult {
	t.Helper()
	Assert(t, r.Stdout == expected)
	return r
}

func (r CliTestStringCmdResult) ExpectStderr(t testing.TB, expected string) CliTestStringCmdResult {
	t.Helper()
	Assert(t, r.Stderr == expected)
	return r
}

func (r CliTestStringCmdResult) ExpectExitCode(t testing.TB, expected int) CliTestStringCmdResult {
	t.Helper()
	Assert(t, r.Cmd.ProcessState.ExitCode() == expected)
	return r
}

