This is for a multiple Java source files across packages.

It builds a Java binary and a Java library.

The binary depends on teh library...

Build the library using:

  % cd src/package2
  % bazel build :util


Build the binary using:

  % bazel build :App

Run in same manner as 02_java...

