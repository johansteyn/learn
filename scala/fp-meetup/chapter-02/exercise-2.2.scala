object O extends App {

//	def isSorted[A](as: Array[A], ordered: (A, A) => Boolean): Boolean = {
//		for (i <- 0 until as.length - 1) {
//			val a = as(i)
//			val b = as(i + 1)
//			if (!ordered(a, b)) false
//		}
//		true
//	}

	def isSorted[A](as: Array[A], ordered: (A, A) => Boolean): Boolean = {
		as.toList match {
//			case first :: second :: _ => if (!ordered(first, second)) false else isSorted(as.tail, ordered)
			case first :: second :: _ => ordered(first, second) && isSorted(as.tail, ordered)
			case _ => true
		}
	}

	val ints1 = Array(1, 2, 3)
	val ints2 = Array(1, 3, 2)
	val chars1 = Array('a', 'b', 'c')
	val chars2 = Array('b', 'a', 'c')
	val strings1 = Array("Alice", "Bob", "Carol")
	val strings2 = Array("Carol", "Alice", "Bob")

	def compareInts(a: Int, b: Int): Boolean = { a < b }
	def compareChars(a: Char, b: Char): Boolean = { a < b }
	def compareStrings(a: String, b: String): Boolean = { a < b }

	println("ints1: " + isSorted(ints1, compareInts))
	println("ints2: " + isSorted(ints2, compareInts))
	println("chars1: " + isSorted(chars1, compareChars))
	println("chars2: " + isSorted(chars2, compareChars))
	println("strings1: " + isSorted(strings1, compareStrings))
	println("strings2: " + isSorted(strings2, compareStrings))

	println("ints1 (anonymous): " + isSorted(ints1, (a: Int, b: Int) => a < b))
	println("ints2 (anonymous): " + isSorted(ints2, (a: Int, b: Int) => a < b))

//	println("ints2 (anonymous): " + isSorted(ints2, (a: AnyVal, b: AnyVal) => a < b))
}


