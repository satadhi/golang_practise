# 📊 Concurrent Word & Letter Counter in Go

This project reads a given text file (`input.txt`) and performs two concurrent operations:

1. **Word Count** – Counts the frequency of each word.
2. **Letter Count** – Counts the frequency of each letter (case-insensitive, alphabetic only).

Both tasks are executed using separate goroutines, demonstrating Go's powerful concurrency features using `channels` and `sync.WaitGroup`.

---

## 🧠 Concepts Covered

- Goroutines
- Channels
- WaitGroups
- File I/O (Read/Write)
- Maps for frequency counting
- String manipulation and cleanup

---

## 📂 Output

- `word_counts.txt` – Contains each word and its count.
- `letter_counts.txt` – Contains each letter and its count.

Example output:
