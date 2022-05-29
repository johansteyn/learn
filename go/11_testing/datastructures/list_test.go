// Same package name as the package we are testing
package datastructures

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// Without a TestMain each test is called directly.
// With a TestMain only it is called directly, once only,
// and each test is called via the Run method.
func TestMain(m *testing.M) {
	fmt.Println("Setup stuff here...")
	exitVal := m.Run()
	fmt.Println("Cleanup here...")
	os.Exit(exitVal)
}

// Test function names must start with (or be equal to) "Test"
func Test(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	//t.Error("Error, but test will continue...")
	//fmt.Println("Continuing...")
	//t.Fatal("Fatal, test will abort.")
	//fmt.Println("This statement cannot be reached")
}

// Test function naming convention: TestPublicFunction or Test_privateFunction
func TestNew(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	list := New()
	expectedSize := 0
	// Since we're in teh same package we can access the "size" field directly
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	_, err := list.Get(0)
	if err == nil {
		t.Errorf("expected error but got none")
	}
}

// Own addition: add an optional _Variation
func TestAdd_10(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	list := New()
	for i := 0; i < 10; i++ {
		list.Add1(11 + i)
	}
	expectedSize := 10
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	expectedValue := 11
	value, err := list.Get(0)
	if err != nil {
		t.Errorf("unexpected expected error")
	}
	if value != expectedValue {
		t.Errorf("expected value %d but got %d\n", expectedValue, value)
	}

	_, err = list.Get(10)
	if err == nil {
		t.Errorf("expected error but got none")
	}

	_, err = list.Get(-1)
	if err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestAdd_100(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	list := New()
	for i := 0; i < 100; i++ {
		list.Add1(i)
	}
	expectedSize := 100
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	expectedValue := 42
	value, err := list.Get(42)
	if err != nil {
		t.Errorf("unexpected expected error")
	}
	if value != expectedValue {
		t.Errorf("expected value %d but got %d\n", expectedValue, value)
	}
}

// Table-driven testing is common in Go
func TestAdd_Table(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
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
			list.Add1(i)
		}
		expectedSize := test.n
		size := list.size
		if size != expectedSize {
			t.Errorf("expected size %d but got %d\n", expectedSize, size)
		}
	}
}

func TestAdd_Randomized(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	fmt.Printf("%s - adding %d elements\n", t.Name(), n)
	list := New()
	for i := 0; i < n; i++ {
		list.Add1(i)
	}
	expectedSize := n
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
}

// Reads the size from a test data file
func TestAdd_File(t *testing.T) {
	n, err := readTestData(t)
	if err != nil {
		t.Fatalf("error reading test data file: %v", err)
	}
	fmt.Printf("%s - adding %d elements\n", t.Name(), n)
	list := New()
	for i := 0; i < n; i++ {
		list.Add1(i)
	}
	expectedSize := n
	size := list.size
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
}

// Tests are run in the package directory
// Subdirectory "testdata" holds files with test data
func readTestData(t *testing.T) (int, error) {
	fmt.Printf("Reading test data...\n")
	var n int
	file, err := os.Open("testdata/testfile")
	if err != nil {
		return 0, err
	}
	// The Cleanup method is used to clean up temporary resources for a single test.
	// Unlike defer, which runs at the end of this function, Cleanup runs when the test completes.
	t.Cleanup(func() {
		file.Close()
		fmt.Printf("Closed test data file.\n")
	})
	_, err = fmt.Fscanln(file, &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Test a private function
func Test_slice(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	n := 10
	list := New()
	for i := 0; i < n; i++ {
		list.Add1(i)
	}
	s := list.slice()
	expectedSize := n
	size := len(s)
	if size != expectedSize {
		t.Errorf("expected size %d but got %d\n", expectedSize, size)
	}
	for i := 0; i < 10; i++ {
		if s[i] != i {
			t.Errorf("expected value %d but got %d\n", i, s[i])
		}
	}
}

/*
// Test main using a mock getTemperature function
func Test_main(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
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
*/
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
			list.Add1(j)
		}
	}
}
func listAdd2(b *testing.B, size int) {
	list := New()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			list.Add2(j)
		}
	}
}
