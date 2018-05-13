load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/spexflix/authentication",
    visibility = ["//visibility:private"],
    deps = [
        "//authentication/authentication:go_default_library",
        "//authentication/server:go_default_library",
        "//common/codecs:go_default_library",
    ],
)

go_binary(
    name = "authentication",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
