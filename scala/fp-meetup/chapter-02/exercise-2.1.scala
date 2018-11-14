object O extends App {
	val n = args(0).toInt
	println("Fib(" + n + ") = " + fib(n))

	def fib(n: Int): Int = {
		@annotation.tailrec
		def loop(i: Int, a: Int, b: Int): Int = {
			if (i == 0) a else loop(i - 1, b, a + b)
		}
		loop(n, 0, 1)
	}
}


