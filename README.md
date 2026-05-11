# ratelimiter-go

[![Go Reference](https://pkg.go.dev/badge/dmx3377/ratelimiter-go.svg)](https://pkg.go.dev/dmx3377/ratelimiter-go)

**ratelimiter-go** is a high-performance, thread-safe rate-limiting middleware for [Go](https://go.dev) applications. 

It utilises the [Token Bucket algorithm](https://grokipedia.com/page/Token_bucket) to provide a *"leaky bucket"* style of traffic shaping, ensuring your services remain resilient against bursts and brute-force attempts.

Unlike traditional limiters that require background goroutines to refill tokens, this library uses lazy evaluation, calculating token availability only when a request arrives—to maximize CPU efficiency and minimize memory footprint.

## Core Features

* **Token Bucket Algorithm:** Provides granular control over request flow with support for "burst" capacity.
* **Zero Dependencies:** Built exclusively with the Go standard library.
* **Thread-Safe Architecture:** Uses `sync.RWMutex` to handle high-concurrency environments without data races.
* **Proxy-Aware IP Extraction:** Automatically identifies real client IPs behind Cloudflare, Nginx, or AWS ALBs by parsing `X-Forwarded-For` headers.
* **Framework Agnostic:** Compatible with `net/http`, Gin, Echo, Fiber, and Chi.

## Installation

`go get github.com/dmx3377/ratelimiter-go`

## Architecture & How it Works

The library operates on two main levels:
1.  **The Bucket:** Each unique key (e.g., an IP address) is assigned a bucket that holds a set number of "tokens."
2.  **The Limiter:** A central manager that maps keys to buckets and orchestrates the middleware logic.

When a request hits the middleware, the library calculates the time elapsed since the last request and "refills" the bucket based on the defined `fillRate`. If the bucket has at least one token, the request proceeds and one token is consumed.

## Usage Examples

### 1. Basic Standard Library *(net/http)*
The simplest way to protect your routes using the standard library:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/dmx3377/ratelimiter-go"
)

func main() {
	// 5 tokens max (burst), refilling at 1 token per second
	limiter := ratelimit.NewLimiter(5, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request successful!")
	})

	// Wrap your entire router
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", limiter.Middleware(mux))
}
```

### 2. Integration with Gin
Because the middleware follows the standard `http.Handler` pattern, integrating with Gin is straightforward:

```go
func main() {
    r := gin.Default()
    limiter := ratelimit.NewLimiter(10, 2)

    // Convert standard middleware to Gin middleware
    r.Use(func(c *gin.Context) {
        if !limiter.Allow(c.ClientIP()) {
            c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
            return
        }
        c.Next()
    })

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Pong!"})
    })
    r.Run()
}

```

## Configuration

| Parameter | Type | Default | Description |
| --- | --- | --- | --- |
| `capacity` | `float64` | Required | Maximum tokens in a bucket (max burst size). |
| `fillRate` | `float64` | Required | Tokens added per second. |

## Best Practices

* **Choosing Capacity:** Set your `capacity` to reflect the maximum burst you can handle at once. A higher capacity allows users to send several requests in quick succession if they have been inactive.
* **Memory Management:** The current version (v1.0.0) keeps buckets in memory. For extremely high-traffic sites with millions of unique IPs, monitor memory usage as each unique key creates a small `Bucket` struct.
* **Proxies:** Ensure your load balancer is configured to pass the `X-Forwarded-For` header; otherwise, the limiter may inadvertently rate-limit the load balancer itself rather than the end-users.

## License

This project is licensed under the **Apache License 2.0**. See the [LICENSE](LICENSE) file for details.

---

**Contributing:** Pull requests are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.
