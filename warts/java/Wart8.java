import java.util.IdentityHashMap ;
import java.util.Map ;

public class Wart8 {
  public static void main(String[] args) {
    Map map = new IdentityHashMap();
    Object o = new Object();
    map.put("a", o);
    map.put("a", o);
    map.put(new String("a"), o);
    map.put(new String("a"), o);
    System.out.println(map.size());
  }
}

