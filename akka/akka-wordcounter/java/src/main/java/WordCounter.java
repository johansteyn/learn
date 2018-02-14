import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.math.BigInteger;
import java.util.StringTokenizer;

import akka.actor.AbstractActor;
import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;

public class WordCounter extends AbstractActor {
	static ActorSystem actorSystem = ActorSystem.create("system");
	static long start;
	static int totalWords;
	static int totalLines;
	static int linesProcessed;

	public static void main(String[] args) {
		final ActorRef wordcounter = actorSystem.actorOf(WordCounter.props(), "wordcounter");
		wordcounter.tell(new WordCounter.Count("words.txt"), wordcounter);
	}

	public static class Count {
		String filename;

		public Count(String filename) {
			this.filename = filename;
		}
	}

	public static class Counted {
		int numWords;

		public Counted(int numWords) {
			this.numWords = numWords;
		}
	}

	public static Props props() {
		return Props.create(WordCounter.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder()
			.match(Count.class, this::count)
			.match(Counted.class, this::counted)
			.build();
		}

	private void count(Count count) throws IOException {
//		System.out.println("Count(filename)");
		start = System.currentTimeMillis();
		BufferedReader br = new BufferedReader(new FileReader(count.filename));
		String line = null;
		while ((line = br.readLine()) != null) {
//			System.out.println("Line: " + line);
			totalLines++;
			ActorRef counter = actorSystem.actorOf(Counter.props());
			counter.tell(new Counter.Count(line), self());
		}
		br.close();
	}

	private void counted(Counted counted) throws IOException {
//		System.out.println("Counted " + counted.numWords + " words");
		totalWords += counted.numWords;
		linesProcessed++;
		if (linesProcessed == totalLines) {
			// All the lines have been counted
			long end = System.currentTimeMillis();
			System.out.println("" + totalWords + " in " + (end - start) + " milliseconds");
			actorSystem.terminate();
		}
	}
}

class Counter extends AbstractActor {
	public static class Count {
		String line;

		public Count(String line) {
			this.line = line;
		}
	}

	public static Props props() {
		return Props.create(Counter.class);
	}

	@Override
	public Receive createReceive() {
		return receiveBuilder().match(Count.class, this::count).build();
	}

	private void count(Count count) {
//		System.out.println("Counting words in line: " + count.line);
		StringTokenizer st = new StringTokenizer(count.line);
		int numWords = st.countTokens();
fibonacci(100);
		sender().tell(new WordCounter.Counted(numWords), self());
	}

	// CPU intensive method to calculate the nth Fibonacci number
	static BigInteger fibonacci(long nth){
		nth = nth - 1;
		long count = 0;
		BigInteger first = BigInteger.ZERO;
		BigInteger second = BigInteger.ONE;
		BigInteger third = null;
		while (count < nth){
			third = new BigInteger(first.add(second).toString());
			first = new BigInteger(second.toString());
			second = new BigInteger(third.toString());
			count++;
		}
		return third;
	}
}

