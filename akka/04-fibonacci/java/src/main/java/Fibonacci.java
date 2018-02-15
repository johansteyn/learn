import java.math.BigInteger;

// A simple, single-threaded Fibonacci calculator
public class Fibonacci {
	public static void main(String[] args) throws Exception {
		int number = Integer.parseInt(args[0]);		
		long start = System.currentTimeMillis();
		for (int i = 1; i <= number; i++) {
			Fibonacci.calculate(i);
		}
		long end = System.currentTimeMillis();
		System.out.println("" + (end - start) + " milliseconds");
	}

	// Calculate the nth Fibonacci number
	public static BigInteger calculate(long nth) {
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

