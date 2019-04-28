object ch02 {

  def exercise01(x: Int): Int = if (x < 0) -1 else if (x > 0) 1 else 0

  def exercise02() = {}

  def exercise03(): Unit = {
    var y: Int = 0
    var x: Unit = ()
    x = y = 1
  }

  def exercise04(): Unit = {
    exercise05(10)
  }

  def exercise05(n: Int): Unit = {
    for (i <- n to 0 by -1) println(i)
  }

  def exercise06(s: String): Long = {
    var x = 1L
    for (c <- s) x *= c
    x
  }
}

