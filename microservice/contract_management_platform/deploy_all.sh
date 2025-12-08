#!/bin/bash
set -e

NAMESPACE="contract-platform"
K8S_DIR="k8s"

# ตรวจสอบ namespace
if ! kubectl get ns $NAMESPACE &> /dev/null; then
  echo "Creating namespace $NAMESPACE..."
  kubectl create namespace $NAMESPACE
else
  echo "Namespace $NAMESPACE already exists."
fi

# ฟังก์ชัน deploy service
deploy_service() {
  SERVICE_NAME=$1
  echo "Deploying $SERVICE_NAME..."
  kubectl apply -f $K8S_DIR/$SERVICE_NAME/ -n $NAMESPACE
  echo "Waiting rollout for $SERVICE_NAME..."
  kubectl rollout status deployment/$SERVICE_NAME -n $NAMESPACE
  echo "$SERVICE_NAME deployed successfully!"
  echo "---------------------------------"
}

# Deploy ทั้ง 4 services
SERVICES=("auth-service" "contract-service" "contract-proxy" "notification-worker")

for service in "${SERVICES[@]}"; do
  deploy_service $service
done

# แสดง pods และ services หลัง deploy
echo "All services deployed!"
kubectl get pods -n $NAMESPACE
kubectl get svc -n $NAMESPACE
