package shell_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/os/shell"
	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestExec(t *testing.T) {
	tt := test.WrapT(t)
	defer tt.Defer()
	cmd := shell.ExecString([]string{"nu", "-c", "echo \"Hello, World!\""})
	test.Assert(tt, cmd.IsOk())
	test.Assert(tt, cmd.Unwrap().Stdout == "Hello, World!\n")
}
