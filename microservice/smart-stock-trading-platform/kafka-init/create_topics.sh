#!/bin/bash
set -e

echo "Waiting for Kafka broker to be ready..."
cub kafka-ready -b broker1:9092 1 20

echo "Creating topic orders..."
kafka-topics --create \
  --topic orders \
  --bootstrap-server broker1:9092 \
  --partitions 1 \
  --replication-factor 1 || echo "Topic already exists"

echo "Topic orders is ready"
