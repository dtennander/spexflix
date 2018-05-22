load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["http_server.go"],
    importpath = "github.com/DiTo04/spexflix/authentication/server",
    visibility = ["//visibility:public"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["http_server_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_urfave_negroni//:go_default_library",
    ],
)
