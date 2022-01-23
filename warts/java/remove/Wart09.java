import java.util.*;

// Derived from Joshua Bloch's "Still More Java Puzzlers"
// http://www.denverjug.org/meetings/files/200408_Puzzles.pdf
// This won't compile
// Unicode escapes are processed before comments.
// The Unicode escape is replaced with an actual line feed!
// So, avoid using Unicode escapes in your code, 
// or at least use \\u000A instead (if you must)
public class Wart09 {
    public static void main(String[] args) {
        // Note: \\u000A is Unicode representation for newline
        char c = 0x000A;
        System.out.println(c);
    }
}

