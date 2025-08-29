Once upon a time I asked ChatGPT to provide me a list of the most common test assignments for a golang developer on job interviews.

Here what we I got:

# 1. Algorithm & Data Structures (LeetCode-style tasks)

Even though Go isn’t usually used for algorithmic competitions, U.S. companies still test core problem-solving with typical LeetCode tasks:

## Topics:
### Arrays & Strings
* Two Sum
* Subarray Sum
* Sliding Window Maximum
* Rotate Array

### Hash Maps & Sets
* Group Anagrams
* Longest Substring Without Repeating Characters
* Frequency Counting

### Linked Lists
* Reverse Linked List
* Merge Two Sorted Lists
* Detect Cycle

### Trees & Graphs
* BFS/DFS traversal
* Lowest Common Ancestor
* Topological Sort

### Sorting & Searching
* Implement QuickSort/MergeSort
* Binary Search variations


# 2. Go-Specific Concurrency & Channels

**This is the #1 unique Go interview area in U.S. companies:**

## Topics:
### Worker Pool
Implement N workers consuming jobs from a channel.

### Fan-in / Fan-out
Combine multiple input channels into one, or split one channel into many.

### Rate Limiter
Implement a ticker-based or token bucket rate limiter.

### Pipeline
Chain multiple goroutines that process data in stages.

### Safe Counter
Build a concurrency-safe counter using mutexes or channels.

### Context Cancellation
Implement a function that respects context.Context for cancellation.

# 3. REST API & Microservices

**Common "take-home" or live coding tasks:**

* Implement a CRUD REST API in Go (often with Gin/Fiber/Chi, or stdlib net/http).
* Add middleware (logging, authentication, rate limiting).
* Use JSON marshaling/unmarshaling (including handling optional fields).
* Implement pagination or filtering for a list endpoint.
* Connect to a simple DB (Postgres/MySQL) or in-memory store.

# 4. System Design in Go (Practical)

**For mid/senior roles, you’ll often get design questions, sometimes with a bit of coding:**

* Design a URL shortener (tinyurl-like).

* Build a caching system (in-memory LRU cache, sometimes distributed cache).

* Implement a message queue (basic pub/sub or job queue).

* Design a rate limiter for API requests.

* Design a chat system or real-time notification system using WebSockets.

# 5. File & Data Processing

**Very common for backend-oriented interviews:**

* Parse a large log file (streaming, not loading whole file).
* Build a CSV/JSON parser (read, transform, and output data).
* Implement a concurrent file downloader with retry logic.
* Process a stream of events (like stock ticks, logs) and compute aggregates.

# 6. Testing & Code Quality

**U.S. companies often want to see clean Go code practices:**

* Write unit tests using testing package.
* Show mocking skills (sometimes testify/mock or manual interfaces).
* Benchmark code with go test -bench.
* Explain how you’d structure code with interfaces for testability.

# 7. Practical Domain-Specific Tasks

**Depending on the company (fintech, SaaS, infra), you may also see:**

* Parsing JSON/YAML configs.
* gRPC service implementation (common in infra-focused companies).
* Implementing retries with backoff.
* Rate limiting per user/IP.
* Time-series data processing (rolling averages, sliding windows).
