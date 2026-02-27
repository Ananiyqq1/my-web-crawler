# 🕷️ Web Crawler

[![CI](https://github.com/Ananiyqq1/my-web-crawler/actions/workflows/ci.yml/badge.svg)](https://github.com/Ananiyqq1/my-web-crawler/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![FastAPI](https://img.shields.io/badge/FastAPI-0.115-009688?logo=fastapi&logoColor=white)](https://fastapi.tiangolo.com)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![Redis](https://img.shields.io/badge/Redis-7-DC382D?logo=redis&logoColor=white)](https://redis.io)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A high-performance, distributed web crawler built with **Go** for the crawling engine and **FastAPI** (Python) for the management REST API. Designed to be polite, concurrent, and scalable.

---

## ✨ Features

- 🚀 **Concurrent crawling** with configurable worker pools (Go goroutines)
- 🤖 **robots.txt compliance** — respects directives and crawl-delay
- ⏱️ **Per-domain rate limiting** via token bucket algorithm
- 🔗 **Link extraction & URL normalization** with depth-limited BFS
- 📡 **REST API** (FastAPI) to start, stop, and monitor crawl jobs
- 📊 **WebSocket live feed** for real-time crawl progress
- 🗄️ **PostgreSQL** for persistent storage of jobs and pages
- 📨 **Redis** message queue for decoupled job distribution
- 🐳 **Dockerized** — one command to run everything
- 📈 **Prometheus + Grafana** metrics and dashboards

---

## 🏗️ Architecture

```
┌──────────────┐       ┌─────────────┐       ┌──────────────────┐
│              │       │             │       │                  │
│   Client /   │──────▶│  FastAPI    │──────▶│     Redis        │
│   Dashboard  │◀──────│  (Python)   │       │  (Job Queue)     │
│              │  WS   │             │       │                  │
└──────────────┘       └──────┬──────┘       └────────┬─────────┘
                              │                       │
                              │                       ▼
                       ┌──────▼──────┐       ┌──────────────────┐
                       │             │       │                  │
                       │ PostgreSQL  │◀──────│   Go Crawler     │
                       │  (Storage)  │       │   (Workers)      │
                       │             │       │                  │
                       └─────────────┘       └──────────────────┘
```

---

## 📂 Project Structure

```
my-web-crawler/
├── crawler/                 # Go crawling engine
│   ├── cmd/                 # CLI entry points
│   │   ├── crawler/         # Standalone crawler CLI
│   │   └── worker/          # Redis consumer worker
│   ├── internal/            # Private packages
│   │   ├── fetcher/         # HTTP client
│   │   ├── parser/          # HTML parsing & link extraction
│   │   ├── robots/          # robots.txt handler
│   │   ├── ratelimit/       # Per-domain rate limiter
│   │   ├── worker/          # Goroutine worker pool
│   │   ├── storage/         # Database layer (pgx)
│   │   └── queue/           # Redis queue consumer
│   ├── migrations/          # Go database migrations
│   ├── go.mod
│   └── go.sum
├── api/                     # FastAPI management service
│   ├── main.py              # App entry point
│   ├── routers/             # Route handlers
│   ├── models/              # Pydantic schemas
│   ├── db/                  # SQLAlchemy models & CRUD
│   ├── services/            # Business logic & queue
│   ├── tests/               # pytest test suite
│   ├── alembic/             # Database migrations
│   ├── requirements.txt
│   └── Dockerfile
├── infra/                   # Infrastructure configs
│   ├── grafana/             # Grafana dashboard JSON
│   ├── prometheus/          # Prometheus config
│   └── docker-compose.yml   # Local dev stack
├── docs/                    # Documentation
│   ├── architecture.md
│   ├── data-models.md
│   ├── tech-stack.md
│   ├── api-spec.yaml
│   ├── setup.md
│   └── deployment.md
├── scripts/                 # Helper scripts
├── .github/
│   ├── ISSUE_TEMPLATE/      # Issue templates
│   └── workflows/           # CI/CD pipelines
├── .gitignore
├── LICENSE
└── README.md
```

---

## 🚀 Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.23+](https://go.dev/dl/) (for local crawler development)
- [Python 3.11+](https://www.python.org/downloads/) (for local API development)

### Run Everything with Docker

```bash
# Clone the repo
git clone https://github.com/Ananiyqq1/my-web-crawler.git
cd my-web-crawler

# Start all services
docker compose -f infra/docker-compose.yml up -d

# API is now available at http://localhost:8000
# Swagger docs at http://localhost:8000/docs
```

### Start a Crawl Job

```bash
# Create a new crawl job
curl -X POST http://localhost:8000/crawl \
  -H "Content-Type: application/json" \
  -d '{"seed_urls": ["https://example.com"], "max_depth": 2}'

# Check job status
curl http://localhost:8000/crawl/{job_id}

# List crawled pages
curl http://localhost:8000/crawl/{job_id}/pages
```

---

## 🧪 Running Tests

### Go Crawler
```bash
cd crawler
go test ./... -race -cover
```

### FastAPI
```bash
cd api
pip install -r requirements.txt
pytest tests/ -v --cov
```

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/crawl` | Start a new crawl job |
| `GET` | `/crawl/{job_id}` | Get job status |
| `GET` | `/crawl/{job_id}/pages` | List crawled pages (paginated) |
| `GET` | `/pages/{page_id}` | Get specific page details |
| `DELETE` | `/crawl/{job_id}` | Stop/cancel a job |
| `WS` | `/crawl/{job_id}/stream` | Live progress feed |
| `GET` | `/health` | Health check |

Full OpenAPI spec: [docs/api-spec.yaml](docs/api-spec.yaml)

---

## 🛠️ Tech Stack

| Component | Technology |
|-----------|-----------|
| Crawler Engine | Go 1.23, goquery, temoto/robotstxt |
| Management API | Python 3.11, FastAPI, Pydantic v2 |
| Database | PostgreSQL 16 |
| Message Queue | Redis 7 |
| Migrations | golang-migrate (Go), Alembic (Python) |
| Containers | Docker, Docker Compose |
| CI/CD | GitHub Actions |
| Monitoring | Prometheus, Grafana |

---

## 🤝 Contributing

See [docs/contributing.md](docs/contributing.md) for guidelines.

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/your-feature`
3. Commit your changes: `git commit -m "feat: add your feature"`
4. Push to the branch: `git push origin feat/your-feature`
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Ananiya H/meskel** — [@Ananiyqq1](https://github.com/Ananiyqq1)
