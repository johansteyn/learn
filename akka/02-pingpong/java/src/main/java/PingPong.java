import akka.actor.AbstractActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

public class PingPong {
	static ActorSystem actorSystem = ActorSystem.create("system");

	public static void main(String[] args) {
		final ActorRef ping = actorSystem.actorOf(Ping.props(), "ping");
		final ActorRef pong = actorSystem.actorOf(Pong.props(), "pong");
		ping.tell(new Ping.Message(), pong);
	}

	static void pause(long millis) {
		try {
			Thread.sleep(millis);
		} catch (InterruptedException ie) {
		}
	}
}

class Ping extends AbstractActor {
	public static class Message {
	}

	public static Props props() {
		return Props.create(Ping.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(Message.class, this::ping).build();
	}

	private void ping(Message message) {
		System.out.println("Ping...");
		PingPong.pause(200);
		sender().tell(new Pong.Message(), self());
	}
}

class Pong extends AbstractActor {
	private static int counter;

	public static class Message {
	}

	public static Props props() {
		return Props.create(Pong.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(Message.class, this::pong).build();
	}

	private void pong(Message message) {
		counter++;
		System.out.println("Pong! [" + counter + "]");
		if (counter >= 42) {
			PingPong.actorSystem.terminate();
			return;
		}
		PingPong.pause(200);
		sender().tell(new Ping.Message(), self());
	}
}

