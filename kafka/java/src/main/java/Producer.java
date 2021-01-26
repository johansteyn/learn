import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Properties;

import org.apache.kafka.clients.producer.Callback;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;


public class Producer {
  private static int count;

  public static void main(String[] args) throws Exception {
    System.out.println("**** Kafka Producer ***");

    boolean async = false;
    if (args.length > 0 && "a".equals(args[0])) {
      async = true;
    }

    System.out.println("Creating properties...");
    Properties properties = new Properties();
    // The first three proerties are required
    properties.put("bootstrap.servers", "localhost:9092");
    properties.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
    properties.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");
    // All other properties are optional
    // These timeouts don't seem to have any effect when the kafka broker is down...
    properties.put("timeout.ms", 1000);
    properties.put("request.timeout.ms", 1000);
    properties.put("metadatafetch.timeout.ms", 1000);


    System.out.println("Creating producer...");
    KafkaProducer producer = new KafkaProducer(properties);

    System.out.println("Populating users...");
    List<String> users = new ArrayList<String>();
    users.add("Alice");
    users.add("Bob");
    users.add("Carol");
    users.add("David");
    users.add("Eva");
    users.add("Fred");
    users.add("George");
    users.add("Harry");
    users.add("Ian");
    users.add("Jack");
    users.add("Kevin");
    users.add("Lucy");
    users.add("Max");
    users.add("Nancy");
    users.add("Oscar");
    users.add("Peter");
    users.add("Quinn");
    users.add("Richard");
    users.add("Sam");
    users.add("Tim");
    users.add("Ursula");
    users.add("Victoria");
    users.add("Wayne");
    users.add("Xenia");
    users.add("Yoshi");
    users.add("Zak");

    System.out.println("Producing and sending user records...");
    for (String user : users) {
      ProducerRecord record = new ProducerRecord("users", user);
      try {
        System.out.println("  Sending message for user '" + user + "'...");
        if (!async) {
          RecordMetadata result = (RecordMetadata) producer.send(record).get();
          onSuccess(result);
          Thread.sleep(1000);
        } else {
          producer.send(record, new ProducerCallback(user));
        }
      } catch (Exception e) {
        onError(user, e);
      }
    }
    if (async) {
      System.out.println("Waiting for async sends to complete...");
      while (count < 26) {
        Thread.sleep(1000);
      }
    }
    System.out.println("Done.");

    // TODO: Avro serialization...
  }

  public static void onSuccess(RecordMetadata result) {
    count++;
    System.out.println("  #" + count + " @ " + new Date(result.timestamp()) + ": partition=" + result.partition() + ", offset=" + result.offset());
  }

  public static void onError(String user, Exception e) {
    System.out.println("Error sending record for user '" + user + "'");
    e.printStackTrace();
  }
}

class ProducerCallback implements Callback {
  private String user;

  public ProducerCallback(String user) {
    this.user = user;
  }

  @Override
  public void onCompletion(RecordMetadata result, Exception e) {
    if (e == null) {
      Producer.onSuccess(result);
    } else {
      Producer.onError(user, e);
    }
  }
}

