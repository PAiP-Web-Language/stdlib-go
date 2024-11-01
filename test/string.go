package test

import (
	"fmt"
	"strings"
	"testing"
)

func Contain(t testing.TB, s string, substr string, args ...any) {
	if !strings.Contains(s, substr) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to contain %v", s, substr), args...)
	}
}

func NotContain(t testing.TB, s string, substr string, args ...any) {
	if strings.Contains(s, substr) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to contain %v", s, substr), args...)
	}
}
