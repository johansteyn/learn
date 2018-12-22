package recfun

object Main {
  def main(args: Array[String]) {
    println("Pascal's Triangle")
    for (row <- 0 to 10) {
      for (col <- 0 to row)
        print(pascal(col, row) + " ")
      println()
    }
  }

  /**
   * Exercise 1
   */
    def pascal(c: Int, r: Int): Int = if (c ==0 || c == r) 1 else pascal(c - 1, r - 1) + pascal(c, r - 1)
  
  /**
   * Exercise 2
   */
    def balance(chars: List[Char]): Boolean = {
      def loop(weight: Int, chars: List[Char]): Boolean = {
        if (chars.isEmpty) weight == 0 else {
          val w = chars.head match {
            case '(' => weight + 1
            case ')' => weight - 1
            case _ => weight
          }
          if (w < 0) false else loop(w, chars.tail)
        }
      }
      loop(0, chars)
    }
  
  /**
   * Exercise 3
   */
    def countChange(money: Int, coins: List[Int]): Int = {
      def recur(x: Int, y: Int): Int = {
        if (x < 0 || (x >= 1 && y <= 0)) 0 else {
          if (x == 0) 1 else recur(x, y - 1) + recur(x - coins(y - 1), y)
        }
      }
      recur(money, coins.size)
    }
  }
