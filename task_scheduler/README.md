# Scope Of Work

| Feature                    | Must Have? | Description                             |
| -------------------------- | ---------- | --------------------------------------- |
| Enqueue Jobs               | ✅         | Add new job to queue                    |
| Worker Pool                | ✅         | Process jobs using goroutines           |
| Retry on Failure           | ✅         | Retry job X times if it fails           |
| Delayed Job Execution      | ✅         | Run job after delay or at specific time |
| Scheduled / Recurring Jobs | ✅         | Use cron-like jobs                      |
| Job Status Tracking        | ✅         | Keep track of current state             |
| Graceful Shutdown          | ✅         | Clean exit of app                       |
| REST API to manage jobs    | ✅         | Interface to trigger/view/cancel        |
| Persistent Storage         | Optional   | Store jobs in Redis/Postgres/etc.       |
| Job Metrics & Logging      | Optional   | For monitoring health and debugging     |
| Priority Queue             | Optional   | Handle critical jobs first              |
| Web UI Dashboard           | Optional   | View jobs visually (like Celery Flower) |

# Phase 1: Basic Project Setup – In-Memory Delayed Task Queue

## Features

- `POST /push` – Add a task with payload and delay
- `GET /consume` – Fetch one task that’s ready
- Goroutine-based scheduler
- In-memory task queue and ready queue

Why it's Poll-Based:
You expose a GET /consume endpoint.

A worker (consumer) polls this endpoint to check if a task is ready.

The server doesn’t push the task — instead, the client must ask repeatedly or on interval.

### 📌 Why It's Poll-Based

- You expose a `GET /consume` endpoint.
- A worker (consumer) **polls** this endpoint to check if a task is ready.
- The server doesn’t **push** the task — instead, the client must **ask repeatedly or at intervals**.
