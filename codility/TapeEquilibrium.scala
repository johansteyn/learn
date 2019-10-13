import annotation.tailrec

object TapeEquilibrium extends App {
  def solution(a: Array[Int]): Int = {
    val total = a. foldLeft(0)((a, b) => a + b)
    val sums = a.drop(1).scanLeft(a.head) {
      case (a, b) => a + b
    }.take(a.size - 1)
    val diffs = sums.map(n => Math.abs(n - (total - n)))
    diffs.min
  }

  def test(a: Array[Int], expected: Int) = {
    println("Testing " + a.mkString(", ") + "...")
    val actual = solution(a)
    try {
      assert(actual == expected)
    } catch {
      case ae: AssertionError => println(s"  Failed: Expected $expected but got $actual")
    }
  }

  test(Array(3, 1, 2, 4, 3), 1)
  test(Array(1, 2, 3, 4, 5, 6, 7, 8, 9), 3)
  //         1  3  6 10 15 21 28 36 45 
  //        44 41 39 35 30 24 17  9 
  //        32 38 33 25 15 3  11 27
}

