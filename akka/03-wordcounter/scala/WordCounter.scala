import scala.concurrent.duration._

import akka.actor.{ Actor, ActorLogging, ActorRef, ActorSystem, Props }
import akka.dispatch.ExecutionContexts._
import akka.pattern.ask
import akka.util.Timeout

// Can use this instead of declaring an implicit execution context
//import scala.concurrent.ExecutionContext.Implicits.global

// Based on:
//   https://www.toptal.com/scala/concurrency-and-fault-tolerance-made-easy-an-intro-to-akka
// NOTE: Messages are the only thing that actors share, so they need to be immutable
//       Case classes are immutable by default, and they also work with pattern matching.
object Main extends App {
	implicit val executionContext = global // Need this for the "future.map" method (or, see import above)
	implicit val timeout = Timeout(25 seconds)	// Need this for the "ask" method
	val actorSystem = ActorSystem("system")
	var wordCounter = actorSystem.actorOf(Props[WordCounter], name = "wordcounter")
	val future = wordCounter ? WordCounter.Count("words.txt") // The "ask"(?) method returns a Future instance
	future.map { 
		totalWords => println("Total number of words: " + totalWords)
		actorSystem.terminate
	}
}

class WordCounter extends Actor {
	private var fileSender: Option[ActorRef] = None
	private var linesRead = 0
	private var linesCounted = 0
	private var totalWords = 0

	def receive = {
		case WordCounter.Count(filename: String) => {
			fileSender = Some(sender) // Save reference to main app
			import scala.io.Source._
			fromFile(filename).getLines.foreach {
				// Create a separate child actor for each line and send it a message (to count the words in that line)
				// Being asynchronous means that any number of lines can be counted concurrently
				// ie. no need to wait for each line to be counted before reading the next line.
				line => context.actorOf(Props[Counter]) ! Counter.Count(line)
				linesRead += 1
			}
		}
		case WordCounter.Counted(numWords) => {
			totalWords += numWords
			linesCounted += 1
			if (linesCounted == linesRead) {
				// All the lines have been counted, so send the result back to the main app
				fileSender.map(_ ! totalWords)
			}
		}
		case _ => {
			println("Message not recognized!")
		}
	}
}

object WordCounter {
	case class Count(filename: String)		// Sent by main method to start counting words in the specified file
	case class Counted(numWords: Integer)	// Sent by child counters to return the number of words in a line
}

class Counter extends Actor {
	def receive = {
		case Counter.Count(string) => {
			val numWords = string.split(" ").length
			// Send the number of words in the line back to the parent actor (ie. the "sender", which is the WordCounter actor)
			sender ! WordCounter.Counted(numWords)
		}
		case _ => println("Error: message not recognized")
	}
}

object Counter {
	case class Count(string: String)	// Sent by parent counter to count the words in a line
}

