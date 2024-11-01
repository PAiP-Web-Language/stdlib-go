package test

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/base"
	"github.com/PAiP-Web-Language/stdlib-go/internal"
)

func CompareEqual[T comparable](t testing.TB, expected, actual T, args ...any) {
	if expected != actual {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func Equal(t testing.TB, expected, actual any, args ...any) {
	if !reflect.DeepEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func DeepEqual(t testing.TB, expected, actual any, args ...any) {
	if !internal.ReflectDeeperEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func CompareNotEqual[T comparable](t testing.TB, expected, actual T, args ...any) {
	if expected == actual {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func NotEqual(t testing.TB, expected, actual any, args ...any) {
	if reflect.DeepEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func NotDeepEqual(t testing.TB, expected, actual any, args ...any) {
	if internal.ReflectDeeperEqual(expected, actual) {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
}

func MatchesRegexp(t testing.TB, regex base.LUI2[string, *regexp.Regexp], value string, args ...any) {
	switch regex.(type) {
	case string:
		r := regex.(string)
		reg, err := regexp.Compile(r)
		if err != nil {
			ErrorFormat(t, 1, fmt.Sprintf("failed to compile regexp %v", r), args...)
		}
		matched:= reg.MatchString(value)
		if !matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does not match regexp %v", value, r), args...)
		}
	case *regexp.Regexp:
		reg := regex.(*regexp.Regexp)
		matched:= reg.MatchString(value)
		if !matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does not match regexp %v", value, reg), args...)
		}
	}
}

func NotMatchesRegexp(t testing.TB, regex base.LUI2[string, *regexp.Regexp], value string, args ...any) {
	switch regex.(type) {
	case string:
		r := regex.(string)
		reg, err := regexp.Compile(r)
		if err != nil {
			ErrorFormat(t, 1, fmt.Sprintf("failed to compile regexp %v", r), args...)
		}
		matched:= reg.MatchString(value)
		if matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does match regexp %v", value, r), args...)
		}
	case *regexp.Regexp:
		reg := regex.(*regexp.Regexp)
		matched:= reg.MatchString(value)
		if matched {
			ErrorFormat(t, 1, fmt.Sprintf("value %v does match regexp %v", value, reg), args...)
		}
	}
}

func NotNil(t testing.TB, value any, args ...any) {
	if value == nil {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to be not nil", value), args...)
	}
}

func Nil(t testing.TB, value any, args ...any) {
	if value != nil {
		ErrorFormat(t, 1, fmt.Sprintf("expected %v to be nil", value), args...)
	}
}
