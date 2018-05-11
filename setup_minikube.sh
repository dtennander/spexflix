#!/usr/bin/env bash

set -e

DIR=$(mktemp -d)

function deleteTempDir {
  rm -rf ${DIR}
}
trap deleteTempDir EXIT

# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out ${DIR}/server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out ${DIR}/server.key

openssl req -new -x509 -sha256 \
    -key ${DIR}/server.key \
    -out ${DIR}/server.crt \
    -days 3650 \
    -subj "/C=SE/ST=David/L=Stockholm/O=Dis/CN=www.example.com"

minikube start

# Set docker to point to the minikube daemon
eval $(minikube docker-env --shell bash)
function unsetDockerVarables {
    deleteTempDir
    eval $(minikube docker-env --shell bash -u)
}
trap unsetDockerVarables EXIT

kubectl create namespace spexflix
kubectl create secret tls tls-secret --cert=${DIR}/server.crt --key=${DIR}/server.key -n spexflix

./redeploy_minikube.sh

minikube dashboard
