Two modules that can be built in parallel.

Both modules will take at least 5 seconds to build.

Build either module individually:

  % bazel build //moduleA
  % bazel build //moduleB

Side note: if you need to specify a specific target within a module:

  % bazel build //moduleA:moduleA
  % bazel build //moduleB:moduleB

Or build both modules in parallel:

  % bazel build //...

When building the modules individually, in sequence, the time taken will be
at least 10 seconds: at least 5 seconds for each module.

But when built in parallel it will take just over 5 seconds, 
ie. about the time it takes to build just one module.

Whether building individually or in parallel, if nothing has changed then subsequent
builds are very fast, ie. the commands (including 5 seconds sleep) are not run.

