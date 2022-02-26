Without Bazel
-------------

Run:

  % go run greet.go

Build:

  % go build greet.go

Test:

  % cd mypackage
  % go test

Initialize module:

  % go mod init greet

NOTE: 
  No need to run "go mod tidy" it seems...
  No need to run "go mod init" either for Bazel?    


With Bazel
----------

Run:

  % bazel run :greet

Build:
  
  % bazel build :greet

Test:

  % bazel test :test_mypackage

