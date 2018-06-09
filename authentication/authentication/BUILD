load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "authDao.go",
        "jwtAuthenticator.go",
    ],
    importpath = "github.com/DiTo04/spexflix/authentication/authentication",
    visibility = ["//visibility:public"],
    deps = [
        "//common/jwt:go_default_library",
        "@com_github_dgrijalva_jwt_go//:go_default_library",
        "@com_github_googlecloudplatform_cloudsql_proxy//proxy/dialers/postgres:go_default_library",
        "@org_golang_x_crypto//bcrypt:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["jwtAuthenticator_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//authentication/authentication/mock_authentication:go_default_library",
        "@com_github_golang_mock//gomock:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
        "@org_golang_x_crypto//bcrypt:go_default_library",
    ],
)
