object O extends App {
	def curry[A,B,C](f: (A, B) => C): A => (B => C)

}


