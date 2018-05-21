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

kubectl create namespace spexflix
kubectl create secret tls tls-secret --cert=${DIR}/server.crt --key=${DIR}/server.key -n spexflix-develop
kubectl create secret generic jwt-secret --from-literal=secret="ADD YOUR SECRET HERE"

./redeploy_minikube.sh

minikube dashboard
