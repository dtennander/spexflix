load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "userService.go",
    ],
    importpath = "github.com/DiTo04/spexflix/users",
    visibility = ["//visibility:private"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_auth0_go_jwt_middleware//:go_default_library",
        "@com_github_dgrijalva_jwt_go//:go_default_library",
        "@com_github_googlecloudplatform_cloudsql_proxy//proxy/dialers/postgres:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
        "@com_github_urfave_negroni//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["main_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
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
    name = "users",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
