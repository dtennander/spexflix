load("@k8s_dev_deploy//:defaults.bzl", "k8s_dev_deploy")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_objects(
   name = "dev",
   objects = [
      ":dev-auth",
      ":dev-content",
      ":dev-front-end",
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