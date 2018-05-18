load("@k8s_dev_deploy//:defaults.bzl", "k8s_dev_deploy")
load("@k8s_production_deploy//:defaults.bzl", "k8s_production_deploy")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_objects(
   name = "dev",
   objects = [
      ":dev-auth",
      ":dev-content",
      ":dev-front-end",
   ]
)

k8s_objects(
   name = "production",
   objects = [
      ":production-auth",
      ":production-content",
      ":production-front-end",
   ]
)

k8s_objects(
   name = "canary",
   objects = [
      ":canary-auth",
      ":canary-content",
      ":canary-front-end",
   ]
)

k8s_dev_deploy(
  name = "dev-auth",
  template = "authentication.yaml",
  images = {
    "gcr.io/spexflix/authentication:dev" : "//authentication:image"
  },
)

k8s_dev_deploy(
  name = "dev-content",
  template = "content.yaml",
  images = {
    "gcr.io/spexflix/content:dev": "//content:image"
  },
)

k8s_dev_deploy(
  name = "dev-front-end",
  template = "front-end.yaml",
  images = {
    "gcr.io/spexflix/front-end:dev": "//front-end:image"
  },
)

k8s_production_deploy(
  name = "production-auth",
  template = "authentication.yaml",
  images = {
    "gcr.io/spexflix/authentication:dev" : "//authentication:image"
  },
)

k8s_production_deploy(
  name = "production-content",
  template = "content.yaml",
  images = {
    "gcr.io/spexflix/content:dev": "//content:image"
  },
)

k8s_production_deploy(
  name = "production-front-end",
  template = "front-end.yaml",
  images = {
    "gcr.io/spexflix/front-end:dev": "//front-end:image"
  },
)

k8s_production_deploy(
  name = "canary-auth",
  template = ":canary-authentication.yaml",
  images = {
    "gcr.io/spexflix/authentication:dev" : "//authentication:image"
  },
)

k8s_production_deploy(
  name = "canary-content",
  template = ":canary-content.yaml",
  images = {
    "gcr.io/spexflix/content:dev": "//content:image"
  },
)

k8s_production_deploy(
  name = "canary-front-end",
  template = ":canary-front-end.yaml",
  images = {
    "gcr.io/spexflix/front-end:dev": "//front-end:image"
  },
)

[genrule(
         name = "canarify-" + yamlFile,
         srcs = [yamlFile],
         outs = ["canary-" + yamlFile],
         cmd = "sed 's/replicas: 2/replicas: 1/g' $< | \
                sed 's/env: production/env: canary/g' | \
                perl -pe 's/(name:\W\w+)-production/\\1-canary/' > $@"
)
for yamlFile in glob(["*.yaml"])]

k8s_dev_deploy(
  name = "dev-ingress",
  template = "ingress.yaml",
)

k8s_production_deploy(
  name = "production-ingress",
  template = "ingress.yaml",
)

