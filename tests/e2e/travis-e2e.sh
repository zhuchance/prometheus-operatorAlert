#!/usr/bin/env bash
# exit immediately when a command fails
set -e
# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail
# error on unset variables
set -u
# print each command before executing it
set -x

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x kubectl
curl -Lo kind https://github.com/kubernetes-sigs/kind/releases/download/v0.7.0/kind-linux-amd64
chmod +x kind

./kind create cluster --image=kindest/node:v1.18.0
# the default kube config location used by kind
export KUBECONFIG="${HOME}/.kube/config"

# create namespace, permissions, and CRDs
./kubectl create -f manifests/setup

# wait for CRD creation to complete
until ./kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done

# create monitoring components
./kubectl create -f manifests/

make test-e2e
