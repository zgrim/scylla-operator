#!/bin/bash
set -euo pipefail

CLUSTER_NAME=meteringlab-c2
EKS_REGION=us-east-1

function wait-for-object-creation {
    for i in {1..30}; do
        { kubectl -n "${1}" get "${2}" && break; } || sleep 1
    done
}

eksctl create cluster -f ./eks-cluster.yaml 

# Configure node disks and network
kubectl apply -f node-setup-daemonset.yaml
wait-for-object-creation default daemonset.apps/node-setup
kubectl rollout status --timeout=5m daemonset.apps/node-setup
