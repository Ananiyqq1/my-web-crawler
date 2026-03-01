# Data Models

This document defines the core entities used across the Go crawler and the Python API.

## Entities

### 1. CrawlJob

Represents a single request to crawl one or more websites.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | UUID | Primary Key |
| `seed_urls` | String[] | List of starting URLs |
| `max_depth` | Integer | How many levels of links to follow |
| `status` | Enum | `pending`, `running`, `completed`, `failed`, `cancelled` |
| `created_at` | Timestamp | When the job was created |
| `updated_at` | Timestamp | When the job was last updated |

### 2. Page

Represents a single crawled HTML page.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | UUID | Primary Key |
| `job_id` | UUID | Foreign Key to `CrawlJob` |
| `url` | String | The absolute URL of the page |
| `title` | String | The HTML `<title>` of the page |
| `html_content` | Text | The raw or cleaned HTML content (optional storage) |
| `http_status` | Integer | The HTTP response code (e.g., 200, 404) |
| `fetched_at` | Timestamp | When the page was fetched |
| `links` | String[] | List of outgoing absolute URLs found on the page |

### 3. QueueItem (Internal)

Used for the URL frontier during the crawl process.

| Field | Type | Description |
| :--- | :--- | :--- |
| `job_id` | UUID | Reference to the crawl job |
| `url` | String | URL to be visited |
| `depth` | Integer | Current depth level |
| `retries` | Integer | Number of failed attempts |

## Relationships

-   **One-to-Many**: A `CrawlJob` has many `Page` records.
-   **One-to-Many**: A `CrawlJob` has many `QueueItem` records (transient).
