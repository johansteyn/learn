import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.math.BigInteger;
import java.util.StringTokenizer;

// A simple, single-threaded Java wordcounter to compare performace with a multi-threaded Akka wordcounter.
public class SimpleWordCounter {
	static int totalWords;

	public static void main(String[] args) throws IOException {
		long start = System.currentTimeMillis();
		String filename = "words.txt";
			BufferedReader br = new BufferedReader(new FileReader(filename));
			String line = null;
			while ((line = br.readLine()) != null) {
//				System.out.println("Line: " + line);
				StringTokenizer st = new StringTokenizer(line);
				totalWords += st.countTokens();
fibonacci(100);
			}
			br.close();
		long end = System.currentTimeMillis();
		System.out.println("" + totalWords + " in " + (end - start) + " milliseconds");
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
