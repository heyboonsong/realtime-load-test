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

### Server-Sent Events (SSE) Load Test

An SSE load test is included in the `docker-compose.yml` file. To run it:

```bash
docker-compose up --scale k6-sse=1
```

This will start a k6 instance that runs the SSE load test (`sse.js`) with 10,000 VUs for 30 seconds.

### WebSocket Load Test

To run the WebSocket load test locally, first ensure the application is running (`docker-compose up -d`), then execute:

```bash
k6 run --vus 10 --duration 30s ws-load-test.js
```

This command will run the WebSocket load test (`ws-load-test.js`) with 10 virtual users for 30 seconds against the `ws://localhost:9000/ws` endpoint.

## Monitoring

cAdvisor is included in the `docker-compose.yml` file to provide container resource monitoring. You can access the cAdvisor web UI at `http://localhost:8080`.