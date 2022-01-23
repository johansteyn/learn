import java.util.*;

// http://weblogs.java.net/blog/arnold/archive/2005/06/generics_consid_1.html
// I don't see why this is a wart... compiles fine for me!
public class Wart01 {
    public static void main(String[] args) {
System.out.println("1");
        List a  = Collections.emptyList(); 
System.out.println("2");
        filter(a);   //compiles 
System.out.println("3");
        filter(Collections.emptyList());  //does not compile 
System.out.println("4");
    }

    public static void filter(List stringList){
    }
}

