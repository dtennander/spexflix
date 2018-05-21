load("@k8s_production_deploy//:defaults.bzl", "k8s_production_deploy")

k8s_production_deploy(
    name = "prod-ingress",
    template = "ingress.yaml",
)

k8s_production_deploy(
    name = "prod-certificate",
    template = "certificate.yaml",
)