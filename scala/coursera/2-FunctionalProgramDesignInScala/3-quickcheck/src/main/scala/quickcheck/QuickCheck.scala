package quickcheck

import common._

import org.scalacheck._
import Arbitrary._
import Gen._
import Prop._

abstract class QuickCheckHeap extends Properties("Heap") with IntHeap {

  lazy val genHeap: Gen[H] = for {
    element <- arbitrary[Int]
    heap <- oneOf(const(empty), genHeap)
  } yield insert(element, heap)

  implicit lazy val arbHeap: Arbitrary[H] = Arbitrary(genHeap)

  property("min1") = forAll { a: Int =>
    val h = insert(a, empty)
    findMin(h) == a
  }

  property("gen1") = forAll { (h: H) =>
    val m = if (isEmpty(h)) 0 else findMin(h)
    findMin(insert(m, h)) == m
  }

  // If you insert any two elements into an empty heap, finding the minimum
  // of the resulting heap should get the smallest of the two elements back.
  property("ins2") = forAll { (a: Int, b: Int) =>
    val heap = insert(b, insert(a, empty))
    val min = if (a < b) a else b
    findMin(heap) == min
  }

  // If you insert an element into an empty heap, then delete the minimum,
  // the resulting heap should be empty.
  property("insEmpty") = forAll { a: Int =>
    val heap = deleteMin(insert(a, empty))
    isEmpty(heap)
  }

  // Given any heap, you should get a sorted sequence of elements 
  // when continually finding and deleting minima. 
  property("sorted") = forAll { (heap: H) =>
    def recur(h: H, l: List[Int]): List[Int] = {
      if (isEmpty(h)) l else findMin(h) :: recur(deleteMin(h), l)
    }
    val list: List[Int] = recur(heap, List())
    list == list.sorted
  }

  // Finding a minimum of the melding of any two heaps should return a minimum of one or the other.
  property("min2") = forAll { (heap1: H, heap2: H) =>
    val heap = meld(heap1, heap2)
    val min1 = findMin(heap1)
    val min2 = findMin(heap2)
    val min = if (min1 < min2) min1 else min2
    findMin(heap) == min
  }

  // For any list of n elements inserted into an empty heap,
  // deleting the smallest element n times should result in an empty heap.
  property("elements") = forAll { (xs: List[Int]) =>
    var h = empty
    for (x <- xs) h = insert(x, h)
    for (x <- xs) h = deleteMin(h)
    isEmpty(h)
  }

  // Melding 2 given heaps should be the same as melding 2 those heaps
  // where the smallest element has been moved from one heap to the other.
  property("meld") = forAll { (heap1: H, heap2: H) =>
    def toList(h: H): List[Int] = {
      def recur(h: H, list: List[Int]): List[Int] = {
        if (isEmpty(h)) list
        else findMin(h) :: recur(deleteMin(h), list)
      }
      recur(h, List())
    }
    val min = findMin(heap1)
    val heap3 = deleteMin(heap1)
    val heap4 = insert(min, heap2)
    toList(meld(heap1, heap2)) == toList(meld(heap3, heap4))
  }
}

