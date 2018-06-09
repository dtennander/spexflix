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

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    commit = "06ea1031745cb8b3dab3f6a236daf2b0aa468b7e",
    importpath = "github.com/dgrijalva/jwt-go",
)

go_repository(
    name = "com_github_auth0_go_jwt_middleware",
    commit = "5493cabe49f7bfa6e2ec444a09d334d90cd4e2bd",
    importpath = "github.com/auth0/go-jwt-middleware",
)

go_repository(
    name = "com_github_googlecloudplatform_cloudsql_proxy",
    commit = "74e2f41327763e6a859d91ceff9133e73a14bb3d",
    importpath = "github.com/GoogleCloudPlatform/cloudsql-proxy",
)

go_repository(
    name = "com_github_lib_pq",
    commit = "90697d60dd844d5ef6ff15135d0203f65d2f53b8",
    importpath = "github.com/lib/pq",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "1e0a3fa8ba9a5c9eb35c271780101fdaf1b205d7",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "com_google_cloud_go",
    commit = "7219c21b03ae13acf812621221f3ca5c0f1e769f",
    importpath = "cloud.google.com/go",
)

go_repository(
    name = "org_golang_google_api",
    commit = "00e3bb8d04691e25ee2fccf98c866bcb7925c3ec",
    importpath = "google.golang.org/api",
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

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "static_nginx_server",
    registry = "registry.hub.docker.com",
    repository = "kyma/docker-nginx",
    digest = "sha256:c7e9c0c5d6b3c9112f644006484926aaadc84d99d960d39894cb2f79c399b026",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "8ac0e0d97ce45cd83d1d7243c060cb8461dda5e9",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_golang_mock",
    commit = "22bbf0ddf08105dfa364d0a2fa619dfa71014af5",
    importpath = "github.com/golang/mock",
)
