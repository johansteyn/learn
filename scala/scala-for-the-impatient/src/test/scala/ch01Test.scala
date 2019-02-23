import ch01._
import scala.math.BigInt
import org.scalatest.{FlatSpec, Matchers}

class ch01Test extends FlatSpec with Matchers {

  "exercise06" should "compute 2^1024" in {
    exercise06() shouldBe BigInt("1797693134862315907729305190789024733617976978942306572734" +
      "300811577326758055009631327084773224075360211201138798713933576587897688144166224928474" +
      "3063947412437776789342486548527630221960124609411945308295208500576883815068234246288147" +
      "3913110540827237163350510684586298239947245938479716304835356329624224137216")
  }
 
  "exercise08" should "yield a string such as 'qsnvbevtomcj38o06kul'" in {
    val result: String = exercise08()
    result.length shouldBe 20
    result.matches("[a-z0-9]+") shouldBe true
  }

  "exercise09a" should "return the first character" in {
    exercise09a("Quick brown fox jumps over the lazy dog") shouldBe 'Q'
  }

  "exercise09b" should "return the first character" in {
    exercise09b("Quick brown fox jumps over the lazy dog") shouldBe 'Q'
  }

  "exercise09c" should "return the last character" in {
    exercise09c("Quick brown fox jumps over the lazy dog") shouldBe 'g'
  }

  "exercise09d" should "return the last character" in {
    exercise09d("Quick brown fox jumps over the lazy dog") shouldBe 'g'
  }
}
