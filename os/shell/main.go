package shell

import (
	"io"
	"os"
	"os/exec"
	"strings"

	. "github.com/PAiP-Web-Language/stdlib-go/base"
	"github.com/PAiP-Web-Language/stdlib-go/errors"
)

func Exec(cmdLine []string) Result[*exec.Cmd] {
	if len(cmdLine) == 0 {
		return Err[*exec.Cmd](errors.FunctionalError{})
	}

	cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		Err[*exec.Cmd](err)
	}
	cmderr := cmd.Wait()
	if cmderr != nil {
		return Err[*exec.Cmd](cmderr)
	}
	return Ok(cmd)
}

func ExecRW(cmdLine []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) Result[*exec.Cmd] {
	if len(cmdLine) == 0 {
		return Err[*exec.Cmd](errors.FunctionalError{})
	}

	cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	if err != nil {
		Err[*exec.Cmd](err)
	}
	cmderr := cmd.Wait()
	if cmderr != nil {
		return Err[*exec.Cmd](cmderr)
	}
	return Ok(cmd)
}

type StringCmdResult struct {
	Cmd *exec.Cmd
	Stdout string
	Stderr string
}

func ExecString(cmdLine []string) Result[StringCmdResult] {
	if len(cmdLine) == 0 {
		return Err[StringCmdResult](errors.FunctionalError{})
	}

	stdout := new(strings.Builder)
	stderr := new(strings.Builder)

	cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
	cmd.Stdin = nil
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	if err != nil {
		Err[StringCmdResult](err)
	}
	cmderr := cmd.Wait()
	if cmderr != nil {
		return Err[StringCmdResult](cmderr)
	}
	return Ok(StringCmdResult{
		Cmd:    cmd,
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	})
}
