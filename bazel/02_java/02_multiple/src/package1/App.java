package package1;

import package2.Util;

public class App {
  public static void main(String[] args) {
    System.out.println("App depends on " + Util.CONSTANT);
  }
}
