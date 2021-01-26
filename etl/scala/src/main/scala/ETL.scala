import java.lang.reflect.Type;
import java.util.logging.{Level, Logger}
import scala.collection.JavaConverters._
import com.datastax.driver.core.{Cluster, Configuration, Host, Metadata, ProtocolOptions, ResultSet, Row, Session, SimpleStatement, SocketOptions}
import com.google.common.reflect.TypeToken
import com.google.gson.Gson

object ETL extends App {
  val logger: Logger = Logger.getLogger("ETL")
  val matchListDeserializationType: Type = new TypeToken[java.util.List[Match]](){}.getType()

  // Stage
  val address = "cass-202342972-2-306335921.stg.ciamatcher-stg.ms-df-cassandra.cdcprod5.prod.walmart.com"
  val user = "team"
  val keyspace = "ciamatcher"

  // Production
  //val address = "10.46.105.78"
  //val user = "app"
  //val keyspace = "cia_matcher"

  def log(message: String) = logger.log(Level.INFO, message)

  val dt = args(0)
  val password = args(1)
  val limit: String = if (args.size > 2) args(2) else null
  val timeout = 60 * 1000 * 10 * 6 * 2 // 2 hours

  log("Connecting to Cassandra cluster at: " + address)
  val cluster: Cluster = Cluster.builder()
    .addContactPoint(address)
    .withCredentials(user, password)
    .withSocketOptions(new SocketOptions().setReadTimeoutMillis(timeout))
    .build()
  try {
    cluster.init()
    val session: Session = cluster.connect(keyspace)
    val query = "SELECT * FROM matches" + (if (limit == null) "" else " LIMIT " + limit)
    println("Creating statement for query: " + query)
    val simpleStatement: SimpleStatement = new SimpleStatement(query)
    log("Executing statement...")
    val resultSet: ResultSet = session.execute(simpleStatement)
    val rows = resultSet.iterator().asScala
    val itemMatches = rows.map(row => processRow(row)).filter(_ != null).toList
    println("Results:")
    for (itemMatch <- itemMatches) println("  " + itemMatch)
    log("" + itemMatches.size + " rows processed.")
  } catch {
    case e: Exception => println("\n Exception: " + e)
  } finally {
    cluster.close()
  }

  def processRow(row: Row): ItemMatch = {
    val exactMatch = row.getString("exact_match")
    val likeMatch = row.getString("like_match")
    val storesMatch = row.getString("stores_match")
    val m: (String, String) = if (exactMatch != null && exactMatch != "") ("exact", exactMatch) else
      if (likeMatch != null && likeMatch != "") ("similar", likeMatch) else
      if (storesMatch != null && storesMatch != "") ("exact", storesMatch) else (null, null)
    val matchType = m._1
    val matchJSON = m._2
    if (matchType == null || matchJSON == null) return null
    val gson: Gson = new Gson()
    val matchList: java.util.List[Match] = gson.fromJson(matchJSON, matchListDeserializationType)
    val matches = matchList.asScala.filter(m => m.primaryMatch == 1)
    if (matches.size <= 0) return null
    val primaryMatch = matches(0)
    ItemMatch(
      row.getString("tenant_id"),
      row.getString("item_id"),
      row.getString("comp_id"),
      primaryMatch.compItemId,
      primaryMatch.url,
      primaryMatch.pipeline,
      matchType,
      primaryMatch.getScore,
      primaryMatch.getMiscData,
      primaryMatch.created_at,
      primaryMatch.updated_at,
      primaryMatch.matchStatus,
      dt)
  }
}

case class ItemMatch(tenantId: String, itemId: String, compId: String, domain: String, 
  compItemId: String, compURL: String, pipeline: String, matchType: String, score: Double, 
  miscData: String,  createdAt: String, updatedAt: String, status: String, dt: String)

object ItemMatch { 
  def apply(tenantId: String, itemId: String, compId: String, 
    compItemId: String, compURL: String, pipeline: String, matchType: String, score: Double, 
    miscData: String,  createdAt: String, updatedAt: String, status: String, dt: String): ItemMatch = 
    ItemMatch(tenantId, itemId, compId, extractDomain(compURL), 
      compItemId, compURL, pipeline, matchType, score, 
      miscData,  createdAt, updatedAt, status, dt)

  // TODO: Unit test...
  def extractDomain(url: String): String = {
    if (url == null) return null;
    val start = if (url.indexOf("://") <= 0) 0 else url.indexOf("://") + 3
    val end = if (url.indexOf('/', start) <= 0) url.size else url.indexOf('/', start)
    url.substring(start, end).replaceFirst("^www.*?\\.", "")
  }
}

case class Match(compItemId: String, url: String, algoScore: java.util.List[java.util.Map[String, String]],
  primaryMatch: Int, pipeline: String, matchStatus: String, created_at: String, updated_at: String,
  miscAttributes: java.util.Map[String, String]) {

  def getScore(): Double = {
    if (algoScore == null || algoScore.size <= 0) return 0.0
    try {
      algoScore.get(0).get("score").toDouble
    } catch {
      // TODO: Do we want to handle this here, or pass it up the call stack?
      case e: NumberFormatException => 0.0
    }
  }

  def getMiscData(): String = {
    if (miscAttributes == null || miscAttributes.size <= 0) return null
    // Scala's toString method will produce somethimg like:
    //   {conversion_factor=1, is_exact_match=true, source=MST}
    // Which we need to convert to:
    //   {"conversion_factor":"3.2", "is_exact_match":"false", "source":"MST"}
    miscAttributes.toString.replaceAll("\\{", "{\"").replaceAll("=", "\":\"").replaceAll(", ", "\", \"").replaceAll("}", "\"}")
  }
}

