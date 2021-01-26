import scala.collection.JavaConverters._
import java.util.Date
import java.util.logging.{Level, Logger}
import com.datastax.driver.core.{Cluster, Configuration, Host, Metadata, ProtocolOptions, ResultSet, Row, Session, SimpleStatement, SocketOptions}

object Cassandra extends App {
  val logger: Logger = Logger.getLogger("Cassandra")
 
  // Stage
  val address = "cass-202342972-2-306335921.stg.ciamatcher-stg.ms-df-cassandra.cdcprod5.prod.walmart.com"
  val user = "team"
  val keyspace = "ciamatcher"

  // Production
  //val address = "10.46.105.78"
  //val user = "app"
  //val keyspace = "cia_matcher"

  def log(message: String) = logger.log(Level.INFO, message)

  val password = args(0)
  val limit: String = if (args.size > 1) args(1) else null
  val timeout = 60 * 1000 * 10 * 6 * 2 // 2 hours
  
  // Cluster
  val cluster: Cluster = Cluster.builder()
      .addContactPoint(address)
      .withCredentials(user, password)
      .withSocketOptions(new SocketOptions().setConnectTimeoutMillis(2000))
      .build()
  cluster.init()
  println("Connected to cluster: " + cluster.getClusterName())
  println("  Configuration:")
  val configuration: Configuration = cluster.getConfiguration()
  println("    Protocol Options:")
  val protocolOptions: ProtocolOptions = configuration.getProtocolOptions()
  println("      Protocol Version: " + protocolOptions.getProtocolVersion())
  println("    Socket Options: ")
  val socketOptions: SocketOptions = configuration.getSocketOptions()
  println("      Read timeout: " + socketOptions.getReadTimeoutMillis())
  socketOptions.setReadTimeoutMillis(timeout)
  println("      Read timeout: " + socketOptions.getReadTimeoutMillis())
  val metadata: Metadata = cluster.getMetadata()
  println("  Metadata:")
  println("    Cluster name: " + metadata.getClusterName())
  println("    Hosts:")
  val hosts = metadata.getAllHosts().asScala
  for (host <- hosts) {
    println("      Datacenter=" + host.getDatacenter + ", Rack=" + host.getRack + ", Address=" +  host.getAddress)
  }

  // Session
  val session: Session = cluster.connect(keyspace)
  val state: Session.State = session.getState()
  println("New session for keyspace: " + session.getLoggedKeyspace())
  println("  Connected hosts:")
  val connectedHosts = state.getConnectedHosts().asScala
  for (host <- connectedHosts) {
    println("    Datacenter=" + host.getDatacenter + ", Rack=" + host.getRack + ", Address=" +  host.getAddress +
      ", # Open/Trashed Connections= " + state.getOpenConnections(host) + "/" +  state.getTrashedConnections(host) + 
      ", # In-flight Queries = " + state.getInFlightQueries(host))
  }

  // Query
  val query = "SELECT * FROM matches" + (if (limit == null) "" else " LIMIT " + limit)
//  val query = "SELECT * FROM matches where item_id = '902053494'" 
  println("Creating statement for query: " + query)
  val simpleStatement: SimpleStatement = new SimpleStatement(query)
  println("  Read timeout: " + simpleStatement.getReadTimeoutMillis() + " milliseconds")
  println("  Setting read timeout to " + timeout + " milliseconds...")
  simpleStatement.setReadTimeoutMillis(timeout)
  println("  Read timeout: " + simpleStatement.getReadTimeoutMillis() + " milliseconds")
  log("Executing statement...")
  val resultSet: ResultSet = session.execute(simpleStatement)
  println("  Results:")
  var count = 0
  try {
    val rows = resultSet.iterator().asScala
/*
    for (row <- rows) {
      if (count % 1000 == 0) {
        println(" " + count + " " + new Date())
        println("    " + row)
      } else {
        print(".")
      }
      count += 1
    }
*/
    for (row <- rows) {
      println(" " + count + " " + new Date())
      println("    " + row)
      println("")
      println("    Columns:")
      println("      item_id=" + row.getString("item_id"))
      println("      tenant_id=" + row.getString("tenant_id"))
      println("      comp_id=" + row.getString("comp_id"))
      println("      exact_match=" + row.getString("exact_match"))
      println("      is_eligible=" + row.getBool("is_eligible"))
      println("      last_match_attempt=" + row.getString("last_match_attempt"))
      println("      last_successful_match_attempt=" + row.getString("last_successful_match_attempt"))
      println("      like_match=" + row.getString("like_match"))
      println("      offer_type=" + row.getString("offer_type"))
      println("      rl1_match=" + row.getString("rl1_match"))
      println("      seller_match=" + row.getString("seller_match"))
      println("      stores_match=" + row.getString("stores_match"))
      println("      updated_at=" + row.getTimestamp("updated_at"))
      println("      variant_match=" + row.getString("variant_match"))
//      println("      =" + row.getMap("", String.class, String.class))
      println("")
      count += 1
    }
  } catch {
    case e: Exception => println("\n Exception: " + e)
  }
  println("")
  log("" + count + " rows processed.")

  cluster.close()
}

