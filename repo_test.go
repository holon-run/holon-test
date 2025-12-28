package holon_test

import (
	"testing"
)

func TestRepoRoot(t *testing.T) {
	// A simple test to verify the test file is working
	got := "hello"
	want := "hello"

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
