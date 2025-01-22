// Different package name as the package we are testing
// This means we can only test the public API
package datastructures_test

import (
	"fmt"
	"testing"
	"testlist/datastructures"
)

func TestPublicAdd(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	list := datastructures.New()
	list.Add1(42)
	expectedSize := 1
	// Cannot access the "size" field directly - must call the Size method
	//size := list.size
	size := list.Size()
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	expectedValue := 42
	value, err := list.Get(0)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if value != expectedValue {
		t.Errorf("expected value %d but got %d\n", expectedValue, value)
	}
}
