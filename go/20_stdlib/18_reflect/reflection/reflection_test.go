package reflection

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type foo struct {
	X int
	S string
	B bool

	Xptr *int
	Sptr *string
	Bptr *bool

	Xs   []*int
	Bar  bar
	Bars []*bar
}

type bar struct {
	A int
	B string
	C bool
}

func TestMain(m *testing.M) {
	fmt.Println("Setup stuff here...")
	exitVal := m.Run()
	fmt.Println("Cleanup here...")
	os.Exit(exitVal)
}

func TestGetFieldByName(t *testing.T) {
	fmt.Printf("%s\n", t.Name())
	x := 42
	y := 7
	z := 13
	s := "The meaning of life"
	b := true
	xptr := &x
	yptr := &y
	zptr := &z
	sptr := &s
	bptr := &b
	xs := []*int{xptr, yptr, zptr}
	bar1 := bar{
		A: 1,
		B: "One",
		C: false,
	}
	bar2 := bar{
		A: 2,
		B: "Two",
		C: true,
	}
	bar3 := bar{
		A: 3,
		B: "Three",
		C: false,
	}
	bars := []*bar{&bar1, &bar2, &bar3}
	f := &foo{
		X:    x,
		S:    s,
		B:    b,
		Xptr: xptr,
		Sptr: sptr,
		Bptr: bptr,
		Xs:   xs,
		Bar:  bar1,
		Bars: bars,
	}

	fx := GetFieldByName(f, "X")
	if fx != x {
		t.Errorf("expected %d but got %v", x, fx)
	}
	if _, ok := fx.(int); !ok {
		t.Errorf("expected type int but got %T", fx)
	}
	fs := GetFieldByName(f, "S")
	if fs != s {
		t.Errorf("expected %s but got %v", s, fs)
	}
	if _, ok := fs.(string); !ok {
		t.Errorf("expected type string but got %T", fs)
	}
	fb := GetFieldByName(f, "B")
	if fb != b {
		t.Errorf("expected %t but got %v", b, fb)
	}
	if _, ok := fb.(bool); !ok {
		t.Errorf("expected type bool but got %T", fb)
	}
	fxptr := GetFieldByName(f, "Xptr")
	if fxptr != xptr {
		t.Errorf("expected %v but got %v", xptr, fxptr)
	}
	if _, ok := fxptr.(*int); !ok {
		t.Errorf("expected type *int but got %T", fxptr)
	}
	fsptr := GetFieldByName(f, "Sptr")
	if fsptr != sptr {
		t.Errorf("expected %v but got %v", sptr, fsptr)
	}
	if _, ok := fsptr.(*string); !ok {
		t.Errorf("expected type *string but got %T", fsptr)
	}
	fbptr := GetFieldByName(f, "Bptr")
	if fbptr != bptr {
		t.Errorf("expected %v but got %v", bptr, fbptr)
	}
	if _, ok := fbptr.(*bool); !ok {
		t.Errorf("expected type *bool but got %T", fbptr)
	}
	fxs := GetFieldByName(f, "Xs")
	if !reflect.DeepEqual(fxs, xs) {
		t.Errorf("expected %v but got %v", xs, fxs)
	}
	if _, ok := fxs.([]*int); !ok {
		t.Errorf("expected type []*int but got %T", fxs)
	}
	fbar := GetFieldByName(f, "Bar")
	if fbar != bar1 {
		t.Errorf("expected %v but got %v", bar1, fbar)
	}
	if _, ok := fbar.(bar); !ok {
		t.Errorf("expected type bar but got %T", fbar)
	}
	fbars := GetFieldByName(f, "Bars")
	if !reflect.DeepEqual(fbars, bars) {
		t.Errorf("expected %v but got %v", bars, fbars)
	}
	if _, ok := fbars.([]*bar); !ok {
		t.Errorf("expected type []*bar but got %T", fbars)
	}
	fmt.Println("All tests passed!")
	/*
		fs := GetFieldByName(f, "S")
		fb := GetFieldByName(f, "B")
		//fmt.Printf("fx=%v (%T)\n", fx, fx)
		//fmt.Printf("fs=%v (%T)\n", fs, fs)
		//fmt.Printf("fb=%v (%T)\n", fb, fb)
		fxptr := GetFieldByName(f, "Xptr")
		fsptr := GetFieldByName(f, "Sptr")
		fbptr := GetFieldByName(f, "Bptr")
		//fmt.Printf("fxptr=%v (%T)\n", fxptr, fxptr)
		//fmt.Printf("fsptr=%v (%T)\n", fsptr, fsptr)
		//fmt.Printf("fbptr=%v (%T)\n", fbptr, fbptr)
		fxs := GetFieldByName(f, "Xs")
		//fmt.Printf("fslice=%v (%T)\n", fxs, fxs)
		fbar := GetFieldByName(f, "Bar")
		//fmt.Printf("fbar=%v (%T)\n", fbar, fbar)
		fbars := GetFieldByName(f, "Bars")
		//fmt.Printf("fbars=%v (%T)\n", fbars, fbars)
	*/

}
