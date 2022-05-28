// Same package name as the package we are testing
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// The test function name must start with (or be equal to) "Test"
func Test(t *testing.T) {
	fmt.Printf("Test\n")
}

func TestEmptyList(t *testing.T) {
	fmt.Printf("TestEmpty\n")
	list := New()
	expectedSize := 0
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	_, err := list.get(0)
	if err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestListWithTenNodes(t *testing.T) {
	fmt.Printf("TestListWithTenNodes\n")
	list := New()
	for i := 0; i < 10; i++ {
		list.add1(11 + i)
	}
	expectedSize := 10
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	expectedValue := 11
	value, err := list.get(0)
	if err != nil {
		t.Errorf("unexpected expected error")
	}
	if value != expectedValue {
		t.Errorf("expected value %d but got %d\n", expectedValue, value)
	}

	_, err = list.get(10)
	if err == nil {
		t.Errorf("expected error but got none")
	}

	_, err = list.get(-1)
	if err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestListWithHundredNodes(t *testing.T) {
	fmt.Printf("TestListWithHundredNodes\n")
	list := New()
	for i := 0; i < 100; i++ {
		list.add1(i)
	}
	expectedSize := 100
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	expectedValue := 42
	value, err := list.get(42)
	if err != nil {
		t.Errorf("unexpected expected error")
	}
	if value != expectedValue {
		t.Errorf("expected value %d but got %d\n", expectedValue, value)
	}
}

// Table-driven testing is common in Go
func TestListUsingTable(t *testing.T) {
	fmt.Printf("TestListUsingTable\n")
	var tests = []struct {
		n int
	}{
		{0},
		{1},
		{10},
		{100},
		{1000},
	}
	for _, test := range tests {
		list := New()
		for i := 0; i < test.n; i++ {
			list.add1(i)
		}
		expectedSize := test.n
		size := list.size
		if size != expectedSize {
			t.Errorf("expected size %d but got %d\n", expectedSize, size)
		}
	}
}

func TestListRandomized(t *testing.T) {
	fmt.Printf("TestListRandomized\n")
	list := New()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	fmt.Printf("Adding %d elements...\n", n)
	for i := 0; i < n; i++ {
		list.add1(i)
	}
	expectedSize := n
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
}

// Test main using a mock getTemperature function
func TestMain(t *testing.T) {
	// In case any other tests need the getTemperature function, save it
	saved := getTemperature
	// And restore it at the end of this test
	defer func() {
		getTemperature = saved
	}()
	// This is the mock implementation
	getTemperature = func() int {
		return 42
	}
	main()
}
func BenchmarkListAdd1_10(b *testing.B) {
	listAdd1(b, 10)
}

func BenchmarkListAdd1_100(b *testing.B) {
	listAdd1(b, 100)
}

func BenchmarkListAdd1_1000(b *testing.B) {
	listAdd1(b, 1000)
}

func BenchmarkListAdd2_10(b *testing.B) {
	listAdd2(b, 10)
}

func BenchmarkListAdd2_100(b *testing.B) {
	listAdd2(b, 100)
}

func BenchmarkListAdd2_1000(b *testing.B) {
	listAdd2(b, 1000)
}

func listAdd1(b *testing.B, size int) {
	list := New()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			list.add1(j)
		}
	}
}
func listAdd2(b *testing.B, size int) {
	list := New()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			list.add2(j)
		}
	}
}
