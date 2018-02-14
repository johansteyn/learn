import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Executors;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Future;

// A multi-threaded Fibonacci calculator
public class MultiFibonacci {
	public static void main(String[] args) throws Exception {
		int number = Integer.parseInt(args[0]);
		int threadPoolSize = Integer.parseInt(args[1]);
		ExecutorService threadPool = Executors.newFixedThreadPool(threadPoolSize);
		List<Future> futures = new ArrayList<Future>();
		long start = System.currentTimeMillis();
		for (int i = 1; i <= number; i++) {
			Worker worker = new Worker(i);
			futures.add(threadPool.submit(worker));
		}
		for (Future future : futures) {
			future.get();
		}
		long end = System.currentTimeMillis();
		System.out.println("" + (end - start) + " milliseconds");
		System.exit(0);
	}

	static class Worker implements Runnable {
		int nth;

		public Worker(int nth) {
			this.nth = nth;
		}

		public void run() {
			Fibonacci.calculate(nth);
		}
	}
}

