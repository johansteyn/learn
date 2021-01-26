import java.io.{FileReader, IOException, Reader}
import java.lang.reflect.Type
import java.util.{HashMap, List, Map}
import java.util.logging.{Level, Logger}
import scala.collection.JavaConverters._
import com.google.common.reflect.TypeToken
import com.google.gson.Gson

object Matcher extends App {
  val logger: Logger = Logger.getLogger("Matcher")
  val matchListDeserializationType: Type = new TypeToken[List[Match]](){}.getType()

  def log(message: String) = logger.log(Level.INFO, message)

  val filename = args(0)
  log("Parsing " + filename + "...")
  val gson: Gson = new Gson()
  val reader: Reader = new FileReader(filename)
  val list: List[Match] = gson.fromJson(reader, matchListDeserializationType)
  val matches = list.asScala
  for (m <- matches) {
    System.out.println(m)
  }
}

case class Match(
  compItemId: String,
  url: String,
  algoScore: List[Map[String, String]],
  verificationRequired: Int,
  primaryMatch: Int,
  pipeline: String, 
  matchStatus: String, 
  created_at: String, 
  updated_at: String, 
  miscAttributes: Map[String, String],
  isLikeMatch: Boolean
)

