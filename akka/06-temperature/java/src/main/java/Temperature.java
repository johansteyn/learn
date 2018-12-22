import java.util.Optional;

import akka.actor.AbstractActor;
import akka.actor.AbstractActor.Receive;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

// Based on Lightbend tutorial code:
//   https://doc.akka.io/docs/akka/current/guide/tutorial_3.html
public class Temperature {
	public static void main(String[] args) {
		Temperature.log("Starting Actor system...");
		ActorSystem actorSystem = ActorSystem.create("system");
		ActorRef device = actorSystem.actorOf(Device.props("group", "device"));
//		device.tell(new Device.RecordTemperature(1L, 24.0), ActorRef.noSender());
//		device.tell(new Device.ReadTemperature(2L), ActorRef.noSender());
		device.tell(new Device.RecordTemperature(1L, 24.0), device);
		device.tell(new Device.ReadTemperature(2L), device);

		try {
			Thread.sleep(2000);
		} catch (InterruptedException ie) {
		}
		log("Shutting down...");
		actorSystem.terminate();
		log("Done.");
	}

	public static void log(String message) {
		System.out.println("[" + Thread.currentThread().getName() + "] " + message);
	}
}

class Device extends AbstractActor {
	final String groupId;
	final String deviceId;

	public Device(String groupId, String deviceId) {
		this.groupId = groupId;
		this.deviceId = deviceId;
	}

	public static Props props(String groupId, String deviceId) {
		return Props.create(Device.class, groupId, deviceId);
	}

	public static final class RecordTemperature {
		final long requestId;
		final double value;

		public RecordTemperature(long requestId, double value) {
			this.requestId = requestId;
			this.value = value;
		}
	}

	public static final class TemperatureRecorded {
		final long requestId;

		public TemperatureRecorded(long requestId) {
			this.requestId = requestId;
		}
	}

	public static final class ReadTemperature {
		final long requestId;

		public ReadTemperature(long requestId) {
			this.requestId = requestId;
		}
	}

	public static final class RespondTemperature {
		final long requestId;
		final Optional<Double> value;

		public RespondTemperature(long requestId, Optional<Double> value) {
			this.requestId = requestId;
			this.value = value;
		}
	}

	Optional<Double> lastTemperatureReading = Optional.empty();

	@Override
	public void preStart() {
		Temperature.log("Device '" + groupId + "-" + deviceId + "' starting...");
	}

	@Override
	public void postStop() {
		Temperature.log("Device '" + groupId + "-" + deviceId + "' stopped.");
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(RecordTemperature.class, r -> {
			Temperature.log("Recorded temperature reading for request " + r.requestId + ": " + r.value);
			lastTemperatureReading = Optional.of(r.value);
			getSender().tell(new TemperatureRecorded(r.requestId), getSelf());
		}).match(ReadTemperature.class, r -> {
			getSender().tell(new RespondTemperature(r.requestId, lastTemperatureReading), getSelf());
		}).build();
	}
}
/*
class Greeter extends AbstractActor {
	static public Props props(String greeting) {
		return Props.create(Greeter.class, () -> new Greeter(greeting));
	}

	static public class Greet {
		public Greet() {
		}
	}

	public Greeter(String greeting) {
		this.greeting = greeting;
		Temperature.log("Created Greeter: " + getSelf());
	}

	@Override
	public void preStart() {
		Temperature.log("Starting '" + greeting + "' Greeter...");
	}

	@Override
	public void postStop() {
		Temperature.log("Stopped '" + greeting + "' Greeter.");
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(Greet.class, this::onGreet)
			.build();
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
*/

