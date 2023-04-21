#!/bin/bash
set -euo pipefail

CLUSTER_NAME=meteringlab-c1
EKS_REGION=us-east-1

function wait-for-object-creation {
    for i in {1..30}; do
        { kubectl -n "${1}" get "${2}" && break; } || sleep 1
    done
}

eksctl create cluster -f ./eks-cluster.yaml 

# Install local volume provisioner
echo "Installing local volume provisioner..."
helm install local-provisioner ../../common/provisioner
echo "Your disks are ready to use."

echo "Starting the cert manager..."
kubectl apply -f ../../common/cert-manager.yaml
kubectl wait --for condition=established --timeout=60s crd/certificates.cert-manager.io crd/issuers.cert-manager.io
wait-for-object-creation cert-manager deployment.apps/cert-manager-webhook
kubectl -n cert-manager rollout status --timeout=5m deployment.apps/cert-manager-webhook

echo "Starting the scylla operator..."
kubectl apply -f ../../common/operator.yaml
kubectl wait --for condition=established crd/scyllaclusters.scylla.scylladb.com
wait-for-object-creation scylla-operator deployment.apps/scylla-operator
kubectl -n scylla-operator rollout status --timeout=5m deployment.apps/scylla-operator

echo "Starting the scylla cluster..."
kubectl apply -f cluster.yaml
