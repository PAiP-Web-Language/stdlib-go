package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/base"
)

func FileRead(t testing.TB, filename string) string {
	t.Helper()
	b, err := os.ReadFile(filename)
	if err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("failed to read file %v", filename))
		return ""
	}
	return string(b)
}

func FileWrite(t testing.TB, filename string, content string) {
	t.Helper()
	dir := filepath.Dir(filename)
	_ = os.MkdirAll(dir, 0700)
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("failed to write file %v", filename))
	}
}

func FileReadJson(t testing.TB, filename string, vref any) {
	t.Helper()
    content := FileRead(t, filename)
	if err := json.Unmarshal([]byte(content), vref); err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("unmarshal %s: %v", filename, err))
	}
}

func FileWriteJson(t testing.TB, filename string, vref any) {
	t.Helper()
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", TestConfig_FileJsonIndent)
	err := enc.Encode(vref)
	if err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("failed to encode json %v", vref))
	}
	FileWrite(t, filename, buf.String())
}

func FileShallowEqual(t testing.TB, got string, expected string) base.Result[[2]string] {
	t.Helper()
	got = strings.TrimSpace(got)
	expected = strings.TrimSpace(expected)

	if got == expected {
		return base.Ok([2]string{})
	}
	return base.ErrWithValue([2]string{got, expected}, fmt.Errorf("expected %v, got %v", expected, got))
}

func FileEqualJson(t testing.TB, filename string, expected any) {
	t.Helper()
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", TestConfig_FileJsonIndent)
	err := enc.Encode(expected)

	if err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("failed to encode json %v", expected))
	}

	content := FileRead(t, filename)

	if FileShallowEqual(t, content, buf.String()).IsErr() {
		if TestConfig_FileFullInfo {
			ErrorFormat(t, 1, fmt.Sprintf("File %s containing %v does not match expected json %v", filename, content, buf.String()))
		} else {
			ErrorFormat(t, 1, fmt.Sprintf("File %s does not match expected json", filename))
		}
	}
}
