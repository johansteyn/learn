import akka.actor.{ Actor, ActorLogging, ActorRef, ActorSystem, Props }

object Main extends App {
	val actorSystem = ActorSystem("system")
	var ping = actorSystem.actorOf(Props[Ping], name = "ping")
	var pong = actorSystem.actorOf(Props[Pong], name = "pong")
	ping.tell(Ping.Message, pong)
}

class Ping extends Actor {
	def receive = {
		case Ping.Message => 
			println("Ping...");
			Thread.sleep(200);
			sender.tell(Pong.Message, self)
	}
}

object Ping {
	case object Message
}

class Pong extends Actor {
	def receive = {
		case Pong.Message => 
			Pong.counter += 1
			println("Pong! [" + Pong.counter + "]");
			if (Pong.counter >= 42) {
				Main.actorSystem.terminate();
			} else {
				Thread.sleep(200);
				sender.tell(Ping.Message, self)
			}
	}
}

object Pong {
	var counter = 0
	case object Message
}

