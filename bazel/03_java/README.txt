This is for a multiple Java source files across packages.

It builds a Java application and a Java library - all from a single BUILD file.

The application depends on the library.

Build both the application and the library using:

  % bazel build :App

This creates the application and library JAR files:

  bazel-bin/App.jar
  bazel-bin/libutil.jar


TODO: BUILD files in subdirs...

