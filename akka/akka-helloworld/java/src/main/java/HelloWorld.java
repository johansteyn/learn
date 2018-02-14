import akka.actor.AbstractActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

public class HelloWorld extends AbstractActor {
	public static class Hello {
		// Definition of a message that this Actor supports
		// Actors do not expose public fields or methods (state & behaviour)
		// Instead, they expose public messages
	}

	public static Props props() {
		return Props.create(HelloWorld.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(Hello.class, this::onHello).build();
	}

	private void onHello(Hello hello) {
		// The method that will be invoked when a "Hello" message is received.
		System.out.println("Hello World!");
	}

	public static void main(String[] args) {
		// Create an Actor system - the runtime that starts and maintains thread pools
		ActorSystem actorSystem = ActorSystem.create("system");

		// Obtain a reference to an Actor, which is not a direct reference to the instance itself.
		// Always give your actor a name, so that you can recuperate a reference to it by name if you need to.
		// (Very expensive, so only do it if absolutely necessary, eg: lost a reference somehow)
		final ActorRef actor = actorSystem.actorOf(HelloWorld.props(), "HelloWorld");
		
		// Send a message to our actor
		// Note that we need to tell the actor which actor is sending the message,
		// but in this case we only have one actor and we are sending from our main method.
		actor.tell(new Hello(), ActorRef.noSender());

		// Quit the Actor system
		actorSystem.terminate();
	}
}


