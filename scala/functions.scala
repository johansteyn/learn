object Main extends App {
	// A normal function definition:
	def addOne(x: Int): Int = x + 1
	// An anonymous (lamda) function assigned to a val:
	val addTwo: (Int) => Int = (x) => x + 2
	// NOTE: We could also use:
	// val addTwo: Int => Int = x => x + 2
	// ie. parentheses are redundant, but can help make the code clearer
	
	// Both can be invoked in the same way:
	println(addOne(12))
	println(addTwo(40))

	// And both can be passed as a function parameter:
	var numbers = List(1, 2, 3, 4)
	numbers = numbers.map(addOne)
	println(numbers.mkString(","))
	numbers = numbers.map(addTwo)
	println(numbers.mkString(","))

	// Q: So, what's the difference between a normal function and an anonymous one?
	// A; The only difference is that we don't need to give a name to an anonymous function.
	// ie. We can pass a function that adds 3 without calling it "addThree"
	numbers = numbers.map((i: Int) => i + 3) 
	println(numbers.mkString(","))

	// Note that we can use infix notation
	// ie. space instead of a dot, and replace the parentheses with curly braces: 
	numbers = numbers map {(i: Int) => i + 4} 
	println(numbers.mkString(","))

	// We can write our own function that takes a function as a parameter.
	// Our function is termed a "higher-order" function
	// Here we want to apply any given function (that takes an Int and returns an Int)
	// to any given number, eg: add one, or two, or three, or multiply by something 
	def add(n: Int, f: Int => Int) = f(n)
	println(add(11, addOne)) // 12
	println(add(40, addTwo)) // 42
	println(add(11, (i: Int) => i * 3)) // 33

	// We can write a function that returns a function.
	// Eg: instead of defining addOne, addTwo, addThree, etc.
	// we can define a function returns a function which
	// can take any number (n) and add the specified increment value:
	def addNum(increment: Int) = (n: Int) => n + increment
	// Or one that multiplies by the specified factor:
	def mulBy(factor: Int) = (n: Int) => n * factor
	println(add(40, addNum(2))) // 42
	println(add(40, addNum(2))) // 42
	println(add(11, mulBy(3))) // 33
}
