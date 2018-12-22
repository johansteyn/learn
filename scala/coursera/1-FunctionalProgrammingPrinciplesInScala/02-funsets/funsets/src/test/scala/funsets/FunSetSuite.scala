package funsets

import org.scalatest.FunSuite


import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner

/**
 * This class is a test suite for the methods in object FunSets. To run
 * the test suite, you can either:
 *  - run the "test" command in the SBT console
 *  - right-click the file in eclipse and chose "Run As" - "JUnit Test"
 */
@RunWith(classOf[JUnitRunner])
class FunSetSuite extends FunSuite {

  /**
   * Link to the scaladoc - very clear and detailed tutorial of FunSuite
   *
   * http://doc.scalatest.org/1.9.1/index.html#org.scalatest.FunSuite
   *
   * Operators
   *  - test
   *  - ignore
   *  - pending
   */

  /**
   * Tests are written using the "test" operator and the "assert" method.
   */
  // test("string take") {
  //   val message = "hello, world"
  //   assert(message.take(5) == "hello")
  // }

  /**
   * For ScalaTest tests, there exists a special equality operator "===" that
   * can be used inside "assert". If the assertion fails, the two values will
   * be printed in the error message. Otherwise, when using "==", the test
   * error message will only say "assertion failed", without showing the values.
   *
   * Try it out! Change the values so that the assertion fails, and look at the
   * error message.
   */
  // test("adding ints") {
  //   assert(1 + 2 === 3)
  // }


  import FunSets._

  test("contains is implemented") {
    assert(contains(x => true, 100))
  }

  /**
   * When writing tests, one would often like to re-use certain values for multiple
   * tests. For instance, we would like to create an Int-set and have multiple test
   * about it.
   *
   * Instead of copy-pasting the code for creating the set into every test, we can
   * store it in the test class using a val:
   *
   *   val s1 = singletonSet(1)
   *
   * However, what happens if the method "singletonSet" has a bug and crashes? Then
   * the test methods are not even executed, because creating an instance of the
   * test class fails!
   *
   * Therefore, we put the shared values into a separate trait (traits are like
   * abstract classes), and create an instance inside each test method.
   *
   */

  trait TestSets {
    val s1 = singletonSet(1)
    val s2 = singletonSet(2)
    val s3 = singletonSet(3)
  }

  /**
   * This test is currently disabled (by using "ignore") because the method
   * "singletonSet" is not yet implemented and the test would fail.
   *
   * Once you finish your implementation of "singletonSet", exchange the
   * function "ignore" by "test".
   */
  test("singletonSet(1) contains 1") {

    /**
     * We create a new instance of the "TestSets" trait, this gives us access
     * to the values "s1" to "s3".
     */
    new TestSets {
      /**
       * The string argument of "assert" is a message that is printed in case
       * the test fails. This helps identifying which assertion failed.
       */
      assert(contains(s1, 1), "Singleton")
    }
  }

  test("union contains all elements of each set") {
    new TestSets {
      val s = union(s1, s2)
      assert(contains(s, 1), "Union 1")
      assert(contains(s, 2), "Union 2")
      assert(!contains(s, 3), "Union 3")
    }
  }


  test("intersect contains only elements common to every set") {
    new TestSets {
      val s12 = union(s1, s2)
      val s13 = union(s1, s3)
      val s = intersect(s12, s13)
      assert(contains(s, 1), "Intersect 1")
      assert(!contains(s, 2), "Intersect 2")
      assert(!contains(s, 3), "Intersect 3")
    }
  }


  test("diff contains elements in first set that are not in the second set") {
    new TestSets {
      val s12 = union(s1, s2)
      val s13 = union(s1, s3)
      val s = diff(s12, s13)
      assert(!contains(s, 1), "Diff 1")
      assert(contains(s, 2), "Diff 2")
      assert(!contains(s, 3), "Diff 3")
    }
  }


  test("filter contains elements in a set that satisfy a given function") {
    new TestSets {
      val s123 = union(s1, union(s2, s3))
      val sid = filter(s123, (x) => (x == x))
      val sless3 = filter(s123, (x) => (x < 3))
      assert(contains(sid, 1), "Filter sid 1")
      assert(contains(sid, 2), "Filter sid 2")
      assert(contains(sid, 3), "Filter sid 3")
      assert(contains(sless3, 1), "Filter sless3 1")
      assert(contains(sless3, 2), "Filter sless3 2")
      assert(!contains(sless3, 3), "Filter sless3 3")
    }
  }

  test("forall tests whether a given predicate is true for all elements of the set.") {
    new TestSets {
      val s12 = union(s1, s2)
      val s13 = union(s1, s3)
      val s123 = union(s1, union(s2, s3))
      assert(forall(s12, (x) => (x < 3)), "Forall s12 less than 3")
      assert(!forall(s123, (x) => (x < 3)), "Forall s123 less than 3")
      assert(forall(s13, (x) => (x % 2 == 1)), "Forall s13 odd numbers")
    }
  }

  test("exists tests whether a set contains at least one element for which the given predicate is true.") {
    new TestSets {
      val s12 = union(s1, s2)
      val s13 = union(s1, s3)
      val s123 = union(s1, union(s2, s3))
      assert(exists(s12, (x) => (x == 2)), "Exists s12 equals 2")
      assert(exists(s12, (x) => (x % 2 == 0)), "Exists s12 even number")
      assert(!exists(s13, (x) => (x % 2 == 0)), "Exists s12 even number")
    }
  }

  test("map transforms a given set into another one by applying to each of its elements the given function.") {
    new TestSets {
      val s123 = union(s1, union(s2, s3))
      printSet(s123)
      val s234 = map(s123, (x) => (x + 1))
      printSet(s234)
      val s246 = map(s123, (x) => (x * 2))
      printSet(s246)
      assert(!exists(s234, (x) => (x == 1)), "Not exists s234 equals 1")
      assert(exists(s234, (x) => (x == 2)), "Exists s234 equals 2")
      assert(exists(s234, (x) => (x == 3)), "Exists s234 equals 3")
      assert(exists(s234, (x) => (x == 4)), "Exists s234 equals 4")
      assert(!exists(s234, (x) => (x == 5)), "Not exists s234 equals 5")
    }
  }


}
