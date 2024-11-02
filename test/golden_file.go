package test

import (
	"fmt"
	"path/filepath"
	"testing"
)

// GoldenFileRun runs function f for each file matching filenameGlob
func GoldenFileRun(t *testing.T, filenameGlob string, f func(t *testing.T, match string)) {
	t.Helper()
	matches, err := filepath.Glob(filenameGlob)
	if err != nil {
		TestUtilErrorFormat(t, 1, fmt.Sprintf("failed to glob %v", filenameGlob))
	}
	for _, match := range matches {
		name := filepath.Base(match)
        t.Run(name, func(t *testing.T) {
            f(t, match)
        })
	}
}
