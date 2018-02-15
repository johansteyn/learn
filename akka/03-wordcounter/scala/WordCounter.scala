import scala.concurrent.duration._
import akka.actor.{ Actor, ActorLogging, ActorRef, ActorSystem, Props }
import akka.dispatch.ExecutionContexts._
import akka.pattern.ask
import akka.util.Timeout

object Main extends App {
	implicit val ec = global // TODO: What is this?
	implicit val timeout = Timeout(25 seconds)	// TODO: And this?
	val actorSystem = ActorSystem("system")
	var wordCounter = actorSystem.actorOf(Props[WordCounter], name = "wordcounter")
	//val future = wordCounter ? "words.txt" // Unrecognized message
	val future = wordCounter ? WordCounter.Count("words.txt") // The "ask"(?) method returns a Future instance
	future.map { 
		totalWords => println("Total number of words " + totalWords)
		actorSystem.terminate
	}
}

class WordCounter extends Actor {
	private var totalWords = 0
	private var totalLines = 0
	private var linesProcessed = 0
	private var fileSender: Option[ActorRef] = None

	def receive = {
		case WordCounter.Count(filename: String) => {
			fileSender = Some(sender) // Save reference to process invoker
			import scala.io.Source._
			fromFile(filename).getLines.foreach {
				// Create a separate child actor for each line and send it a message containing the line
				// NOTE: This measn that any number of lines can be counted concurrently
				line => context.actorOf(Props[Counter]) ! Counter.Count(line)
				totalLines += 1
			}
		}
		case WordCounter.Counted(numWords) => {
			totalWords += numWords
			linesProcessed += 1
			if (linesProcessed == totalLines) {
				// All the lines have been counted, so send the result back to the process invoker
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
	case class Count(string: String)
}

