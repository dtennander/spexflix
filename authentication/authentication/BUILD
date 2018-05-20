load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "jwtAuthenticator.go",
    ],
    importpath = "github.com/DiTo04/spexflix/authentication/authentication",
    visibility = ["//visibility:public"],
    deps = [
        "//common/jwt:go_default_library",
        "@com_github_dgrijalva_jwt_go//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
    ],
    embed = [":go_default_library"],
)
