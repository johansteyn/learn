This is for a multiple Java source files across packages.

It builds a Java binary and a Java library.

The binary depends on the library...

Build the library using:

  % cd src/package2
  % bazel build :util

This creates JAR file:

  bazel-bin/src/package2/libutil.jar

Build the binary using:

  % bazel build :App

TODO: Not working yet...


