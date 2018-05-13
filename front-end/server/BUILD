load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "authClient.go",
        "homepageHandlers.go",
        "loginHandlers.go",
        "server.go",
    ],
    importpath = "github.com/DiTo04/spexflix/front-end/server",
    visibility = ["//visibility:public"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["authClient_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//common/codecs:go_default_library",
        "//common/mocks:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)
