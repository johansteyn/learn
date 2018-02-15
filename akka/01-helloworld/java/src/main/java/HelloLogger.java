import akka.actor.AbstractLoggingActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

public class HelloLogger extends AbstractLoggingActor {
	public static class Hello {
	}

	public static Props props() {
		return Props.create(HelloLogger.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(Hello.class, this::onHello).build();
	}

	private void onHello(Hello hello) {
		log().info("Hello Logger!");
	}

	public static void main(String[] args) {
		ActorSystem actorSystem = ActorSystem.create("system");
		final ActorRef actor = actorSystem.actorOf(HelloLogger.props(), "HelloLogger");
		actor.tell(new Hello(), ActorRef.noSender());
		actorSystem.terminate();
	}
}


