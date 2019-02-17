// From the "Traits" section in:
// https://docs.scala-lang.org/tutorials/scala-for-java-programmers.html

trait Ord {
  def < (that: Any): Boolean
  def <=(that: Any): Boolean = (this < that) || (this == that)
  def > (that: Any): Boolean = !(this <= that)
  def >=(that: Any): Boolean = !(this < that)
}

// This is the "normal" class they use
class Date(y: Int, m: Int, d: Int) extends Ord {
  def year = y
  def month = m
  def day = d
  override def toString(): String = year + "-" + month + "-" + day

  override def equals(that: Any): Boolean =
    that.isInstanceOf[Date] && {
      val o = that.asInstanceOf[Date]
      o.day == day && o.month == month && o.year == year
    }

  def <(that: Any): Boolean = {
    if (!that.isInstanceOf[Date])
      sys.error("cannot compare " + that + " and a Date")
    val o = that.asInstanceOf[Date]
    (year < o.year) || (year == o.year && (month < o.month || (month == o.month && day < o.day)))
  }
}

// This is one I made, using a "case" class and pattrn matching
case class MyDate(y: Int, m: Int, d: Int) extends Ord {
  def <(that: Any): Boolean = that match {
    case MyDate(year, month, day) => (y < year) || (y == year && (m < month || (m == month && d < day)))
    case _ => sys.error("cannot compare " + that + " and a MyDate")
  }
}

object Main extends App {
  val date1 = new Date(1967, 7, 15)
  val date2 = new Date(1969, 7, 14)
  println("" + date1 + " < " + date2 + " ? " + (date1 < date2))
  println("" + date2 + " < " + date1 + " ? " + (date2 < date1))
  val mydate1 = MyDate(1967, 7, 15)
  val mydate2 = MyDate(1969, 7, 14)
  println("" + mydate1 + " < " + mydate2 + " ? " + (mydate1 < mydate2))
  println("" + mydate2 + " < " + mydate1 + " ? " + (mydate2 < mydate1))
}

