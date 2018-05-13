load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/spexflix/front-end",
    visibility = ["//visibility:private"],
    deps = [
        "//common/codecs:go_default_library",
        "//front-end/content_client:go_default_library",
        "//front-end/server:go_default_library",
    ],
)

go_binary(
    name = "front-end",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
