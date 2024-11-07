package test

import (
	"io"
	"net/http"
	"testing"
)

// LogResponseBody is a function that will log response body from provided http.Response
func LogResponseBody(t testing.TB, response *http.Response) {
	t.Helper()
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := io.ReadAll(response.Body)
    if err != nil {
		FatalFormat(t, 1, "Failed to read response body: %v", err)
    }
    t.Log(string(body))
}
