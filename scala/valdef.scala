object Main extends App {
  // For val's the RHS expression is only evaluated once: at the place where it is defined.
  val sum1 = { println("Adding 1 and 2..."); 1 + 2 }  // Evaluated here
  // For def's the RHS expression is not evaulated where it is defined.
  // Instead, it is evaluated every time it is referenced.
  def sum2 = { println("Adding 2 and 3..."); 2 + 3 }  // NOT evaluated here
  // Note that the RHS of a val can be a function literal (ie. anonymous function, or lamda)
  // However, evaluating the RHS expression simply means we are defining the anonymous function.
  // We are not calling the function, so the function body (expression) is not run when it is defined.
  // After defining the function we simply assign it to the (sum3) val so it can be called later.
  // Effectively we have given an anonymous function a name,
  // which is equivalent to using "def" to define a named function.
  // Therefore, as with a "def" it is run each time the function is called.
  val sum3 = () => { println("Adding 3 and 4..."); 3 + 4 }  // Not evaluated here either (even though we use a "val")
  println("------------------------")
  println("sum1: " + sum1) // Not re-evaluated here
  println("sum1: " + sum1) // Nor here
  println("sum2: " + sum2) // Evaluated for first time here
  println("sum2: " + sum2) // And evaluated again here
  println("sum3: " + sum3()) // Evaluated for first time here
  println("sum3: " + sum3()) // And evaluated again here
  println("------------------------")
  // Interesting...
  // In https://docs.scala-lang.org/tour/basics.html they say that 
  // the "def" keyword is used to define methods - not functions...
  // I don't quite agree with that, and prefer this explanaton:
  //   https://www.tutorialspoint.com/scala/scala_functions.htm
  // So, the only real difference between a function and a method
  // is that methods are defined inside classes/traits/objects.
  // Here is a more detailed explanation:
  //   https://stackoverflow.com/questions/2529184/difference-between-method-and-function-in-scala
}

