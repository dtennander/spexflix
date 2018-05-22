load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "html_pages",
    srcs = glob(["html-templates/*"]),
    mode = "0644",
    package_dir = "/html-templates",
)

pkg_tar(
    name = "public_files",
    srcs = glob(["public/*"]),
    mode = "0644",
    package_dir = "/public",
)

container_image(
    name = "template_image",
    base = "@go_image_base//image",
    tars = [
        ":html_pages",
        ":public_files",
    ],
)

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

go_image(
    name = "image",
    embed = [":go_default_library"],
    goos = "linux",
    pure = "on",
    base = ":template_image",
    visibility = ["//visibility:public"],
)
