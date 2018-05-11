#!/usr/bin/env bash

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