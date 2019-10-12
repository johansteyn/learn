object CyclicRotation extends App {
  def solution(a: Array[Int], k: Int): Array[Int] = {
    require(k >= 0 && k <= 100)
    val n = k % a.size
    if (a.size <= 0 || n == 0) a else {
      val b = new Array[Int](a.size)
      for (i <- 0 until a.size) {
        val j = (i + a.size - n) % a.size
        if (a(j) < -1000 || a(j) > 1000) throw new IllegalArgumentException("Element out of range: " + a(j))
        b(i) = a(j)
      }
      b
    }
  }

  def test(a: Array[Int], k: Int, expected: Array[Int]) = {
    println("Testing " + a.mkString(", ") + s" rotated $k times...")
    try {
      val actual = solution(a, k)
      try {
        assert(actual.sameElements(expected))
      } catch {
        case ae: AssertionError => println("  Failed: Expected " + expected.mkString(", ") + " but got " + actual.mkString(", "))
      }
    } catch {
      case iae: IllegalArgumentException => println("  Failed: " + iae)
    }
  }


  test(Array(1), 1, Array(11))
  test(Array(1), 99, Array(1))
  test(Array(1, 2, 3), 0, Array(1, 2, 3))
  test(Array(1, 2, 3), 1, Array(3, 1, 2))
  test(Array(1, 2, 3), 2, Array(2, 3, 1))
  test(Array(1, 2, 3), 3, Array(1, 2, 3))
  test(Array(1, 2, 3), 11, Array(2, 3, 1))
  test(Array(1, 2, 3), 33, Array(1, 2, 3))
  test(Array(1, 2, 3, 4, 5, 6, 7), 3, Array(5, 6, 7, 1, 2, 3, 4))
  test(Array(1, 2222, 3), 1, Array())
  test(Array(1, -2222, 3), 1, Array())
}

