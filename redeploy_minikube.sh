#!/usr/bin/env bash

# Set docker to point to the minikube daemon
eval $(minikube docker-env --shell bash)
function unsetDockerVarables {
    eval $(minikube docker-env --shell bash -u)
}
trap unsetDockerVarables EXIT

# Make all images
cd authentication
make
cd ../content
make
cd ../front-end/
make
cd ..

# Start up the cluster
kubectl delete -f k8s/ -n spexflix
kubectl create -f k8s/ -n spexflix