import java.util.ArrayList;
import java.util.List;

public class Wart6 {
  public static void main(String[] args) {
    List<Integer> list = new ArrayList<Integer>();

    list.add(7);
    list.add(new Integer(12));
    list.add(null);

    int x = list.get(0);
    System.out.println(x);
    int y = list.get(1);
    System.out.println(y);
    int z = list.get(2);
    System.out.println(z);
  }
}

