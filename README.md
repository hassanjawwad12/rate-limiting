# Rate Limiting in Go

Rate limiting is a technique used to control the amount of incoming and outgoing traffic to or from a network, server, or service. This repository demonstrates three different approaches to implementing rate limiting in Go applications.

<img src="token-bucket.png"/>

## What is Rate Limiting?

Rate limiting restricts how many requests a client can make to an API within a specified time period. When a client exceeds the allowed limit, the server typically responds with a 429 (Too Many Requests) status code, indicating that the client should retry after a certain period.

## Why is Rate Limiting Important?

1. **Prevent DoS Attacks**: Rate limiting is a crucial defense mechanism against Denial of Service (DoS) attacks, where attackers flood a service with excessive requests to overwhelm it.

2. **Resource Management**: It helps manage server resources efficiently by preventing any single client from consuming too much bandwidth or processing power.

3. **Cost Control**: For services that rely on pay-per-use APIs, rate limiting helps control costs by capping usage.

4. **Improved Reliability**: By preventing traffic spikes, rate limiting ensures consistent performance for all users.

5. **Fair Usage**: It ensures fair distribution of resources among all clients, preventing any single client from monopolizing the service.

## Rate Limiting Techniques Implemented

### 1. Per-client Rate Limiting

This approach assigns a separate rate limit to each client, typically identified by IP address or API key. It allows for more granular control over resource allocation and prevents individual clients from affecting others.

Key features:
- Individual tracking of request counts per client
- Customizable limits for different client tiers
- Isolation of client behavior

### 2. Token Bucket Algorithm

The Token Bucket algorithm is a flexible rate limiting mechanism that works by filling a virtual "bucket" with tokens at a constant rate. Each request consumes one token, and if the bucket is empty, the request is rejected.

Key features:
- Allows for bursts of traffic (up to the bucket size)
- Smooth rate limiting over time
- Configurable token refill rate and bucket capacity
- Simple and efficient implementation

### 3. Tollbooth Library

This implementation uses the [Tollbooth](https://github.com/didip/tollbooth) library, a simple middleware for rate limiting HTTP requests in Go applications.

Key features:
- Easy integration with standard Go HTTP servers
- Configurable rate limits and time windows
- Custom response messages for rate-limited requests
- IP-based client identification
- Middleware approach that can be applied to specific routes

## Getting Started

### Prerequisites
- Go 1.16 or higher

### Installation
```bash
git clone https://github.com/yourusername/rate-limiting.git
go get 

### Running the Examples

# Run the per-client rate limiting example
cd per-client
go run main.go

# Run the token bucket example
cd ../token-bucket
go run main.go

# Run the tollbooth example
cd ../toolbooth
go run main.go

# Test with curl
`for i in {1..10}; do curl -i http://localhost:8080/ping; done`
# Test with parallel requests
`for i in {1..10}; do curl -s http://localhost:8080/ping &done`