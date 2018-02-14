// Similarities between traits and classes:
//   Can have abstract and concrete fields and methods.
//   Can have a super (parent) class/trait.

// Differences between traits and classes:
//   Traits can have primary constructors but no auxiliary constructors
//   And a trait's primary constructor cannot have parameters
//   Traits can extend either classes or traits, 
//     but only classes with parameterless constructors
//   A class can extend only one class but any number of traits
//   An object can have a trait added at contruction time.

// NOTE: The println statements are to show the construction orders

trait Logger {
	println("--- Logger")

	// An abstract method in a trait does not require the "abstract" keyword
	def log(message: String)

	// A trait can also have concrete methods (ie. like abstract classes in Java)
	def debug(message: String) { 
		log("DEBUG: " + message) 
	}

	def info(message: String) { 
		log("INFO: " + message) 
	}

	def warn(message: String) { 
		log("WARN: " + message) 
	}

	def error(message: String) { 
		log("INFO: " + message) 
	}
}


trait ConsoleLogger extends Logger {
	println("--- ConsoleLogger")
	// A concrete implementation does not require the "override" keyword when implementing an abstract method
	def log(message: String) { 
		println(message)
	}
}

trait TimestampLogger extends ConsoleLogger {
	println("--- TimestampLogger")
	// We need the "override" keyword when overriding a method that is already implemented (in a parent class)
	override def log(message: String) {
		super.log(s"${java.time.Instant.now()} $message")
	}
}

trait ShortLogger extends ConsoleLogger {
	println("--- ShortLogger")
	// A trait can have concrete fields
	var maxLength = 20
	override def log(message: String) {
		super.log(if (message.length <= maxLength) message else s"${message.substring(0, maxLength)}...")
	}
}

//trait FileLogger extends Logger {
trait FileLogger extends ConsoleLogger {
	println("--- FileLogger")
	import java.io.{FileWriter, PrintWriter}
	val pw = new PrintWriter(new FileWriter("loggers.log", true))

//	def log(message: String) {
	override def log(message: String) {
		Main.pw.println(message);
		Main.pw.flush()
	}
}

trait Singer {
	println("--- Singer")
	def sing
}

trait Marker {
	// Empty trait to illustrate construction order
	println("--- Marker")
}

class Fruit extends Marker {
	println("=== Fruit")
}

// A class can have traits
class Bananas extends Fruit with Singer with ConsoleLogger {
	println("=== Bananas")
	def sing {
		log("Yes, we have no bananas. We have no bananas today.");
	}
}

trait CountLogger extends Logger {
	println("--- CountLogger")
	// A trait can also have abstract fields
	var count: Int
	// But if a method uses an abstract field, then it must be abstract
	abstract override def log(message: String) {
		count += 1
		super.log(s"[$count] $message")
	}
}

// Here we don't extend Fruit (just to see how/if construction order is affected)
class Pineapple (number: Int = 0) extends Singer with ConsoleLogger with CountLogger {
	println("=== Pineapple")
	// Make the abstract field concrete
	var count = number
	def sing {
		log("Agadoo-doo-doo, push pineapple, shake the tree!");
	}
}

// Just as a class can extend a trait, a trait extend a class.
// BUT... only certain classes can use that trait:
//  - A class that doesn't extend any other class (it will end up extending the trait's class)
//  - A class that extends the same class or a subclass of the one tehtrait extends
trait LoggedException extends Exception with ConsoleLogger {
	println("--- LoggedException")
	def log() { 
		log(getMessage())	// The getMessage is inherited from Exception
	}
}

class ArrghException extends LoggedException {
	println("=== ArrghException")
	override def getMessage() = "Arrgh!"
}

import java.io.IOException
class UffException extends IOException with LoggedException {
	println("=== UffException")
	override def getMessage() = "Uff!"
}

// An object can have traits
object Main extends App with FileLogger with TimestampLogger {
	log("================================================");
	debug("Debug message");
	info("Informational message");
	warn("Warning message");
	error("Error message");

	println("------------------------------------------------");
	var bananas = new Bananas()
	bananas.sing

	println("------------------------------------------------");
	// An object can have extra traits mixed in when constructed
	bananas = new Bananas() with TimestampLogger
	bananas.sing
	bananas = new Bananas() with ShortLogger
	bananas.sing
	bananas = new Bananas() with ShortLogger with TimestampLogger
	bananas.sing

	println("------------------------------------------------");
	// An abstract field can be made concrete with a default value...
	var pineapple = new Pineapple() with CountLogger
	pineapple.sing
	pineapple.sing
	pineapple.sing
	println("------------------------------------------------");
	// ...or a specified value
	pineapple = new Pineapple(7) with CountLogger
	pineapple.sing
	pineapple.sing
	pineapple.sing

	println("------------------------------------------------");
	// We can change the value of a field in a trait
	bananas = new Bananas() with TimestampLogger with ShortLogger
	bananas.asInstanceOf[ShortLogger].maxLength = 35
	bananas.sing

	println("------------------------------------------------");
	// Even though TimestampLogger extends ConsoleLogger, 
	// this uses FileLogger due to its position in the list
	val b = new Bananas() with FileLogger with TimestampLogger
	b.sing

	println("------------------------------------------------");
	val arrgh = new ArrghException()
	arrgh.log()
	val uff = new UffException()
	uff.log()
}
/*
Construction order for last "Bananas" instance, called "b":

1. Marker trait
2. Fruit class
3. Singer trait
4. Logger trait (created before ConsoleLogger sub-trait)
5. ConsoleLogger trait
6. Bananas class (created before super class and all superclass traits, but before its own traits)
7. FileLogger trait (Logger super-trait already created)
8. TimestampLogger trait (both Logger and ConsoleLogger super-traits already created)

Class/trait hierarchy:
                               Logger
                               /   |
    Fruit-------Marker        /    |
      |                      /     |
  Bananas-------Singer--------ConsoleLogger
      |                    /       |
      |                   /        |
      |                  /         |
      |                 /          |
      |                /           |
      b-------FileLogger-----TimestampLogger
*/
