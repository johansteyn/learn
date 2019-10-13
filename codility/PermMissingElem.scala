object PermMissingElem extends App {
  def solution(a: Array[Int]): Int = {
    require(a.size <= 100000)
    val set = scala.collection.mutable.SortedSet(1)
    for (i <- 2 to (a.size + 1)) set += i
    for (n <- a) set -= n
    set.head
  }

  def test(a: Array[Int], expected: Int) = {
    println("Testing " + a.mkString(", ") + s"... ")
    val actual = solution(a)
    try {
      assert(actual == expected)
      println("Passed.")
    } catch {
      case ae: AssertionError => println(s"  Failed: Expected $expected but got $actual")
    }
  }

  test(Array(2, 3, 1, 5), 4)
  test(Array(1, 4, 3, 9, 6, 7, 5, 8), 2)
}

