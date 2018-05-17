load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push", "container_image")

go_image(
    name = "image",
    srcs = ["main.go"],
    goos = "linux",
    pure = "on",
    importpath = "github.com/DiTo04/spexflix/authentication",
    visibility = ["//visibility:public"],
    deps = [
        "//authentication/authentication:go_default_library",
        "//authentication/server:go_default_library",
        "//common/codecs:go_default_library",
    ],
)