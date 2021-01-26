import java.io.IOException;
import java.util.Date;
import java.util.logging.Level;
import java.util.logging.Logger;

import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.Configuration;
import com.datastax.driver.core.Host;
import com.datastax.driver.core.Metadata;
import com.datastax.driver.core.ProtocolOptions;
import com.datastax.driver.core.ResultSet;
import com.datastax.driver.core.Row;
import com.datastax.driver.core.Session;
import com.datastax.driver.core.SimpleStatement;
import com.datastax.driver.core.SocketOptions;

public class Cassandra {
  private static Logger logger = Logger.getLogger(Cassandra.class.getName());

  // Stage
  private static String address = "cass-202342972-2-306335921.stg.ciamatcher-stg.ms-df-cassandra.cdcprod5.prod.walmart.com";
  private static String user = "team";
  private static String keyspace = "ciamatcher";

  // Production
  //private static String address = "10.46.105.78";
  //private static String user = "app";
  //private static String keyspace = "cia_matcher";

  private static void log(String message) {
    logger.log(Level.INFO, message);
  }

  public static void main(String[] args) throws IOException {
    String password = args[0];
    String limit = null;
    if (args.length > 1) limit = args[1];
    int timeout = 60 * 1000; // 1 minute
    timeout *= 10; // 10 minutes
    timeout *= 6; // 1 hour
    timeout *= 2; // 2 hours

    // Cluster
    Cluster cluster = Cluster.builder()
      .addContactPoint(address)
      .withCredentials(user, password)
      .withSocketOptions(new SocketOptions().setConnectTimeoutMillis(2000))
      .build();
    cluster.init();
    System.out.println("Connected to cluster: " + cluster.getClusterName());
    System.out.println("  Configuration:");
    Configuration configuration = cluster.getConfiguration();
    System.out.println("    Protocol Options:");
    ProtocolOptions protocolOptions = configuration.getProtocolOptions();
    System.out.println("      Protocol Version: " + protocolOptions.getProtocolVersion());
    System.out.println("    Socket Options: ");
    SocketOptions socketOptions = configuration.getSocketOptions();
    System.out.println("      Read timeout: " + socketOptions.getReadTimeoutMillis());
    socketOptions.setReadTimeoutMillis(timeout);
    System.out.println("      Read timeout: " + socketOptions.getReadTimeoutMillis());
    Metadata metadata = cluster.getMetadata();
    System.out.println("  Metadata:");
    System.out.println("    Cluster name: " + metadata.getClusterName());
    System.out.println("    Hosts:");
    for (Host host : metadata.getAllHosts()) {
      System.out.printf("      Datacenter=%s, Rack=%s, Address=%s\n", host.getDatacenter(), host.getRack(), host.getAddress());
    }

    // Session
    Session session = cluster.connect(keyspace);
    Session.State state = session.getState();
    System.out.println("New session for keyspace: " + session.getLoggedKeyspace());
    System.out.println("  Connected hosts:");
    for (Host host : state.getConnectedHosts()) {
      System.out.printf("    Datacenter=%s, Rack=%s, Host=%s, # Open/Trashed Connections= %s/%s, # In-flight Queries = %s\n",
          host.getDatacenter(), host.getRack(), host.getAddress(), 
          state.getOpenConnections(host), state.getTrashedConnections(host), state.getInFlightQueries(host));
    }

    // Query
    String query = "SELECT * FROM matches";
    if (limit != null) query += " LIMIT " + limit;
    System.out.println("Creating statement for query: " + query);
    SimpleStatement simpleStatement = new SimpleStatement(query);
    System.out.println("  Read timeout: " + simpleStatement.getReadTimeoutMillis() + " milliseconds");
    System.out.println("  Setting read timeout to " + timeout + " milliseconds...");
    simpleStatement.setReadTimeoutMillis(timeout);
    System.out.println("  Read timeout: " + simpleStatement.getReadTimeoutMillis() + " milliseconds");
    log("Executing statement...");
    ResultSet resultSet = session.execute(simpleStatement);
    System.out.println("  Results:");
    int count = 0;
    try {
      for (Row row : resultSet) {
        if (count % 1000 == 0) {
          System.out.println(" " + count + " " + new Date());
          System.out.println("    " + row);
        } else {
          System.out.print(".");
        }
        count++;
      }
    } catch (Exception e) {
      System.out.println("\n Exception: " + e);
    }
    System.out.println("");
    log("" + count + " rows processed.");

    cluster.close();
  }
}

