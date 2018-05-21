load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "jwtMiddleware.go",
        "server.go",
    ],
    importpath = "github.com/DiTo04/spexflix/content/server",
    visibility = ["//visibility:public"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
        "@com_github_urfave_negroni//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "jwtMiddleware_test.go",
        "server_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//common/codecs:go_default_library",
        "//common/mocks:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)
