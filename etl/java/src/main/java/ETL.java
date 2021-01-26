import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

// Cassandra
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

// JSON
import com.google.common.reflect.TypeToken;
import com.google.gson.Gson;

public class ETL {
  private static Logger logger = Logger.getLogger(ETL.class.getName());
  private static Type matchListDeserializationType = new TypeToken<List<Match>>(){}.getType();

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

  public static void main(String[] args) {
    String dt = args[0];
    String password = args[1];
    String limit = null;
    if (args.length > 2) limit = args[2];
    int timeout = 60 * 1000; // 1 minute
    timeout *= 10; // 10 minutes
    timeout *= 6; // 1 hour
    timeout *= 2; // 2 hours

    log("Connecting to Cassandra cluster at: " + address);
    Cluster cluster = Cluster.builder()
      .addContactPoint(address)
      .withCredentials(user, password)
      .withSocketOptions(new SocketOptions().setReadTimeoutMillis(timeout))
      .build();
    cluster.init();
    Session session = cluster.connect(keyspace);
    String query = "SELECT * FROM matches";
    if (limit != null) query += " LIMIT " + limit;
    log("Creating statement for query: " + query);
    SimpleStatement simpleStatement = new SimpleStatement(query);
    log("Executing statement...");
    ResultSet resultSet = session.execute(simpleStatement);
    System.out.println("  Results:");
    Gson gson = new Gson();
    int rowCount = 0;
    int matchCount = 0;
    try {
      for (Row row : resultSet) {
        rowCount++;
        System.out.println("\nRow #" + rowCount);
        //System.out.println("" + row);
        String matchType = null;
        String matchJSON = row.getString("exact_match");
        if (matchJSON != null && !"".equals(matchJSON)) {
          matchType = "exact";
        } else {
          matchJSON = row.getString("like_match");
          if (matchJSON != null && !"".equals(matchJSON)) {
            matchType = "similar";
          } else {
            matchJSON = row.getString("stores_match");
            if (matchJSON != null && !"".equals(matchJSON)) {
              matchType = "exact";
            } else {
              // No match found
              continue;
            }
          }
        }
        // Only consider a match whose "primaryMatch" value is 1
        Match match = null;
        List<Match> matches = gson.fromJson(matchJSON, matchListDeserializationType);
        for (Match m : matches) {
          if (m.primaryMatch == 1) {
            match = m;
            break;
          }
        }
        if (match == null) {
          // No match found
          continue;
        }
        String tenantId = row.getString("tenant_id");
        String itemId = row.getString("item_id");
        String compId = row.getString("comp_id");
        String compItemId = match.compItemId;
        String compURL = match.url;
        String pipeline = match.pipeline;
        double score = match.getScore();
        String miscData = match.getMiscData();
        String createdAt = match.created_at;
        String updatedAt = match.updated_at;
        String status = match.matchStatus;
        ItemMatch itemMatch = new ItemMatch(tenantId, itemId, compId,compItemId, compURL, pipeline, matchType, score, miscData, createdAt, updatedAt, status, dt);
        System.out.println("Match #" + matchCount);
        System.out.println("" + itemMatch);
        matchCount++;
      }
    } catch (Exception e) {
      System.out.println("\n Exception: " + e);
      e.printStackTrace();
    }
    System.out.println("");
    log("Found " + matchCount + " matches out of " + rowCount + " rows.");
    cluster.close();
  }
}

class ItemMatch {
  String tenantId;
  String itemId;
  String compId;
  String domain;
  String compItemId;
  String compURL;
  String pipeline;
  String matchType;
  double score;
  String miscData;
  String createdAt;
  String updatedAt;
  String status;
  String dt;

  public ItemMatch(String tenantId, String itemId, String compId, String compItemId, String compURL, String pipeline,
    String matchType, double score, String miscData, String createdAt, String updatedAt, String status, String dt) {
    this.tenantId = tenantId;
    this.itemId = itemId;
    this.compId = compId;
    this.domain = extractDomain(compURL);
    this.compItemId = compItemId;
    this.compURL = compURL;
    this.pipeline = pipeline;
    this.matchType = matchType;
    this.score = score;
    this.miscData = miscData;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.status = status;
    this.dt = dt;
  }

  private String extractDomain(String url) {
    if (url == null) return null;
    String domain = url;
    int index = url.indexOf("://");
    if (index >= 0) {
      domain = url.substring(index + 3);
    }
    index = domain.indexOf('/');
    if (index >= 0) {
      domain = domain.substring(0, index);
    }
    domain = domain.replaceFirst("^www.*?\\.", "");
    return domain;
  }

  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("ItemMatch: ");
    sb.append("\n  tenantId=" + tenantId);
    sb.append("\n  itemId=" + itemId);
    sb.append("\n  compId=" + compId);
    sb.append("\n  domain=" + domain);
    sb.append("\n  compItemId=" + compItemId);
    sb.append("\n  compURL=" + compURL);
    sb.append("\n  pipeline=" + pipeline);
    sb.append("\n  matchType=" + matchType);
    sb.append("\n  score=" + score);
    sb.append("\n  miscData=" + miscData);
    sb.append("\n  createdAt=" + createdAt);
    sb.append("\n  updatedAt=" + updatedAt);
    sb.append("\n  dt=" + dt);
    sb.append("\n");
    return sb.toString();
  }
}

class Match {
  String compItemId;
  String url;
  List<Map<String, String>> algoScore;
  int primaryMatch;
  String pipeline;
  String matchStatus;
  String created_at;
  String updated_at;
  Map<String, String> miscAttributes;

  public double getScore() {
    double score = 0.0;
    if (algoScore != null) {
      // There should only be one algoScore, but in case there are more only take the first one
      Map<String, String> map = algoScore.get(0);
      String scoreString = map.get("score");
      if (scoreString != null) {
        score = Double.parseDouble(scoreString);
      }
    }
    return score;
  }

  public String getMiscData() {
    if (miscAttributes == null) {
      return null;
    }
    StringBuilder sb = new StringBuilder("{");
    int i = 0;
    for (Map.Entry<String,String> entry : miscAttributes.entrySet()) {
      if (i > 0) {
        sb.append(", ");
      }
      sb.append("\"");
      sb.append(entry.getKey());
      sb.append("\":\"");
      sb.append(entry.getValue());
      sb.append("\"");
      i++;
    }
    sb.append("}");
    return sb.toString();
  }

  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("Match: ");
    sb.append("\n  compItemId=" + compItemId);
    sb.append("\n  url=" + url);
    sb.append("\n  algoScore=" + algoScore);
    sb.append("\n  primaryMatch=" + primaryMatch);
    sb.append("\n  pipeline=" + pipeline);
    sb.append("\n  matchStatus=" + matchStatus);
    sb.append("\n  created_at=" + created_at);
    sb.append("\n  updated_at=" + updated_at);
    sb.append("\n  miscAttributes=" + miscAttributes);
    sb.append("\n");
    return sb.toString();
  }
}

