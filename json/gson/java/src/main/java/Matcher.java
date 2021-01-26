import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import com.google.common.reflect.TypeToken;
import com.google.gson.Gson;

public class Matcher {
  private static Logger logger = Logger.getLogger(Matcher.class.getName());
  private static String filename;
  private static Type matchListDeserializationType = new TypeToken<List<Match>>(){}.getType();

  private static void log(String message) {
    logger.log(Level.INFO, message);
  }

  public static void main(String[] args) throws IOException {
    filename = args[0];
    log("Parsing " + filename + "...");
    Gson gson = new Gson();
    try (Reader reader = new FileReader(filename)) {
      List<Match> matches = gson.fromJson(reader, matchListDeserializationType);
      for (Match match : matches) {
        System.out.println(match);
      }
    }
  }
}

class Match {
  String compItemId;
  String url;
  List<Map<String, String>> algoScore;
  int verificationRequired;
  int primaryMatch;
  String pipeline;
  String matchStatus;
  String created_at;
  String updated_at;
  Map<String, String> miscAttributes;
  boolean isLikeMatch;

  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("Match: ");
    sb.append("\n  compItemId=" + compItemId);
    sb.append("\n  url=" + url);
    sb.append("\n  algoScore=" + algoScore);
    sb.append("\n  verificationRequired=" + verificationRequired);
    sb.append("\n  primaryMatch=" + primaryMatch);
    sb.append("\n  pipeline=" + pipeline);
    sb.append("\n  matchStatus=" + matchStatus);
    sb.append("\n  created_at=" + created_at);
    sb.append("\n  updated_at=" + updated_at);
    sb.append("\n  miscAttributes=" + miscAttributes);
    sb.append("\n  isLikeMatch=" + isLikeMatch);
    sb.append("\n");
    return sb.toString();
  }
}

