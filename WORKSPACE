http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz",
    sha256 = "c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz",
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "27c94dec66c3c9fdb478c33994471c5bfc15b6eb",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

git_repository(
    name = "io_bazel_rules_k8s",
    commit = "8c9b9cbc6a46a4c8db4a1da8565252b326d90331",
    remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_gorilla_mux",
    commit = "e3702bed27f0d39777b0b37b664b6280e8ef8fbf",
    importpath = "github.com/gorilla/mux",
)

go_repository(
    name = "com_github_stretchr_testify",
    commit = "c679ae2cc0cb27ec3293fea7e254e47386f05d69",
    importpath = "github.com/stretchr/testify",
)

go_repository(
    name = "com_github_gorilla_context",
    commit = "08b5f424b9271eedf6f9f0ce86cb9396ed337a42",
    importpath = "github.com/gorilla/context",
)

go_repository(
    name = "com_github_urfave_negroni",
    commit = "22c5532ea862c34fdad414e90f8cc00b4f6f4cab",
    importpath = "github.com/urfave/negroni",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load(
    "@io_bazel_rules_docker//docker:docker.bzl",
    "docker_repositories",
)

docker_repositories()

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories", "k8s_defaults")

k8s_repositories()

k8s_defaults(
    # This becomes the name of the @repository and the rule
    # you will import in your BUILD files.
    name = "k8s_dev_deploy",
    namespace = "spexflix-develop",
    cluster = "gke_spexflix_europe-west1-b_develop",
)

k8s_defaults(
    # This becomes the name of the @repository and the rule
    # you will import in your BUILD files.
    name = "k8s_production_deploy",
    namespace = "spexflix-production",
    cluster = "gke_spexflix_europe-west1-b_develop",
)
