Based on https://www.youtube.com/watch?v=e9dlx2ilwr0

This is for building a Java binary from a single source file.

Note that the location of source file needs be specified (somehow... TODO).
Else it needs to be inferred by being placed in some known location.
If no source file is found it will result in error:

main_class was not provided and cannot be inferred: 
  source path doesn't include a known root (java, javatests, src, testsrc)

In this case, we place the single Java source file in "src".

Run using:

  % bazel run :App

But that outputs Bazel stuff too, so rather build and run separately.

Build using:

  % bazel build :App

This also works:

  % bazel build //:App

The .class file will be in:

  bazel-bin/_javac/App/App_classes

So it can be run using:
  
  % java -classpath bazel-bin/_javac/App/App_classes App

But it will also create a JAR file in:

  bazel-bin

Which can be run using:

  % java -classpath bazel-bin/App.jar App

Note that this doesn't work, as it doesn't have a manifest:

  % java -jar bazel-bin/App.jar

Finally, it also produces a binary which can be run using:

  % bazel-bin/App


Other Bazel commands
--------------------

Display runtime info about the bazel server:

  % bazel info


  

