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

**Storage: PostgreSQL via Docker Compose**
URLs are persisted in a PostgreSQL 18 database running as a Docker container alongside the app. Configuration is injected via environment variables — no hardcoded credentials. A named Docker volume ensures data survives container restarts. Database migrations are managed with `golang-migrate` and run automatically on startup, so the schema is always in sync with the code.

---

## Running locally

**Prerequisites:** Docker + Docker Compose

```bash
git clone https://github.com/miraclemenikelechi/url-shortener-go
cd url-shortener-go
cp .env.example .env
docker compose up --build
```

One command starts everything — the Go app and the PostgreSQL database. Migrations run automatically on startup.

To stop without losing data:
```bash
docker compose down
```

To stop and wipe the database:
```bash
docker compose down -v
```

---

## How Docker is used

The build uses a **multi-stage Dockerfile**:
- Stage 1 (`golang:1.26`) compiles the binary
- Stage 2 (`alpine`) runs it

The final image carries only the compiled binary — not the Go toolchain. This keeps the image small and the attack surface minimal. First build takes ~2 minutes downloading layers. Subsequent builds run in under 10 seconds thanks to layer caching.

Docker Compose wires the app and database together on an internal network. The app reaches Postgres via the service name `postgres:5432` — no hardcoded IPs, no host machine ports exposed unnecessarily.

---

## Infrastructure as Code — Terraform

The `terraform/` folder provisions the entire infrastructure the app runs on. Locally that means Docker containers, images, and networks. On a real cloud it would mean servers, managed databases, firewalls, and DNS — the same Terraform code, different variable values per environment.

**Docker Compose vs Terraform:**
Docker Compose is the runtime config — it describes how containers run together. Terraform is the infrastructure layer underneath — it provisions and validates what needs to exist before the app can run at all. In production they work at different levels: Terraform creates the servers and networking, Docker Compose (or Kubernetes) runs the containers on top of them.

**Why this matters for multiple environments:**
The same Terraform configuration can provision dev, staging, and production with different variable files. One codebase, three environments, no configuration drift. What runs in staging is identical to production — just smaller.

```bash
# Provision everything
cd terraform
terraform init
terraform apply

# Tear down everything cleanly
terraform destroy
```

Sensitive values like database passwords are declared as `sensitive = true` variables and stored in `terraform.tfvars` which is never committed to git.

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

Adding the database introduced `context.Context` properly for the first time. In Go, every database call carries a context — a cancellation signal that travels with the request. If the HTTP client disconnects, the database query cancels automatically. It lives up to everything people say about it.

On Terraform: the concept didn't fully click until I understood what it's actually for. It's not a replacement for Docker Compose — they operate at different levels. Terraform provisions what needs to exist. Docker runs what lives inside it. The real value shows at scale — the same configuration file can provision dev, staging, and production without touching a single server manually.

On Docker: I came into this having avoided Docker for a long time. Building this changed that. When you start thinking about distributed systems and how services actually run in production, Docker isn't optional — it's the bedrock. Without it you're signing up for a battle you can't even start.

---

## What's next

- [x] URL shortening with validation
- [x] PostgreSQL persistence via Docker Compose
- [x] Automatic database migrations on startup
- [x] Terraform infrastructure as code (Docker provider locally)
- [ ] GitHub Actions CI/CD pipeline
- [ ] Prometheus metrics + Grafana dashboard
- [ ] Kubernetes deployment
