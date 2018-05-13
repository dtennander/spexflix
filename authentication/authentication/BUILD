load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "Authenticator.go",
        "InMemorySessionPool.go",
        "session.go",
    ],
    importpath = "github.com/DiTo04/spexflix/authentication/authentication",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "InMemorySessionPool_test.go",
        "session_test.go",
    ],
    embed = [":go_default_library"],
)
