# metatheif
A high-performance concurrent URL metadata fetcher written in Go (Golang).

ConcurFetch is a CLI tool that fetches metadata (title, status code, content type, etc.) from multiple URLs concurrently, demonstrating efficient goroutine and channel usage in Go. Built for learning purposes, it showcases:

✔ Concurrency patterns (worker pools, goroutines, channels)
✔ HTTP client optimization (timeouts, keep-alives)
✔ Structured output (JSON, CSV, or plaintext)
✔ Error handling & retries for robustness

Why?
Learn Go’s concurrency model in a practical project.

Compare performance vs. sequential fetching.

Demonstrate idiomatic Go (interfaces, structs, error handling).

