import annotation.tailrec

object BinaryGap extends App {
  def solution(n: Int): Int = {
    @tailrec
    def f(n: Int, acc: Int, max: Int): Int = {
      if (n <= 1) max else {
        if (((n >> 1) & 1) == 1) 
          f(n >> 1, 0, max)
        else
          f(n >> 1, acc + 1, if (acc + 1 > max) acc + 1 else max)
      }
    }
    @tailrec
    def strip(n: Int): Int = if (n <= 0) n else if ((n & 1) == 1) n else strip(n >> 1)
    f(strip(n), 0, 0)
  }

  def test(n: Int, expected: Int) = {
    val b = n.toBinaryString
    println(s"Testing $n ($b)...")
    val actual = solution(n)
    try {
      assert(actual == expected)
    } catch {
      case ae: AssertionError => println(s"  Failed for n=$n: Expected $expected but got $actual")
    }
  }

  test(0, 0)    // 0
  test(1, 0)    // 1
  test(7, 0)    // 111
  test(5, 1)    // 101
  test(9, 2)    // 1001
  test(72, 2)   // 1001000
  test(1185, 4) // 10010100001
}

