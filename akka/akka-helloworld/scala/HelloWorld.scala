import akka.actor.{ Actor, ActorLogging, ActorRef, ActorSystem, Props }

class HelloWorld extends Actor {
	def receive = {
		case HelloWorld.Hello => println("Hello World!");
	}
}

object HelloWorld {
	case object Hello
}


class HelloLogger extends Actor with ActorLogging {
	def receive = {
		case HelloLogger.Hello => log.info("Hello Logger!");
	}
}

object HelloLogger {
	case object Hello
}

object Main extends App {
	val actorSystem = ActorSystem("system")
	var actor = actorSystem.actorOf(Props[HelloWorld], name = "HelloWorld")

//	actor.tell(HelloWorld.Hello, ActorRef.noSender)
//	actor ! HelloWorld.Hello
	import HelloWorld.Hello
	actor ! Hello

	actor = actorSystem.actorOf(Props[HelloLogger], name = "HelloLogger")
	actor ! HelloLogger.Hello

	actorSystem.terminate();
}


