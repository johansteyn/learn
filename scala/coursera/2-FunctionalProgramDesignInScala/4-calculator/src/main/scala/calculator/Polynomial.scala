package calculator

object Polynomial {
  def computeDelta(a: Signal[Double], b: Signal[Double], c: Signal[Double]): Signal[Double] = {
    Signal { b() * b() - 4 * a() * c() }
  }

  def computeSolutions(a: Signal[Double], b: Signal[Double], c: Signal[Double], delta: Signal[Double]): Signal[Set[Double]] = {
    Signal {
      delta() match {
        case d if (d > 0) => {
          val x = Math.sqrt(delta())
          val root1 = (-b() + x) / (2 * a())
          val root2 = (-b() - x) / (2 * a())
          Set(root1, root2)
        }
        case d if (d == 0) => Set(-b() / (2 * a()))
        case _ => Set()
      }
    }
  }
}

