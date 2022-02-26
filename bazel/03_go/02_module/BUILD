load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
  name = "greet",
  srcs = ["greet.go"],
  deps = [":mypackage"]
)

go_library(
  name = "mypackage",
  importpath = "greet/mypackage",
  srcs = ["mypackage/person.go"],
)

go_test(
  name = "test_mypackage",
  srcs = ["mypackage/person_test.go"],
  embed = [":mypackage"]
)

