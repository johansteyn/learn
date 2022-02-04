Based on https://www.youtube.com/watch?v=BZYj6yfA6Bs

First create an empty WORKSPACE file to mark this folder as a Bazel project:

  % touch WORKSPACE

Then create the BUILD file (which could also be named BUILD.bazel).

Then build using:

  % bazel build //...

or:

  % bazel build :helloworld

This will create a few folders off /private/var/tmp/_bazel_johan.steyn, and links to those folders.

The resulting output file will be:

  bazel-bin/helloworld.txt 


