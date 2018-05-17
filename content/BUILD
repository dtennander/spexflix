load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "image",
    srcs = ["main.go"],
    goos = "linux",
    pure = "on",
    importpath = "github.com/DiTo04/spexflix/content",
    visibility = ["//visibility:public"],
    deps = [
        "//common/codecs:go_default_library",
        "//content/content:go_default_library",
        "//content/server:go_default_library",
    ],
)
