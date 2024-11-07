package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/base"
)

// Contain checks if string contains substring
func Contain(t testing.TB, s string, substr string, args ...any) {
	t.Helper()
	if !strings.Contains(s, substr) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to contain %v", s, substr), args...)
	}
}

// NotContain checks if string does not contain substring
func NotContain(t testing.TB, s string, substr string, args ...any) {
	t.Helper()
	if strings.Contains(s, substr) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to contain %v", s, substr), args...)
	}
}

// MatchesRegexp checks if string matches regexp
func MatchesRegexp(t testing.TB, regex base.LUI2[string, *regexp.Regexp], value string, args ...any) {
	t.Helper()
	switch r := regex.(type) {
	case string:
		reg, err := regexp.Compile(r)
		if err != nil {
			ErrorFormat(t, 1, fmt.Sprintf("failed to compile regexp %v", r), args...)
		}
		matched := reg.MatchString(value)
		if !matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does not match regexp %v", value, r), args...)
		}
	case *regexp.Regexp:
		matched := r.MatchString(value)
		if !matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does not match regexp %v", value, r), args...)
		}
	}
}

// NotMatchesRegexp checks if string does not match regexp
func NotMatchesRegexp(t testing.TB, regex base.LUI2[string, *regexp.Regexp], value string, args ...any) {
	t.Helper()
	switch r := regex.(type) {
	case string:
		reg, err := regexp.Compile(r)
		if err != nil {
			ErrorFormat(t, 1, fmt.Sprintf("failed to compile regexp %v", r), args...)
		}
		matched := reg.MatchString(value)
		if matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does match regexp %v", value, r), args...)
		}
	case *regexp.Regexp:
		matched := r.MatchString(value)
		if matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does match regexp %v", value, r), args...)
		}
	}
}
