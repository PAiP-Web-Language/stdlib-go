package test_test

import (
	"regexp"
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestContain1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "Contain should pass on correct substring")
    test.Contain(t, "test", "est", "Contain should pass on correct substring")
}

func TestContain2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "Contain should fail on incorrect substring")
    test.Contain(t, "test", "est2", "Contain should fail on incorrect substring (Expected fail)")
}

func TestNotContain1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "NotContain should pass on incorrect substring")
    test.NotContain(t, "test", "est2", "NotContain should pass on incorrect substring")
}

func TestNotContain2(_t *testing.T) {
    t := test.WrapTB(_t)
    t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "NotContain should fail on correct substring")
    test.NotContain(t, "test", "est", "NotContain should fail on correct substring (Expected fail)")
}

func TestMatchesRegexp1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "MatchesRegexp should pass on correct regexp")
    test.MatchesRegexp(t, "^t.*$", "test", "MatchesRegexp should pass on correct regexp")
}

func TestMatchesRegexp2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "MatchesRegexp should fail on incorrect regexp")
    test.MatchesRegexp(t, "^t2.*$", "test", "MatchesRegexp should fail on incorrect regexp (Expected fail)")
}

func TestNotMatchesRegexp1(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "NotMatchesRegexp should pass on incorrect regexp")
    test.NotMatchesRegexp(t, "^t2.*$", "test", "NotMatchesRegexp should pass on incorrect regexp")
}

func TestNotMatchesRegexp2(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "NotMatchesRegexp should fail on correct regexp")
    test.NotMatchesRegexp(t, "^t.*$", "test", "NotMatchesRegexp should fail on correct regexp (Expected fail)")
}

func TestMatchesRegexp3(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "MatchesRegexp should pass on correct regexp")
    test.MatchesRegexp(t, regexp.MustCompile("^t.*$"), "test", "MatchesRegexp should pass on correct regexp")
}

func TestMatchesRegexp4(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "MatchesRegexp should fail on incorrect regexp")
    test.MatchesRegexp(t, regexp.MustCompile("^t2.*$"), "test", "MatchesRegexp should fail on incorrect regexp (Expected fail)")
}

func TestNotMatchesRegexp3(_t *testing.T) {
    t := test.WrapTB(_t)
    defer t.Defer()
    test.LogIfFailed(t, "NotMatchesRegexp should pass on incorrect regexp")
    test.NotMatchesRegexp(t, regexp.MustCompile("^t2.*$"), "test", "NotMatchesRegexp should pass on incorrect regexp")
}

func TestNotMatchesRegexp4(_t *testing.T) {
    t := test.WrapTB(_t)
	t.ExpectedFailure = true
    defer t.Defer()
    test.LogIfFailed(t, "NotMatchesRegexp should fail on correct regexp")
    test.NotMatchesRegexp(t, regexp.MustCompile("^t.*$"), "test", "NotMatchesRegexp should fail on correct regexp (Expected fail)")
}
