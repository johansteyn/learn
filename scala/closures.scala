object Main extends App {
	// The "free" variable n has an initial value
	var n = 10
	val foo = (x: Int) => x + n
	// Form a closure using the value of n
	println(foo(1)) // 11

	// Change the value of n
	n = 20
	// Form a separate closure using the new value of n
	println(foo(1)) // 21

	def bar(f: Int => Int) = {
		// This n is not a "free" variable - it isn't in any enclosing scope of foo
		val n = 100
		// No closure is formed here - a closure is only formed when bar is invoked
		f(1)
	}
	// A closure is formed here, with foo using the n that is in scope here
	println(bar(foo)) // 21 
}
