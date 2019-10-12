object FrogJmp extends App {
  val max = 1000000000
  def solution(x: Int, y: Int, d: Int): Int = {
    require(x >= 1 && x <= max && y >= 1 && y <= max && d >= 1 && d <= max & x <= y)
    (y - x) / d + (if ((y - x) % d > 0) 1 else 0)
  }

  def test(x: Int, y: Int, d: Int, expected: Int) = {
    print(s"Testing x=$x, y=$y, d=$d... ")
    try {
      val actual = solution(x, y, d)
      try {
        assert(actual == expected)
        println("Passed.")
      } catch {
        case ae: AssertionError => println(s"  Failed: Expected $expected but got $actual")
      }
    } catch {
      case iae: IllegalArgumentException => println("  Failed: " + iae)
    }
  }

  println("Expected failures:")
  test(-1, 1, 1, 0)
  test(1, -1, 1, 0)
  test(1, 1, -1, 0)
  test(9, 1, 1, 0)
  test(1234567890, 1, 1, 0)

  println("Expected passes:")
  test(10, 85, 30, 3)
  test(10, 1000000000, 2, 499999995)
}

