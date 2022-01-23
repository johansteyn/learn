public class Wart2 {
  public static void main(String[] args) {
    test(100, 100);
    test(200, 200);
  }

  static void test(Integer i1, Integer i2) {
    if (i1 == i2) {
      System.out.println("Equal!");
    } else {
      System.out.println("Not equal!");
    }
  }
}

