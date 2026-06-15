# url-shortener-go

A URL shortener written in Go. Send it a valid URL, get back a short code. Visit the short code, get redirected to the original destination.

Built as a learning vehicle for Go and infrastructure engineering — the app is intentionally simple so the focus stays on how it's built and operated, not what it does.

---

## Architecture

```
POST /send-raw-text   →  validates URL, generates 6-char code, stores mapping
GET  /{code}          →  looks up code, redirects to original URL
GET  /health          →  returns { status: "ok" } for health checks
```

**Why Go?**
Go compiles to a single static binary that runs the same everywhere. No runtime, no dependencies, no "works on my machine." That property makes it a natural fit for containerised infrastructure — and it's the language Kubernetes, Docker, and Terraform are written in.

**Why the standard library and not a framework?**
Go's `net/http` is production-grade on its own. Most infrastructure tooling at companies like Cloudflare and Google uses it directly. Learning the standard library first means understanding what frameworks are actually abstracting.

**Why a health endpoint?**
When this app runs inside Docker, Kubernetes, or behind a load balancer — nothing can see inside the container. The `/health` endpoint is how the outside world asks "are you alive?" Monitoring tools, orchestrators, and load balancers all rely on it to decide whether to send traffic to a container or restart it.

**Current storage: in-memory hash map**
URLs are stored in a `map[string]string` — fast, simple, and intentionally temporary. The store resets on every restart. A persistent database (PostgreSQL) is the next addition.

---

## Running locally

**Prerequisites:** Go 1.22+

```bash
git clone https://github.com/miraclemenikelechi/url-shortner-go
cd url-shortner-go
go run main.go
```

---

## Running with Docker

**Build:**
```bash
docker build -t urlshortener .
```

**Run:**
```bash
docker run -p 8649:8649 urlshortener
```

The Docker build uses a **multi-stage build**:
- Stage 1 (`golang:1.26`) compiles the binary
- Stage 2 (`alpine`) runs it

The final image carries only the compiled binary — not the Go toolchain. This keeps the image small and the attack surface minimal. First build takes ~2 minutes downloading layers. Subsequent builds run in under 10 seconds thanks to layer caching.

---

## Testing the endpoints

```bash
# Health check
curl http://localhost:8649/health

# Shorten a URL
curl -X POST http://localhost:8649/send-raw-text \
  -H "Content-Type: application/json" \
  -d '{"raw_url": "https://google.com"}'

# Follow the redirect (replace CODE with the value returned above)
curl -L http://localhost:8649/CODE
```

---

## What I learned

Coming from Node.js, the biggest adjustment was Go's type system and pattern for handling HTTP. In Node you reach for `req.body` and duck-type everything. In Go you define a struct for every request and response shape — more upfront work, but the compiler catches mistakes before they reach production.

URL validation was also more intentional — checking both the scheme (`http`/`https`) and the host explicitly, rather than relying on a library to decide what "valid" means.

---

## What's next

- [ ] PostgreSQL persistence via Docker Compose
- [ ] Terraform to provision the infrastructure
- [ ] GitHub Actions CI/CD pipeline
- [ ] Prometheus metrics + Grafana dashboard
- [ ] Kubernetes deployment