import com.segment.analytics.Analytics;
import com.segment.analytics.messages.TrackMessage;
import java.util.LinkedHashMap;
import java.util.Map;

public class SegmentClient {
  public static void main(String[] args) {
    System.out.println("Segment Client");
    String writeKey = System.getenv("SEGMENT_WRITE_KEY");
    Analytics analytics = Analytics.builder(writeKey).build();
    Map<String, Object> properties = new LinkedHashMap<>();
    properties.put("property-key", "property-value");
    analytics.enqueue(TrackMessage
      .builder("Java Test") // The track event name
      .userId("User ID"));
      .properties(properties)
    analytics.flush();
    analytics.shutdown();
  }
}
