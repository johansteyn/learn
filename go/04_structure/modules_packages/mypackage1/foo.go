package mypackage1

func Foo() string {
	return "Foo"
}

func Foobar() string {
	// Even though "bar" is not exported, it is visible from other files in the same package
	return Foo() + bar()
}

