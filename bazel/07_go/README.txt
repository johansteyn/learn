Without Bazel
-------------

Run:

  % go run helloworld.go

Build:

  % go build helloworld.go


With Bazel
----------

https://github.com/bazelbuild/rules_go

Run:

  % bazel run :helloworld

Build:

  % bazel run :helloworld

Once built, it can be run using:

  % bazel-bin/helloworld_/helloworld

