load("@k8s_dev_deploy//:defaults.bzl", "k8s_dev_deploy")
load("@k8s_production_deploy//:defaults.bzl", "k8s_production_deploy")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_objects(
    name = "dev",
    objects = [
        ":dev-auth",
        ":dev-content",
        ":dev-front-end",
    ],
)

k8s_objects(
    name = "production",
    objects = [
        ":production-auth",
        ":production-content",
        ":production-front-end",
    ],
)

k8s_objects(
    name = "canary",
    objects = [
        ":canary-auth",
        ":canary-content",
        ":canary-front-end",
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
    name = "dev-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = ":develop-content.yaml",
)

k8s_dev_deploy(
    name = "dev-front-end",
    images = {
        "gcr.io/spexflix/front-end:dev": "//front-end:image",
    },
    template = ":develop-front-end.yaml",
)

k8s_production_deploy(
    name = "production-auth",
    images = {
        "gcr.io/spexflix/authentication:dev": "//authentication:image",
    },
    template = "authentication.yaml",
)

k8s_production_deploy(
    name = "production-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = "content.yaml",
)

k8s_production_deploy(
    name = "production-front-end",
    images = {
        "gcr.io/spexflix/front-end:dev": "//front-end:image",
    },
    template = "front-end.yaml",
)

k8s_production_deploy(
    name = "canary-auth",
    images = {
        "gcr.io/spexflix/authentication:dev": "//authentication:image",
    },
    template = ":canary-authentication.yaml",
)

k8s_production_deploy(
    name = "canary-content",
    images = {
        "gcr.io/spexflix/content:dev": "//content:image",
    },
    template = ":canary-content.yaml",
)

k8s_production_deploy(
    name = "canary-front-end",
    images = {
        "gcr.io/spexflix/front-end:dev": "//front-end:image",
    },
    template = ":canary-front-end.yaml",
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

k8s_dev_deploy(
    name = "dev-ingress",
    template = "ingress.yaml",
)

k8s_production_deploy(
    name = "production-ingress",
    template = "ingress.yaml",
)
