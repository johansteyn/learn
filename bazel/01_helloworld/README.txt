Based on https://www.youtube.com/watch?v=BZYj6yfA6Bs

First create an empty WORKSPACE file to mark this folder as a Bazel project:

  % touch WORKSPACE

Then create the BUILD file (which could also be named BUILD.bazel).

Then build using:

  % bazel build //...

This will create a few folders off /private/var/tmp/_bazel_johan.steyn, and links to those folders.

The "helloworld.txt" output file will be in:

  bazel-01_helloworld/bazel-out/darwin-fastbuild/bin/helloworld.txt 


