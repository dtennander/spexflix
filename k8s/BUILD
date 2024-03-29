load("@k8s_dev_deploy//:defaults.bzl", "k8s_dev_deploy")
load("@k8s_production_deploy//:defaults.bzl", "k8s_production_deploy")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_objects(
    name = "dev",
    objects = [
        ":dev-auth",
        ":dev-content",
        ":dev-web",
        ":dev-api",
        ":dev-users",
    ],
)

k8s_objects(
    name = "production",
    objects = [
        ":production-auth",
        ":production-content",
        ":production-web",
        ":production-api",
        ":production-users",
    ],
)

k8s_objects(
    name = "canary",
    objects = [
        ":canary-auth",
        ":canary-content",
        ":canary-web",
        ":canary-api",
        ":canary-users",
    ],
)

k8s_dev_deploy(
    name = "dev-auth",
    images = {
        "gcr.io/spexflix/authentication:dev": "//authentication:image",
    },
    template = ":develop-authentication.yaml",
)

k8s_dev_deploy(
    name = "dev-users",
    images = {
        "gcr.io/spexflix/users:dev": "//users:image",
    },
    template = ":develop-users.yaml",
)

k8s_dev_deploy(
    name = "dev-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = ":develop-content.yaml",
)

k8s_dev_deploy(
    name = "dev-web",
    images = {
        "gcr.io/spexflix/web:dev": "//web:image",
    },
    template = ":develop-web.yaml",
)

k8s_dev_deploy(
    name = "dev-api",
    images = {
        "gcr.io/spexflix/api:dev": "//api:image",
    },
    template = ":develop-api.yaml",
)

k8s_production_deploy(
    name = "production-auth",
    images = {
        "gcr.io/spexflix/authentication:dev": "//authentication:image",
    },
    template = "authentication.yaml",
)

k8s_production_deploy(
    name = "production-users",
    images = {
        "gcr.io/spexflix/users:dev": "//users:image",
    },
    template = ":users.yaml",
)

k8s_production_deploy(
    name = "production-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = "content.yaml",
)

k8s_production_deploy(
    name = "production-web",
    images = {
        "gcr.io/spexflix/web:dev": "//web:image",
    },
    template = ":web.yaml",
)

k8s_production_deploy(
    name = "production-api",
    images = {
        "gcr.io/spexflix/api:dev": "//api:image",
    },
    template = ":api.yaml",
)

k8s_production_deploy(
    name = "canary-auth",
    images = {
        "gcr.io/spexflix/authentication:dev": "//authentication:image",
    },
    template = ":canary-authentication.yaml",
)

k8s_production_deploy(
    name = "canary-users",
    images = {
        "gcr.io/spexflix/users:dev": "//users:image",
    },
    template = ":canary-users.yaml",
)

k8s_production_deploy(
    name = "canary-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = ":canary-content.yaml",
)

k8s_production_deploy(
    name = "canary-web",
    images = {
        "gcr.io/spexflix/web:dev": "//web:image",
    },
    template = ":canary-web.yaml",
)

k8s_production_deploy(
    name = "canary-api",
    images = {
        "gcr.io/spexflix/api:dev": "//api:image",
    },
    template = ":canary-api.yaml",
)

[genrule(
    name = "canarify-" + yamlFile,
    srcs = [yamlFile],
    outs = ["canary-" + yamlFile],
    cmd = "sed 's/replicas: 2/replicas: 1/g' $< | \
                sed 's/env: production/env: canary/g' | \
                perl -pe 's/(name:\W\w+)-production/\\1-canary/' > $@",
) for yamlFile in glob(["*.yaml"])]

[genrule(
    name = "developify-" + yamlFile,
    srcs = [yamlFile],
    outs = ["develop-" + yamlFile],
    cmd = "sed 's/replicas: 2/replicas: 1/g' $< | \
                sed 's/env: production/env: develop/g' | \
                perl -pe 's/(name:\W\w+)-production/\\1-develop/' > $@",
) for yamlFile in glob(["*.yaml"])]
