This is for a multiple Java source files across packages.

It builds a Java binary and a Java library.

The binary depends on the library...

Build the library using:

  % bazel build :util

This will create the JAR file: 
  bazel-bin/libutil.jar

Build the binary using:

  % bazel build :App

This will build both the library JAR (if not already built) and the application JAR:
  bazel-bin/libutil.jar
  bazel-bin/App.jar



