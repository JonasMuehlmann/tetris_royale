#!/usr/bin/env bash

# This is meant to be executed on the machine holding the kubernetes cluster!

SERVICES=("user_service" "matchmaking_service" "gateway")

minikube kubectl -- delete -f microservice-deployment.yaml
minikube stop
minikube start
eval $(minikube docker-env)

# Build all service's docker images
for service in ${SERVICES[@]}; do
    docker build -t ${service} -f ./cmd/$service/Dockerfile .
done

minikube kubectl -- apply -f microservice-deployment.yaml
# kubectl port-forward service/user-service 8080:8080
# kubectl port-forward service/statistics-service 8081:8080

# Useful commands:
# kubectl get services
# kubectl describe deployment user-service
# kubectl describe pod user-service
# kubectl get pods -l app=user-service
# kubectl logs satistics-service-<long-hash>
