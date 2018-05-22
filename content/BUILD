load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/spexflix/content",
    visibility = ["//visibility:private"],
    deps = [
        "//common/codecs:go_default_library",
        "//content/content:go_default_library",
        "//content/server:go_default_library",
    ],
)

go_image(
    name = "image",
    embed = [":go_default_library"],
    goos = "linux",
    pure = "on",
    visibility = ["//visibility:public"],
)
