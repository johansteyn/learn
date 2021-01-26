import java.io.File;

import org.apache.avro.file.DataFileReader;
import org.apache.avro.file.DataFileWriter;
import org.apache.avro.io.DatumReader;
import org.apache.avro.io.DatumWriter;
import org.apache.avro.specific.SpecificDatumReader;
import org.apache.avro.specific.SpecificDatumWriter;

import steyn.johan.User;

// By default this app will write (serialize) and read (deserialize) to the same file.
// Specifying "readonly" will skip the write, allowing you to:
// - Do an initial write + read
// - Modify the schema
// - Do a read only, to verify that the schema change was compatible 
public class Avro {
  public static void main(String[] args) throws Exception {
    System.out.println("============ Avro ============");

    File file = new File("users.avro");
    if (!(args.length > 0 && "readonly".equals(args[0]))) {
      System.out.println("Constructing User instances...");
      // Three ways to create an instances...
      // 1. Default constructor + setters
      User alice = new User();
      alice.setName("Alice");
      alice.setLuckyNumber(2);
      // Let favourite colour default to null

      // 2. Constructor with parameters
      //User bob = new User("Bob", 7, "Red");
      //User bob = new User("Bob", 7, "Red", "bob@builder.com");
      User bob = new User("Bob", 7, "bob@builder.com");

      // 3. Builder
      User carol = User.newBuilder()
        .setName("Carol")
        .setLuckyNumber(null)
        //.setFavouriteColour("blue")
        .build();
      // Lucky number won't default to null - need to set it explicitly
      // Note that builders will validate fields as they are set, whereas constructors will not.
      // But then instances created via constructors could fail when serialized...

      System.out.println("Serializing User instances to file...");
      DatumWriter<User> userDatumWriter = new SpecificDatumWriter<User>(User.class);
      DataFileWriter<User> dataFileWriter = new DataFileWriter<User>(userDatumWriter);
      dataFileWriter.create(alice.getSchema(), file);
      dataFileWriter.append(alice);
      dataFileWriter.append(bob);
      dataFileWriter.append(carol);
      dataFileWriter.close();
    }

    // Deserialize user objects from file
    System.out.println("Deserializing User instances from file...");
    DatumReader<User> userDatumReader = new SpecificDatumReader<User>(User.class);
    DataFileReader<User> dataFileReader = new DataFileReader<User>(file, userDatumReader);
    User user = null;
    for (int i = 1; dataFileReader.hasNext(); i++) {
      // Reusing the user instance reduces garbage collection
      user = dataFileReader.next(user);
      System.out.println("User #" + i + ": " + user);
    }
  }
}
