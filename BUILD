load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    command = "fix",
    extra_args = [
        "-build_file_name",
        "BUILD,BUILD.bazel",
    ],
    prefix = "github.com/DiTo04/spexflix",
)
