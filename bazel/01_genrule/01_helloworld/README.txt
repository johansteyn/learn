Based on https://www.youtube.com/watch?v=BZYj6yfA6Bs

First create an empty WORKSPACE file to mark this folder as a Bazel project:

  % touch WORKSPACE

Then create the BUILD file (which could also be named BUILD.bazel).

Then build using any of:

  % bazel build //...
  % bazel build ...
  % bazel build //:helloworld
  % bazel build :helloworld

The // refers to the root of the workspace.
The ... refers to all targets in the specified path of the workspace.
The :helloworld refers to a specific target in the workspace.

This will create a few folders off /private/var/tmp/_bazel_johan.steyn, and links to those folders.

The resulting output file will be:

  bazel-bin/helloworld.txt 

Note that we have a .bazelversion file, which specifies the version to be used.
Without that file, it will use the latest version of Bazel.


