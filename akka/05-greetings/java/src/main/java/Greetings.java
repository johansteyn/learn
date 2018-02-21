import java.util.Optional;

import akka.actor.AbstractActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

// Based on Lightbend quickstart and tutorial code:
//   https://developer.lightbend.com/guides/akka-quickstart-java
//   https://doc.akka.io/docs/akka/current/guide/tutorial_1.html
// Differs in that each greeter has it's own printer, so we can see the actor hierarchy.
// And prints to stdout instead of using akka.event.Logging 
// TODO: Rather use "futures" to detect when all the work is done...
public class Greetings {
	private static final int NUMBER = 1000;
	public static int count = 0; // Used to detect when all the work is done.

	public static void main(String[] args) {
		Greetings.log("Starting Actor system...");
		ActorSystem actorSystem = ActorSystem.create("system");
		final ActorRef helloGreeter = actorSystem.actorOf(Greeter.props("Hello"), "helloGreeter");
		final ActorRef howdyGreeter = actorSystem.actorOf(Greeter.props("Howdy"), "howdyGreeter");
		final ActorRef goodDayGreeter = actorSystem.actorOf(Greeter.props("Good day"), "goodDayGreeter");
		helloGreeter.tell(new Greeter.SetName("Alice"), ActorRef.noSender());
		helloGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		howdyGreeter.tell(new Greeter.SetName("Bob"), ActorRef.noSender());
		howdyGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		howdyGreeter.tell(new Greeter.SetName("Carol"), ActorRef.noSender());
		howdyGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		goodDayGreeter.tell(new Greeter.SetName("Bob"), ActorRef.noSender());
		goodDayGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		helloGreeter.tell(new Greeter.SetName("Dave"), ActorRef.noSender());
		// We expect a greeting to Dave will fail...
		helloGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		// Akka will automatically restart the greeter, so we can retry with a different name...
		helloGreeter.tell(new Greeter.SetName("David"), ActorRef.noSender());
		helloGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		for (int i = 0; i < NUMBER; i++) {
			// We also expect a greeting to #666 to fail...
			helloGreeter.tell(new Greeter.SetName("#" + i), ActorRef.noSender());
			helloGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		}
		howdyGreeter.tell(new Greeter.SetName("Zebra"), ActorRef.noSender());
		howdyGreeter.tell(new Greeter.Greet(), ActorRef.noSender());
		// Don't terminate until all the work is done (ie. all greetings have been made)
		while (count < NUMBER - 1 + 6) {
			try {
				Thread.sleep(200);
			} catch (InterruptedException ie) {
			}
		}
		Greetings.log("Shutting down...");
		actorSystem.terminate();
		Greetings.log("Done.");
	}

	public static void log(String message) {
		System.out.println("[" + Thread.currentThread().getName() + "] " + message);
	}
}

class Greeter extends AbstractActor {
	private String greeting;
	private String name;
	private ActorRef printer;

	static public Props props(String greeting) {
		return Props.create(Greeter.class, () -> new Greeter(greeting));
	}

	static public class SetName {
		public final String name;

		public SetName(String name) {
			this.name = name;
		}
	}

	static public class Greet {
		public Greet() {
		}
	}

	public Greeter(String greeting) {
		this.greeting = greeting;
		Greetings.log("Created Greeter: " + getSelf());
	}

	@Override
	public void preStart() {
		Greetings.log("Starting '" + greeting + "' Greeter...");
	}

	@Override
	public void postStop() {
		Greetings.log("Stopped '" + greeting + "' Greeter.");
	}

	// We should not see any greeters being restarted
	// as exception only occurs in a printer...
	@Override
	public void preRestart(Throwable reason, Optional<Object> message) throws Exception {
		Greetings.log("Restarting '" + greeting + "' Greeter...");
		super.preRestart(reason, message);
	}

	@Override
	public void postRestart(Throwable reason) throws Exception {
		super.postRestart(reason);
		Greetings.log("Restarted '" + greeting + "' Greeter.");
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(SetName.class, this::onSetName)
			.match(Greet.class, this::onGreet)
			.build();
	}

	private void onSetName(SetName setName) {
		this.name = setName.name;
	}

	private void onGreet(Greet greet) {
		if (printer == null) {
			// Lazily create printer as and when needed
			printer = getContext().actorOf(Printer.props(), "printer");
		}
		if ("Dave".equals(name) || "#666".equals(name)) {
			printer.tell(new Printer.Fail(), getSelf());
		} else {
			printer.tell(new Printer.Print(greeting + ", " + name), getSelf());
		}
	}
}

class Printer extends AbstractActor {

	static public Props props() {
		return Props.create(Printer.class, () -> new Printer());
	}

	static public class Print {
		public final String message;

		public Print(String message) {
			this.message = message;
		}
	}

	static public class Fail {
	}

	public Printer() {
		Greetings.log("Created Printer: " + getSelf());
	}

	@Override
	public void preStart() {
		Greetings.log("Starting Printer...");
	}

	@Override
	public void postStop() {
		Greetings.log("Stopped Printer.");
	}

	@Override
	public void preRestart(Throwable reason, Optional<Object> message) throws Exception {
		Greetings.log("Restarting Printer...");
		super.preRestart(reason, message);
	}

	@Override
	public void postRestart(Throwable reason) throws Exception {
		super.postRestart(reason);
		Greetings.log("Restarted Printer.");
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(Print.class, this::onPrint)
			.match(Fail.class, this::onFail)
			.build();
	}

	private void onPrint(Print print) {
		Greetings.log("============ " + print.message);
		Greetings.count++;
	}

	private void onFail(Fail fail) throws Exception {
		Greetings.log("!!! Printer failing...");
		// Throwing an exception will result in the actor being stopped and restarted,
		// as well as a log entry with stack trace on stdout.
		throw new Exception("Printer Failed!");
	}
}

