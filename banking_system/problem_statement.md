# 💡 Proof of Concept: Concurrent Banking System Simulation

## 🧾 Overview

This POC demonstrates a simple concurrent banking system using Go. It simulates multiple bank accounts and concurrent money transfers between them, ensuring data consistency using goroutines and mutexes.

---

## 🎯 Objective

- Simulate multiple accounts and concurrent transactions.
- Use `mutex` to prevent race conditions.
- Use `goroutines` for parallel transfer simulation.
- Validate final balance to ensure no money is lost or created.

---

## ⚙️ Key Concepts

- **Goroutines**: To perform transfers in parallel.
- **Mutex**: To lock account balances during updates.
- **WaitGroup**: To wait for all transfers to complete.
- **Optional Channels**: For logging or monitoring.

---

## 🔄 Flow

1. Initialize N accounts with a fixed balance.
2. Perform M random transfers concurrently.
3. Ensure locking during balance updates.
4. Verify total system balance remains constant.

---

## ✅ Success Criteria

- No race conditions (`go run -race`)
- No deadlocks
- Consistent total balance at end

---

## 🧠 Learnings

- Safe concurrent operations in Go.
- Managing shared state with mutexes.
- Designing a simple simulation with real-world concurrency problems.

---
