load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "CloudStorageService.go",
        "CloudStorageUrlSigner.go",
        "controller.go",
        "controller_mocks.go",
    ],
    importpath = "github.com/DiTo04/spexflix/content/server",
    visibility = ["//visibility:public"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_auth0_go_jwt_middleware//:go_default_library",
        "@com_github_dgrijalva_jwt_go//:go_default_library",
        "@com_github_golang_mock//gomock:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
        "@com_github_urfave_negroni//:go_default_library",
        "@com_google_cloud_go//storage:go_default_library",
        "@org_golang_google_api//iterator:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["controller_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//common/codecs:go_default_library",
        "@com_github_golang_mock//gomock:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)
