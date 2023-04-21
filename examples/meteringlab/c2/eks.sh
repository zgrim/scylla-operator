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
