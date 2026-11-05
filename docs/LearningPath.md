# Go Projects: Beginner to Advanced

The projects are grouped into tiers. Each lists what you build, the fundamentals it exercises, and a stretch goal. The tiers build on each other, so it's worth doing them roughly in order.

## Tier 1: Language Basics

*Goal: get comfortable with syntax, types, and the standard toolchain.*

**1. Unit Converter CLI**

- **Build:** A command-line tool that converts temperature, length, and weight (`convert 100 km mi`).
- **Practices:** `os.Args`, the `flag` package, `switch`, functions with multiple return values, `strconv`, basic error handling.
- **Stretch:** Add a lookup table with `map[string]float64` and support chained units.

**2. Word Frequency Counter**

- **Build:** Read a text file and print the top N most-used words.
- **Practices:** `bufio.Scanner`, `strings`, maps, slices, `sort.Slice`, file I/O with `defer file.Close()`.
- **Stretch:** Ignore stop words, handle Unicode correctly (runes vs. bytes), and read from stdin as well as files.

**3. Todo List with JSON Persistence**

- **Build:** `todo add "buy milk"`, `todo list`, `todo done 3`, stored in a JSON file.
- **Practices:** Structs, methods, slices of structs, `encoding/json` and struct tags, pointers vs. values, the `os` package.
- **Stretch:** Add due dates (`time.Time`), priorities, and a subcommand structure.

**4. Bank Account Simulator**

- **Build:** Accounts with deposit, withdraw, and transfer operations.
- **Practices:** Methods on pointer receivers, custom error types, `errors.Is` / `errors.As`, error wrapping with `%w`, `fmt.Stringer`.
- **Stretch:** Add transaction history and overdraft rules, with a table-driven test for each rule.

## Tier 2: Idiomatic Go

*Goal: learn interfaces, packages, testing, and the standard library.*

**5. Shapes & Geometry Library**

- **Build:** A `Shape` interface (`Area()`, `Perimeter()`) with circle, rectangle, and triangle implementations.
- **Practices:** Interfaces and implicit satisfaction, type switches, embedding, organizing code into your own package, `go mod init`.
- **Stretch:** Add a `Scaler` interface and show interface composition.

**6. Duplicate File Finder**

- **Build:** Walk a directory, hash files, and report duplicates.
- **Practices:** `filepath.WalkDir`, `crypto/sha256`, `io.Copy`, `io.Reader`/`io.Writer`, maps of slices.
- **Stretch:** Pre-filter by file size to avoid hashing everything, and add a `--delete` flag with confirmation.

**7. Log Analyzer**

- **Build:** Parse an Apache/nginx-style log and report requests per status code, top endpoints, and slowest requests.
- **Practices:** `regexp`, `time.Parse`, custom types, sorting, `text/tabwriter` for output.
- **Stretch:** Time-window filtering (`--since 1h`) and CSV export.

**8. In-Memory Key-Value Store with Tests**

- **Build:** A package with `Get`, `Set`, `Delete`, and TTL expiry.
- **Practices:** Table-driven tests, subtests (`t.Run`), `sync.RWMutex`, benchmarks (`testing.B`), a clean public API, `go vet`.
- **Stretch:** Run `go test -race` and a background goroutine that evicts expired keys.

**9. Expression Evaluator**

- **Build:** Parse and evaluate `3 + 4 * (2 - 1)`.
- **Practices:** A tokenizer, recursive-descent parsing, recursion, structs and interfaces for an AST, error reporting with positions.
- **Stretch:** Add variables, functions (`sqrt(16)`), and a REPL.

## Tier 3: Concurrency

*Goal: internalize goroutines, channels, `select`, `sync`, and `context`.*

**10. Concurrent Link Checker / Web Crawler**

- **Build:** Given a URL, crawl same-domain pages and report broken links.
- **Practices:** Goroutines, `sync.WaitGroup`, channels, a mutex-protected visited set, `net/http` clients, `context` with timeouts.
- **Stretch:** Limit concurrency to N, respect `robots.txt`, and add depth limits.

**11. Worker Pool File Processor**

- **Build:** A pool of workers that process jobs, such as resizing images or compressing files.
- **Practices:** The worker pool pattern, buffered vs. unbuffered channels, fan-out/fan-in, graceful shutdown, closing channels correctly.
- **Stretch:** Make it a multi-stage pipeline (read → transform → write) where each stage is its own goroutine group.

**12. Rate Limiter**

- **Build:** Implement a token bucket and a sliding window limiter.
- **Practices:** `time.Ticker`, `select` with timeouts, `sync/atomic`, designing a concurrency-safe API.
- **Stretch:** Per-key limiting (per IP), with cleanup of stale entries.

**13. In-Process Pub/Sub Broker**

- **Build:** `Subscribe(topic)` returns a channel and `Publish(topic, msg)` fans out to subscribers.
- **Practices:** Channels of channels, avoiding goroutine leaks, `context` cancellation, handling slow consumers.
- **Stretch:** Support wildcard topics and configurable back-pressure policies (drop vs. block).

**14. Concurrent Port Scanner**

- **Build:** Scan a host's ports with configurable parallelism.
- **Practices:** `net.DialTimeout`, bounded concurrency with a semaphore channel, result aggregation, deadlines.
- **Stretch:** Streaming output as results arrive, sorted at the end.

## Tier 4: Networking, Web, and Generics

*Goal: build real services and use modern Go features.*

**15. REST API for a Notes/Bookmarks Service**

- **Build:** CRUD endpoints using only `net/http` (use the Go 1.22+ pattern-based `ServeMux`).
- **Practices:** Handlers, JSON encoding/decoding, middleware (logging, auth, recovery), `httptest`, request validation, structured logging with `log/slog`.
- **Stretch:** Pagination, ETags, and graceful shutdown with `http.Server.Shutdown`.

**16. URL Shortener with a Database**

- **Build:** Shorten URLs, redirect, and track click counts, backed by SQLite or Postgres.
- **Practices:** `database/sql`, prepared statements, transactions, migrations, dependency injection through interfaces, config via environment variables.
- **Stretch:** Add a cache layer in front of the DB and write integration tests.

**17. TCP Chat Server**

- **Build:** Multiple clients connect via `nc` or telnet, with rooms, nicknames, and broadcast.
- **Practices:** `net.Listener`, one goroutine per connection, a central hub goroutine, `bufio.Reader`, handling disconnects cleanly.
- **Stretch:** Switch to WebSockets with a small browser front-end.

**18. Generic Data Structures Library**

- **Build:** `Stack[T]`, `Queue[T]`, `Set[T comparable]`, `LinkedList[T]`, an LRU cache, and `Map/Filter/Reduce` helpers.
- **Practices:** Type parameters, constraints, `comparable`, `cmp.Ordered`, `iter.Seq` range-over-func iterators, when not to use generics.
- **Stretch:** A generic concurrent-safe map and a priority queue using `container/heap`.

**19. Reverse Proxy / Load Balancer**

- **Build:** Forward requests to a pool of backends using round-robin, with health checks.
- **Practices:** `net/http/httputil.ReverseProxy`, background health-check goroutines, atomic counters, middleware, `context` propagation.
- **Stretch:** Weighted and least-connections strategies, plus a metrics endpoint.

## Tier 5: Advanced and Systems

*Goal: apply everything in larger, design-heavy projects.*

**20. Persistent Key-Value Store (Bitcask-style)**

- **Build:** An append-only log with an in-memory index, compaction, and crash recovery.
- **Practices:** File I/O with `os.File`, `encoding/binary`, checksums (`hash/crc32`), `mmap` or `io.ReaderAt`, locking, durability trade-offs (`fsync`).
- **Stretch:** Move to an LSM-tree design with memtables, SSTables, and bloom filters.

**21. Redis Clone**

- **Build:** A server speaking the RESP protocol that supports `GET`, `SET`, `DEL`, `EXPIRE`, lists, and pub/sub, testable with `redis-cli`.
- **Practices:** Protocol parsing, a custom TCP server, concurrency strategy, command dispatch, persistence (snapshots or AOF).
- **Stretch:** Transactions (`MULTI/EXEC`) and benchmarking with `redis-benchmark`.

**22. Interpreter for a Small Language**

- **Build:** Lexer → parser → AST → evaluator for something like *Monkey* (from *Writing an Interpreter in Go*) with variables, functions, closures, and arrays.
- **Practices:** Interfaces for AST nodes, maps for environments, recursion, extensive tests, error handling at scale.
- **Stretch:** Compile to bytecode and write a VM, then compare performance.

**23. Distributed Job Queue**

- **Build:** A server that accepts jobs over HTTP or gRPC, with workers that pull, execute, retry with backoff, and report status.
- **Practices:** gRPC and Protocol Buffers, leases/visibility timeouts, idempotency, retries, persistent storage, observability with metrics and tracing.
- **Stretch:** Scheduled and cron jobs, dead-letter queues, and a small dashboard.

**24. Raft Consensus (MIT 6.824 style)**

- **Build:** Leader election and log replication across nodes communicating via RPC.
- **Practices:** Heavy concurrency with timers and state machines, `net/rpc` or gRPC, careful locking, fault injection, and testing with network partitions.
- **Stretch:** Snapshots, then build a replicated KV store on top.

**25. Mini Git**

- **Build:** `init`, `add`, `commit`, `log`, `diff`, `branch`, and `checkout`, using a content-addressed object store.
- **Practices:** SHA-1/SHA-256 hashing, `compress/zlib`, tree structures, diff algorithms (Myers), filesystem manipulation, a robust CLI design.
- **Stretch:** Merge with conflict detection, and `push`/`pull` between two local repos.

## Tips for Getting the Most Out of These

- **Write tests from Tier 2 onward.** Table-driven tests and `go test -race` are central to how Go is written.
- **Use the standard library first.** Reach for third-party packages only after you understand what the stdlib gives you.
- **Run the tooling regularly:** `gofmt`, `go vet`, `staticcheck`, and `pprof` for profiling.
- **Read other people's Go.** After each project, compare your design to similar code in the stdlib or well-known repos (`net/http`, `bbolt`, `etcd`).
- **Don't skip the stretch goals in Tiers 3 and 5.** That's usually where the real lessons about leaks, races, and error handling show up.
