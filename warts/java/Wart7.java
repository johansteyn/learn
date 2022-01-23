import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;

public class Wart7 {
  public static void main(String[] args) {
    Collection<Object> c = new ArrayList<Object>();
    Collections.addAll(c, "aaa", "bbb", "ccc");
    c.remove(0);
    System.out.println(c);
  }
}

