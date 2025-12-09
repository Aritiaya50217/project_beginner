#!/bin/bash
set -e

# =========================
# Configuration
# =========================
# ชื่อ service ใช้ build image
SERVICES=("auth-service" "contract-service" "contract-proxy" "notification-worker")

# ชื่อ container ใน docker-compose
CONTAINERS=("auth_service" "contract_service" "contract_proxy" "notification_worker")

DOCKER_USER="artis50217"

# Optional: ใช้ commit SHA เป็น tag
GIT_SHA=$(git rev-parse --short HEAD)

# =========================
# 1. Build & Test each service
# =========================
echo "==== Building & Testing Services ===="
for svc in "${SERVICES[@]}"; do
  echo "---- $svc ----"
  
  # Install dependencies
  if [ -f "./$svc/package.json" ]; then
    echo "Installing dependencies..."
    npm install --prefix "./$svc"
  fi

  # Run tests
  if [ -f "./$svc/package.json" ]; then
    echo "Running tests..."
    npm test --prefix "./$svc"
  fi

  # Build Docker image with latest & sha tag
  echo "Building Docker image..."
  docker build -t "$DOCKER_USER/$svc:latest" -t "$DOCKER_USER/$svc:$GIT_SHA" "./$svc"
done


# =========================
# 2. Deploy with docker-compose
# =========================
echo "==== Deploying with docker compose ===="

docker compose down
docker compose up --build -d


# =========================
# 3. Wait for containers to be healthy
# =========================
echo "==== Waiting for containers to be healthy ===="

for cname in "${CONTAINERS[@]}"; do
  echo "Waiting for $cname..."

  # ตรวจว่ามี healthcheck มั้ย
  HAS_HEALTHCHECK=$(docker inspect -f '{{json .State.Health}}' "$cname" 2>/dev/null)

  # ถ้าไม่มี healthcheck → ข้าม
  if [ "$HAS_HEALTHCHECK" == "null" ]; then
    echo "No healthcheck for $cname, skipping..."
    continue
  fi

  # ถ้ามี healthcheck → รอจน healthy
  until [ "$(docker inspect -f '{{.State.Health.Status}}' "$cname")" == "healthy" ]; do
    echo "$cname is not healthy yet..."
    sleep 3
  done

  echo "➡ $cname is healthy!"
done


# =========================
# 4. Optional: Tail logs
# =========================
# echo "==== Showing logs ===="
# docker compose logs -f
