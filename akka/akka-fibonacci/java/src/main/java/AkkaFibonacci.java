import java.util.ArrayList;
import java.util.List;
import java.math.BigInteger;

import akka.actor.AbstractActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

public class AkkaFibonacci extends AbstractActor {
	static ActorSystem actorSystem = ActorSystem.create("system");
	static int number;
	static long start;
	static int calculated;

	public static void main(String[] args) {
		number = Integer.parseInt(args[0]);
		ActorRef fibonacci = actorSystem.actorOf(AkkaFibonacci.props(), "fibonacci");
		List<ActorRef> calculators = new ArrayList<ActorRef>();
		for (int i = 1; i <= number; i++) {
		 	ActorRef calculator = actorSystem.actorOf(Calculator.props());
			calculators.add(calculator);
		}
		start = System.currentTimeMillis();
		for (int i = 1; i <= number; i++) {
			ActorRef calculator = calculators.get(i - 1);
			calculator.tell(new Calculator.Calculate(i), fibonacci);
		}
	}

	public static class Done {
	}

	public static Props props() {
		return Props.create(AkkaFibonacci.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(Done.class, this::done)
			.build();
		}

	private void done(Done done) {
		calculated++;
		if (calculated == number) {
			long end = System.currentTimeMillis();
			System.out.println("" + (end - start) + " milliseconds");
			actorSystem.terminate();
		}
	}
}

class Calculator extends AbstractActor {
	public static class Calculate {
		int nth;

		public Calculate(int nth) {
			this.nth = nth;
		}
	}

	public static Props props() {
		return Props.create(Calculator.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(Calculate.class, this::calculate)
			.build();
		}

	private void calculate(Calculate calculate) {
		Fibonacci.calculate(calculate.nth);
		sender().tell(new AkkaFibonacci.Done(), self());
	}
}

