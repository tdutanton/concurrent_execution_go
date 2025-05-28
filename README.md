# Concurrent Execution in Go

This repository contains three small Go programs that demonstrate different approaches to **concurrent programming** using goroutines, channels, and synchronization primitives.

Each program is implemented as a separate subproject and showcases idiomatic use of Go for solving typical concurrency-related problems.

## 📦 Projects

### 1. ⏱️ Asynchronous Task Stopwatch

**Description:**  
This program launches `N` goroutines, each of which sleeps for a random duration up to `M` milliseconds. It then waits for all goroutines to complete and prints the sleep durations sorted in **descending** order.

**Usage:**

```bash
task exe_01 n=10 m=500
or
go run main.go -n 10 -m 500
```

**Parameters:**

- `-n` — number of goroutines
- `-m` — maximum sleep duration (ms)

**Features:**

- Uses `sync.WaitGroup` to wait for all goroutines
- No channels used
- Prints: `<goroutine index>: <sleep duration>` sorted by sleep duration (high to low)

---

### 2. 🧮 Square Number Generator

**Description:**  
This program concurrently generates numbers from `K` to `N`, passes them through a channel to another function which computes their squares and prints them.

**Usage:**

```bash
task exe_02 k=1 n=10
or
go run main.go -k 1 -n 10
```

**Parameters:**

- `-k` — start of the number range
- `-n` — end of the number range

**Features:**

- Two goroutines: one for generation, one for squaring
- Uses channels with proper read/write restrictions
- Channels are closed correctly
- Squaring is done **sequentially** as numbers are read

---

### 3. ⏰ Ticker with Signal Handling

**Description:**  
A ticker that prints a message every `K` seconds and stops gracefully when it receives a `SIGINT` or `SIGTERM`.

**Usage:**

```bash
task exe_03 k=2
or
go run main.go -k 2
```

**Parameters:**

- `-k` — tick interval in seconds

**Features:**

- Prints: `Tick <i> since <elapsed time in seconds>`
- Handles OS signals gracefully
- Uses `os/signal.Notify` for interruption
- Implements ticker using `time.Sleep` (no `time.Ticker`)

---

## 🛠️ Technologies

- Go 1.24+
- Standard Go libraries only (`sync`, `time`, `os`, `flag`, `signal`, etc.)

## 📁 Structure

```bash
.
├── cmd/
	├── 01_asyncstopwatch/  # Project 1 / exe_01
	├── 02_gensquares/		# Project 2 / exe_02
	└── 03_ticker/			# Project 3 / exe_03
└── internal/
	├── 01_asyncstopwatch/
	├── 02_gensquares/
	└── 03_ticker/
```

Each directory contains a standalone Go application with its own `main.go`.

---

## 🚀 How to Run

Clone the repository:

```bash
git clone https://github.com/yourusername/concurrent-execution-go.git
cd concurrent-execution-go
```

Run any project:

```bash
make (install Taskfile if needed)
task (read main targets and info)

task build (building exe files)

task exe_01 n=... m=...

task exe_02 k=... n=...

task exe_03 k=...

OR JUST

cd stopwatch
go run main.go -n 5 -m 1000
```

---

## 📄 Author
- [Anton Evgenev](https://t.me/tdutanton)

2025