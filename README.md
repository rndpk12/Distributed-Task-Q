# Distributed Task Queue

A production-ready distributed task queue built using Go, Redis, PostgreSQL, WebSockets, Docker, and Railway.

## Live Demo

Dashboard:
https://distributed-task-q-production.up.railway.app

## Overview

This project demonstrates the design and deployment of a distributed task processing system capable of handling asynchronous workloads through worker nodes.

Tasks are submitted through an API, persisted in PostgreSQL, queued in Redis, processed by distributed workers, and monitored in real-time through a WebSocket-powered dashboard.

## Features

### Core Queue System
- Asynchronous task processing
- Redis-backed task queue
- PostgreSQL task persistence
- Distributed worker architecture
- Worker heartbeat monitoring
- Retry mechanism
- Dead Letter Queue (DLQ)

### Real-Time Monitoring
- Live dashboard
- WebSocket event streaming
- Worker health tracking
- Queue depth monitoring
- Throughput visualization
- Task history tracking

### Infrastructure
- Dockerized services
- Railway deployment
- Neon PostgreSQL integration
- Redis Cloud integration
- Multi-service architecture

---

## Architecture

```text
                    ┌─────────────────┐
                    │     Dashboard   │
                    │   (WebSocket)   │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │    API Server   │
                    │      (Gin)      │
                    └───────┬─────────┘
                            │
            ┌───────────────┼───────────────┐
            │                               │
            ▼                               ▼

    ┌──────────────┐               ┌──────────────┐
    │ PostgreSQL   │               │    Redis     │
    │    (Neon)    │               │ (RedisCloud) │
    └──────────────┘               └──────┬───────┘
                                          │
                                          ▼

                                ┌─────────────────┐
                                │ Worker Service  │
                                │   (Railway)     │
                                └─────────────────┘
```

## System Workflow

```text
Client
   │
   ▼
POST /tasks
   │
   ▼
PostgreSQL
   │
   ▼
Redis Queue
   │
   ▼
Worker
   │
   ▼
Task Processing
   │
   ▼
Status Update
   │
   ▼
Dashboard
```

## Tech Stack

### Backend

- Go
- Gin
- GORM

### Database

- PostgreSQL (Neon)

### Queue

- Redis Cloud

### Realtime

- WebSockets

### Infrastructure

- Docker
- Docker Compose
- Railway

---

## API Endpoints

### Tasks

```http
POST /tasks
GET /tasks
```

### Monitoring

```http
GET /metrics
GET /workers
GET /dlq
```

### DLQ

```http
POST /dlq/retry/:id
```

### Realtime

```http
GET /ws
```

---

## Example Task

Request:

```json
{
  "type": "email.send",
  "payload": {
    "user_id": 42,
    "email": "user@example.com"
  }
}
```

Response:

```json
{
  "id": "task-id",
  "status": "pending"
}
```

---

## Local Development

### Clone Repository

```bash
git clone https://github.com/rndpk12/Distributed-Task-Q.git
cd Distributed-Task-Q
```

### Environment Variables

Create a `.env` file:

```env
DATABASE_URL=your_postgres_connection_string
REDIS_ADDR=your_redis_host
REDIS_USERNAME=default
REDIS_PASSWORD=your_redis_password
```

### Start Infrastructure

```bash
docker compose up
```

### Run API

```bash
go run cmd/api/main.go
```

### Run Worker

```bash
go run cmd/worker/main.go
```

---

## Deployment

### API Service

- Railway
- Dockerfile

### Worker Service

- Railway
- Dockerfile.worker

### Database

- Neon PostgreSQL

### Queue

- Redis Cloud

---

## Screenshots

### Dashboard Overview

_Add screenshot here_

### Worker Monitoring

_Add screenshot here_

### Task History

_Add screenshot here_

### Event Stream

_Add screenshot here_

---

## Current Capabilities

- Task creation
- Task processing
- Worker heartbeat tracking
- Dead Letter Queue
- Retry support
- Real-time updates
- Cloud deployment
- Distributed architecture

---

## Future Improvements

- Exponential backoff retries
- Priority queues
- Scheduled jobs
- Prometheus metrics
- Grafana dashboards
- GitHub Actions CI/CD
- Multiple worker replicas
- Rate limiting
- Email service integration

---

## Resume Highlights

- Built a distributed task queue system using Go, Redis, and PostgreSQL.
- Designed asynchronous job processing with worker heartbeats and failure handling.
- Implemented real-time monitoring using WebSockets and a custom dashboard.
- Containerized services using Docker and deployed them on Railway.
- Integrated Neon PostgreSQL and Redis Cloud for cloud-native infrastructure.

---

## Author

**R N Dhanapraveen Krishna**

GitHub:
https://github.com/rndpk12

LinkedIn:
https://www.linkedin.com/in/r-n-dhanapraveenkrishna
