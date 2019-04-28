import ch02._
import org.scalatest.{FlatSpec, Matchers}

class ch02Test extends FlatSpec with Matchers {

  "exercise01" should "return the 'signum' of a number" in {
    exercise01(-123) shouldBe -1
    exercise01(0) shouldBe 0
    exercise01(456) shouldBe 1
  }
  
  "exercise02" should "return Unit" in {
    exercise02() shouldBe ()
  }
  
  "exercise03" should "return Unit" in {
    exercise03() shouldBe ()
  }

  // Exercises 4 and 5 are procedures, ie. return Unit, so check console output for the side-effects

  "exercise04" should "return Unit" in {
    exercise04() shouldBe ()
  }

  "exercise05" should "return Unit" in {
    exercise05(3) shouldBe ()
  }

  "exercise06" should "compute the product of Unicode letters of a String" in {
    exercise06("Hello") shouldBe 9415087488L
    // H * e * l * l * o
    // 0x48L * 0x65L * 0x6CL * 0x6CL * 0x6FL
  }
}
