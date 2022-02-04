package mypackage

import (
	"testing"
)

func TestName(t *testing.T) {
	expected := "Johan"
	got := Name()
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}
