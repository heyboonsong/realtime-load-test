#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

echo "Starting the application with Docker Compose..."
docker-compose up -d

echo "Waiting for the application to be ready..."
sleep 5

echo "Running the Polling load test..."
k6 run --vus 100 --duration 30s polling.js

echo "Running the SSE load test..."
k6 run --vus 100 --duration 30s sse.js

echo "Running the WebSocket load test..."
k6 run --vus 100 --duration 30s ws-load-test.js

echo "Stopping the application..."
docker-compose down

echo "All load tests completed."