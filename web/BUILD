load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
genrule(
    name = "yarn_install",
    srcs = [":package.json"],
    outs = [":node_modules"],
    cmd = "yarn install --prod --modules-folder $@ --non-interactive --frozen-lockfile --cwd $(location package.json)"
)

pkg_tar(
    name = "locked_deps",
    strip_prefix = ".",
    srcs = [
        ":yarn_install",
        "package.json",
        "yarn.lock",
    ],
)

pkg_tar(
    name = "workspace",
    strip_prefix = ".",
    srcs = glob(["src/**/*", "public/**/*"]),
    deps = [
        ":locked_deps"
    ],
)

genrule(
    name = "yarn_build",
    srcs = [":workspace"],
    outs = ["build.tar"],
    cmd = "tar -xf $(location :workspace) && \
           node_modules/react-scripts/bin/react-scripts.js build && \
           tar -cf $@ build"
)