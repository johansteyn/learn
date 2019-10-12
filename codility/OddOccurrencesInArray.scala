import annotation.tailrec

object OddOccurrencesInArray extends App {
  def solution(a: Array[Int]): Int = {
    var set = scala.collection.mutable.Set[Int]()
    for (n <- a) if (set.contains(n)) set -= n else set += n
    set.toList(0)
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

  test(Array(1), 1)
  test(Array(1, 2, 1), 2)
  test(Array(1, 2, 1, 2, 1, 3, 1), 3)
  test(Array(1, 1, 1, 1, 1, 2, 1, 3, 1, 4, 4, 4, 2, 3, 1), 4)
}

