import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Random;

import io.github.resilience4j.bulkhead.Bulkhead;
import io.github.resilience4j.bulkhead.BulkheadConfig;
import io.github.resilience4j.bulkhead.BulkheadFullException;
import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.circuitbreaker.CircuitBreakerConfig;
import io.github.resilience4j.circuitbreaker.CallNotPermittedException;


public class Worker {
  static int iterations = 100;
  static int callCount = 0;
  static int successCount = 0;
  static int failureCount = 0;
  static int blockCount = 0;
  int time = 100;
  int timeout = 3;
  int probability = 50;

  public String work() throws RuntimeException {
    callCount++;
    System.out.println("  Working in thread " + Thread.currentThread().getName() + "... (" + probability + "%)");
    Random random = new Random();
    int r = random.nextInt(100);
    if (r < probability) {
      if (r % 2 == 0) {
        throw new ImmediateException();
      } else {
        sleep(timeout * 1000);
        throw new TimeoutException();
      }
    }
    sleep(time);
    probability = probability < 100 ? probability + 1 : probability;
    return "SUCCESS";
  }

  public static void main(String[] args) {
    Worker worker = new Worker();
    try {
      if (args.length >= 2) {
        worker.time = Integer.parseInt(args[1]);
      }
      if (args.length >= 3) {
        worker.timeout = Integer.parseInt(args[2]);
      }
      if (args.length >= 4) {
        worker.probability = Integer.parseInt(args[3]);
      }
      if (args.length >= 5) {
        iterations = Integer.parseInt(args[4]);
      }
    } catch (NumberFormatException nfe) {
      System.out.println("\nInvalid number!");
      usage();
      System.exit(1);
    }
    if ("circuit-breaker".equals(args[0])) {
      runCircuitBreaker(worker);
    } else if ("bulkhead".equals(args[0])) {
      runBulkhead(worker);
    } else {
      System.out.println("\nSpecify feature, one of: circuit-breaker, bulkhead");
      usage();
      System.exit(1);
    }
    System.out.println("SUMMARY: ");
    System.out.println("  " + callCount + " calls made (" + successCount + " successes, " + failureCount + " failures)");
    System.out.println("  " + blockCount + " calls blocked");
  }

  private static void runCircuitBreaker(Worker worker) {
    System.out.println("Running circuit-breaker...");
    int iterations = 100;
    //CircuitBreaker circuitBreaker = CircuitBreaker.ofDefaults("Worker");
    //CircuitBreakerConfig config = circuitBreaker.getCircuitBreakerConfig();
    CircuitBreakerConfig config = CircuitBreakerConfig.custom()
      .failureRateThreshold(50.0f)
      .slidingWindowSize(10)
      .build();
    CircuitBreaker circuitBreaker = CircuitBreaker.of("Worker", config);
    System.out.println("Config: " + config);
		for (int i = 0; i < iterations; i++) {
			System.out.println("" + (i + 1) + "/" + iterations + " " + new Date());
			try {
        String status = circuitBreaker.executeSupplier(() -> worker.work());
        System.out.println("  " + status);
        successCount++;
      } catch (ImmediateException e) {
        System.out.println("  !!!!!!!!!!!! " + e);
        failureCount++;
      } catch (TimeoutException e) {
        System.out.println("  ............ " + e);
        failureCount++;
      } catch (CallNotPermittedException cnpe) {
        System.out.println("  ============ " + cnpe);
        blockCount++;
        int seconds = 10;
        System.out.println("  Waiting " + seconds + " seconds...");
        Random random = new Random();
        // Choose a random probability of -50%, -40%, -30% ... 0% ... 30%, 40%, 50%
        worker.probability = random.nextInt(11) * 10 - 50; 
        sleep(seconds * 1000);
        //i--; // Retry...
//        // Checking the state prevents it from changing...
//        CircuitBreaker.State state = circuitBreaker.getState();
//        while (state.equals(CircuitBreaker.State.OPEN)) {
//          int seconds = 4;
//          System.out.println("  Waiting " + seconds + " seconds...");
//          sleep(seconds * 1000);
//        }
      } catch (Exception e) {
        System.out.println("UNEXPECTED ERROR: " + e);
        failureCount++;
      }
      System.out.println("  " + circuitBreaker.getMetrics().getNumberOfSuccessfulCalls()
        + " succeeded and " + circuitBreaker.getMetrics().getNumberOfFailedCalls()
        + " failed (" + circuitBreaker.getMetrics().getFailureRate() 
        + "% failure rate) Circuit-breaker is " + circuitBreaker.getState());
    }
  }

  private static void runBulkhead(Worker worker) {
    System.out.println("Running bulkhead... ");
    //Bulkhead bulkhead = Bulkhead.ofDefaults("Worker");
    //BulkheadConfig config = bulkhead.getBulkheadConfig();
    BulkheadConfig config = BulkheadConfig.custom().maxConcurrentCalls(10).build();
    Bulkhead bulkhead = Bulkhead.of("Worker", config);
    System.out.println("Config:");
    System.out.println("  MaxConcurrentCalls=" + config.getMaxConcurrentCalls());
    System.out.println("  getMaxWaitDuration=" + config.getMaxWaitDuration());
    System.out.println("  isFairCallHandlingEnabled=" + config.isFairCallHandlingEnabled());
    System.out.println("  isWritableStackTraceEnabled=" + config.isWritableStackTraceEnabled());
    List<Thread> list = new ArrayList<Thread>();
		for (int i = 0; i < iterations; i++) {
			System.out.println("" + (i + 1) + "/" + iterations + " " + new Date());
			Thread thread = new Thread(() -> {
  			try {
          String status = bulkhead.executeSupplier(() -> worker.work());
          System.out.println("  [" + Thread.currentThread().getName() + "] " + status);
          successCount++;
        } catch (ImmediateException e) {
          System.out.println("  [" + Thread.currentThread().getName() + "] !!!!!!!!!!!! " + e);
          failureCount++;
        } catch (TimeoutException e) {
          System.out.println("  [" + Thread.currentThread().getName() + "] ............ " + e);
          failureCount++;
        } catch (BulkheadFullException bfe) {
          System.out.println("  [" + Thread.currentThread().getName() + "] ============ " + bfe);
          blockCount++;
        } catch (Exception e) {
          System.out.println("  [" + Thread.currentThread().getName() + "] UNEXPECTED ERROR: " + e);
          failureCount++;
        }
      }, "#" + (i + 1));
      System.out.println("  " + bulkhead.getMetrics().getAvailableConcurrentCalls()
        + " available concurrent calls, " + bulkhead.getMetrics().getMaxAllowedConcurrentCalls()
        + " maximum allowed concurrent calls.");
			System.out.println("  Starting thread " + thread.getName() + " at " + new Date() + "...");
      thread.start();
      list.add(thread);
			sleep(50);
		}
    try {
      for (Thread thread : list) {
        thread.join();
      }
    } catch (InterruptedException ie) {
      System.out.println("[" + Thread.currentThread().getName() + "] UNEXPECTED ERROR: " + ie);
      failureCount++;
    }
  }

  static void usage() {
    System.out.println("");
    System.out.println("USAGE:");
    System.out.println("  java -jar worker.jar circuit-breaker <time> <timeout> <probability> <iterations>");
    System.out.println("  java -jar worker.jar bulkhead <time> <timeout> <probability> <iterations>");
    System.out.println("");
    System.out.println("Where:");
    System.out.println("    time = Number of milliseconds that each job will normally take");
    System.out.println("    timeout = Number of seconds for a job to timeout");
    System.out.println("    probability = Percentage chance a job will fail (starting value, will increase by 1% with each success)");
    System.out.println("    iterations = Number of times work will be done");
    System.out.println("");
  }

  static void sleep(int millis) {
    try {
      Thread.sleep(millis);
    } catch (InterruptedException ie) {
      throw new Error("Sleep interrupted!");
    }
  }
}

class ImmediateException extends RuntimeException {
  public ImmediateException() {
    super("An exception occurred immediately!");
  }
}

class TimeoutException extends RuntimeException {
  public TimeoutException() {
    super("A timeout occurred...");
  }
}

