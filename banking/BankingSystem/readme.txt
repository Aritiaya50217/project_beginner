Banking System with Hexagonal Architecture

- Gin: HTTP framework

- MongoDB: Database

- Docker: Containerize app + MongoDB

- CI/CD: GitHub Actions pipeline build-test-deploy

- Kubernetes: Deploy app + MongoDB

- Hexagonal Architecture: แยกชั้นชัดเจน (Domain, Application, Infrastructure, API)

วิธีรันด้วย Docker Compose

- docker-compose up --build

วิธีรันบน Kubernetes (Minikube / Cloud)

Step 1: สร้าง Docker Image แล้ว Push

- docker build -t yourdockerhubusername/banking-app:latest .

- docker push yourdockerhubusername/banking-app:latest


Step 2: Deploy MongoDB + App

- kubectl apply -f k8s/mongo-deployment.yaml

- kubectl apply -f k8s/mongo-service.yaml

- kubectl apply -f k8s/backend-deployment.yaml

- kubectl apply -f k8s/backend-service.yaml


Step 3: ตรวจสอบว่า Service ทำงาน

- kubectl get pods
- kubectl get svc

วิธี Stop

- docker-compose down # สำหรับ Docker Compose

- kubectl delete -f k8s/ <filename> 