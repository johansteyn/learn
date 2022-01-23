Wart0
-----
Outputs "ava".

This is because the constructor is passed a char,
which is an integer type - not a string.

Derived from Joshua Bloch's "Still More Java Puzzlers"
  http://www.denverjug.org/meetings/files/200408_Puzzles.pdf


Wart1
-----
Outputs 0, instead of 12.

This is because we don't have a programmer-defined constructor.
Look carefully... Wart1 has 'void' in front, so it's simply a method!
Methods are allowed to have the same name as the class.

Derived from Joshua Bloch's "Still More Java Puzzlers"
  http://www.denverjug.org/meetings/files/200408_Puzzles.pdf


Wart2
-----
Outputs "Equal!" and "Not Equal!".

The Java specification indicates that certain primitives, 
namely all boolean and byte values, all short and int values
between -127 and 127, and char values from \u0000 \u007F
are to be boxed into the same immutable wrapper objects,
which are then cached and reused.

So, the first test should state that the values are equal,
whereas the second test may state that they are not equal.

However, you cannot depend on the second case as some JVM's may
choose to optimize the code and create one instance for both,
thereby indicating a true result.


Wart3
-----
Outputs 17777, instead of 66666.

This is a low trick...
Look carefully: depending on the font you use the second literal
has the letter 'l', not the digit 1. This makes it a long with value 2.

So, use uppercase L instead of lowercase l for long literals,
and never use l as a variable name.

Derived from Joshua Bloch's "Still More Java Puzzlers"
  http://www.denverjug.org/meetings/files/200408_Puzzles.pdf


Wart4
-----
Outputs 5.

This is because of overflow.
24 * 60 * 60 * 1000 * 1000 is larger than Integer.MAX_VALUE (2147483647)

But, both variables are of type long...

Yes, but the individual literals are ints.
The results of multiplying those ints are also ints,
that just happen to be assigned to longs.

To avoid this, use long for at least one literal
Eg: 24L * 60 * 60 * 1000 * 1000

http://mcqueeney.com/roller/page/tom?entry=new_features_in_jdk_5
http://www.denverjug.org/meetings/files/200408_Puzzles.pdf


Wart5
-----
Outputs 0, instead of 7.

This is because: (float) 1234567890 == (float) 1234567897

So, don't use floats as loop indices!

Derived from Joshua Bloch's "Still More Java Puzzlers"
  http://www.denverjug.org/meetings/files/200408_Puzzles.pdf


Wart6
-----
Outputs 7 the 12 but then throws a null pointer exception.

Everything seems fine at compile time.

And at run time you can add null to the list, and you can also retrieve it.
But you cannot assign the null value to an int.

http://www.jroller.com/page/vprise?entry=java_5_language_changes_were


Wart7
-----
Outputs [bbb, 1, 2]

ArrayList has 2 remove methods: 
  - one that takes an index (of type int),
  - one that takes an object. 

The first will remove the element at the specified index,
whereas the second will remove the specified element.
So, calling remove(0) on an ArrayList will remove the first element.

Collection has only has the second remove method.

Since the instance of ArrayList was declared as type Collection,
it will first look for the method in the Collection class. 
Before Java 5, it would not find it there, so it would look in ArrayList.
But with autoboxing in Java 5, it will convert the 0 to an instance
of Integer, and therefor determine that the remove method in the
Collection class matches, invoke it and then do nothing since there
is no element of type Integer with a value of zero in the collection.


Wart8
-----
Outputs 3, instead of 4.

This is because literals are interned

Derived from Joshua Bloch's "Still More Java Puzzlers"
  http://www.denverjug.org/meetings/files/200408_Puzzles.pdf

