# Realtime Load Test

This project contains a Go application with three realtime endpoints and corresponding k6 load tests.

## Endpoints

-   **Polling**: `GET /polling` - A simple HTTP endpoint that returns the current time.
-   **Server-Sent Events (SSE)**: `GET /sse` - An endpoint that streams time updates every second.
-   **WebSocket**: `GET /ws` - An endpoint that establishes a WebSocket connection and exchanges messages every second.

## Prerequisites

-   [Docker](https://www.docker.com/get-started)
-   [Docker Compose](https://docs.docker.com/compose/install/)
-   [Go](https://go.dev/doc/install) (for local development)
-   [k6](https://k6.io/docs/get-started/installation/) (for running load tests locally)

## Running the Application

The easiest way to run the application is using Docker Compose:

```bash
docker-compose up -d
```

This will start the Go application on port 9000, along with a cAdvisor instance for monitoring.

## Running Load Tests

### Running All Load Tests

A convenience script is provided to run all load tests sequentially:

```bash
./run-all-tests.sh
```

This script will:
1.  Start the application using `docker-compose up -d`.
2.  Run the Polling, SSE, and WebSocket load tests.
3.  Stop the application using `docker-compose down`.

### Running Individual Load Tests

To run a specific load test, first ensure the application is running (`docker-compose up -d`), then execute one of the following commands:

**Polling Load Test:**
```bash
k6 run --vus 100 --duration 30s polling.js
```

**Server-Sent Events (SSE) Load Test:**
```bash
k6 run --vus 100 --duration 30s sse.js
```

**WebSocket Load Test:**
```bash
k6 run --vus 100 --duration 30s ws-load-test.js
```

## Monitoring

cAdvisor is included in the `docker-compose.yml` file to provide container resource monitoring. You can access the cAdvisor web UI at `http://localhost:8080`.