# Technology Stack

Rationale for the technology choices made in the Web Crawler project.

## Core Services

### 1. Go (Crawler Engine)
- **Why**: High concurrency performance via goroutines and lightweight threads.
- **Benefits**: 
  - Compiled language (fast execution).
  - Excellent networking standard library.
  - Low memory footprint for parallel I/O.
- **Key Libraries**:
  - `net/http`: For robust HTTP client operations.
  - `goquery`: For jQuery-like HTML parsing.
  - `golang.org/x/time/rate`: For per-domain rate limiting.
  - `pgx`: High-performance PostgreSQL driver.

### 2. FastAPI (Management API)
- **Why**: Modern, high-performance Python framework for building APIs.
- **Benefits**:
  - Built-in asynchronous support (`async`/`await`).
  - Automatic OpenAPI/Swagger documentation generation.
  - Pydantic for data validation and serialization.
- **Key Libraries**:
  - `SQLAlchemy`: Async ORM for PostgreSQL.
  - `redis-py`: For communication with Redis.
  - `pytest`: For unit and integration testing.

## Infrastructure

### 3. PostgreSQL (Persistence)
- **Why**: Relational database with support for complex queries and constraints.
- **Benefits**:
  - ACID compliance.
  - Powerful indexing for URL lookups.
  - JSONB support for flexible metadata storage if needed.

### 4. Redis (Queue & Pub/Sub)
- **Why**: Extremely fast, in-memory data store.
- **Benefits**:
  - Acts as a job queue (list-based).
  - Acts as a message broker for live WebSocket updates (Pub/Sub).
  - Lightweight and easy to scale.

### 5. Docker & Docker Compose
- **Why**: Simplified development and deployment.
- **Benefits**:
  - Guaranteed environment consistency.
  - Easy orchestration of multiple services (API, Crawler, DB, Redis).
