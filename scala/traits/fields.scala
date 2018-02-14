// https://stackoverflow.com/questions/25141594/when-overriding-a-trait-why-the-value-is-strange

// Problem:
// Memory is reserved for trait B, including space for "a" (with value uninitialized, ie. zero)
// When trait A's constructor is run, the value of "a" (zero) is used when computing the value for "b".
// Only when B's constructor is run does "a" get the value 10.

// Solution:
// Make "b" lazy so that its value is only computed when first accessed, ie. after "a" has value 10.

// Is it really a problem?
// I don't think so.
// It's a contrived example - I cannot think of a situation where I would 
// calculate one field's value based of the values of other fields.

trait A {
	val a = 3
//	val b = a + 2
	lazy val b = a + 2
}

trait B extends A {
	override val a = 10
}

object X extends B

println(X.b)

