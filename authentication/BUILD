load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push", "container_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/spexflix/authentication",
    visibility = ["//visibility:private"],
    deps = [
        "//authentication/authentication:go_default_library",
        "//authentication/server:go_default_library",
        "//common/codecs:go_default_library",
        "@com_github_urfave_negroni//:go_default_library",
    ],
)

go_image(
    name = "image",
    embed = [":go_default_library"],
    goos = "linux",
    pure = "on",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "authentication",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
