
# Distributed Reddit-Style Social Platform (Advanced)

## Tech Stack
Go, PostgreSQL, Redis, Docker, Docker Compose

## Features
- Time-decay feed ranking (Reddit-style hot score)
- Threaded comment trees
- Upvote / Downvote system
- Redis-backed worker pools
- Load testing with k6 (2k+ RPS)

## Run
docker compose up --build

API: http://localhost:8080
