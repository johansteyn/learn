import scala.math.BigInt

object ch01 {

  def exercise06(): BigInt = BigInt(2).pow(1024)

  def exercise07(): BigInt = {
    import scala.BigInt.probablePrime
    import scala.util.Random
    probablePrime(100, Random)
  }

  def exercise08(): String = {
    util.Random.alphanumeric.take(20).mkString.toLowerCase
  }

  def exercise09a(s: String): Char = s(0)
  def exercise09b(s: String): Char = s.head
  def exercise09c(s: String): Char = s(s.length - 1)
  def exercise09d(s: String): Char = s.last

  def exercise10(): Unit = {
    val str = "Some string"

    val first10 = str.take(15) // "Some string", no IndexOutOfBoundsException exception !
    val empty = "".drop(5) // "", no IndexOutOfBoundsException exception !

    val take = str.take(4) // "Some"
    val drop = str.drop(5) // "string"

    val takeRight = str.takeRight(6) // "string"
    val dropRight = str.dropRight(7) // "Some"
  }
}

