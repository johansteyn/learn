package main

import (
	"fmt"
	"reflection/reflection"
)

func main() {
	fmt.Println("Go Standard Library: reflect")
	fmt.Println()

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
	fmt.Printf("x=%d (%T)\n", x, x)
	fmt.Printf("s=%s (%T)\n", s, s)
	fmt.Printf("b=%t (%T)\n", b, b)
	fmt.Printf("xptr=%v (%T)\n", xptr, xptr)
	fmt.Printf("yptr=%v (%T)\n", yptr, yptr)
	fmt.Printf("zptr=%v (%T)\n", zptr, zptr)
	fmt.Printf("sptr=%v (%T)\n", sptr, sptr)
	fmt.Printf("bptr=%v (%T)\n", bptr, bptr)
	fmt.Printf("xs=%v (%T)\n", xs, xs)
	fmt.Printf("bar1=%v (%T)\n", bar1, bar1)
	fmt.Printf("bar2=%v (%T)\n", bar2, bar2)
	fmt.Printf("bar3=%v (%T)\n", bar3, bar3)
	fmt.Printf("bars=%v (%T)\n", bars, bars)
	fmt.Println()

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
	fmt.Printf("f=%v (%T)\n", f, f)
	fmt.Println()

	fx := reflection.GetFieldByName(f, "X")
	fs := reflection.GetFieldByName(f, "S")
	fb := reflection.GetFieldByName(f, "B")
	fmt.Printf("fx=%v (%T)\n", fx, fx)
	fmt.Printf("fs=%v (%T)\n", fs, fs)
	fmt.Printf("fb=%v (%T)\n", fb, fb)
	fxptr := reflection.GetFieldByName(f, "Xptr")
	fsptr := reflection.GetFieldByName(f, "Sptr")
	fbptr := reflection.GetFieldByName(f, "Bptr")
	fmt.Printf("fxptr=%v (%T)\n", fxptr, fxptr)
	fmt.Printf("fsptr=%v (%T)\n", fsptr, fsptr)
	fmt.Printf("fbptr=%v (%T)\n", fbptr, fbptr)
	fxs := reflection.GetFieldByName(f, "Xs")
	fmt.Printf("fslice=%v (%T)\n", fxs, fxs)
	fbar := reflection.GetFieldByName(f, "Bar")
	fmt.Printf("fbar=%v (%T)\n", fbar, fbar)
	fbars := reflection.GetFieldByName(f, "Bars")
	fmt.Printf("fbars=%v (%T)\n", fbars, fbars)

	reflection.SetFieldByName(f, "X", 100)
	reflection.SetFieldByName(f, "S", "New String")
	reflection.SetFieldByName(f, "B", false)
	reflection.SetFieldByName(f, "Xptr", yptr)
	reflection.SetFieldByName(f, "Sptr", sptr)
	reflection.SetFieldByName(f, "Bptr", bptr)
	reflection.SetFieldByName(f, "Xs", []*int{zptr, yptr, xptr})
	reflection.SetFieldByName(f, "Bar", bar2)
	reflection.SetFieldByName(f, "Bars", []*bar{&bar3, &bar2, &bar1})
	fmt.Printf("f=%v (%T)\n", f, f)

	fmt.Println()
}

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
