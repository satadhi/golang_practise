# Task Scheduler with Leader Election

This project implements a distributed task scheduler with **leader election**, ensuring:

- All instances can **receive tasks** via `POST /push`.
- **Only one instance (the leader)** is responsible for task **consumption** via `POST /consume`.

---

## 📐 Architecture Overview

| Endpoint        | Handled By      | Description                               |
| --------------- | --------------- | ----------------------------------------- |
| `POST /push`    | All instances   | Accept and store tasks in a shared queue  |
| `POST /consume` | Only the leader | Periodically consumes and processes tasks |

---

## 🛠️ Components

### ✅ Shared Task Store

All instances interact with a **shared backend** (e.g., Redis, MySQL) to:

- Enqueue tasks
- Dequeue/process tasks (only by the leader)

---

### 🔁 Leader Election Approaches

Option 2: Bully Algorithm (Simple DIY)

- Each instance has a unique ID and pings nodes with higher IDs. If no higher node responds, it becomes the leader.



### 📦 Future Enhancements
- Add metrics/logs for leader health & status
- Use gRPC for internal instance communication
- Add persistence layer for task retries and DLQ
- Build dashboard to monitor leader and tasks

