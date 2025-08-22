# BlogAggre-Gator 🐊

Guided project of [boot.dev](https://boot.dev) academy.

A straightforward RSS/blog aggregator written in Go.

## Project overview

Gator is a lightweight command-line application that fetches and aggregates blog or RSS feed posts. It uses [sqlc](https://sqlc.dev/) for type-safe SQL queries and stores results in a relational database.

The project is designed to be simple, extensible, and developer-friendly.

## Features

- Fetch and aggregate RSS/blog feeds.
- SQL database integration using sqlc for type-safe query handling.
- User accounts with feed following/unfollowing support.
- Minimal dependencies—pure Go logic and SQL-defined schema/query layer.
- Easily extensible into other Go projects.

## Installation

### Prerequisites

- Go 1.18 or higher

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/LucasSim0n/gator.git
   cd gator```

2. Install dependencies:
    
    ```bash
    go mod download
    ```
3. Configure your database schema and queries using sql/ and sqlc.yaml.

I used postgres and goose for the migrations.

4. Build the aplication:
    
    ```bash
    go build -o gator main.go
    ```

## Usage

Run the help command to see all the avaliable commands and their usage.
    
    ```bash
    gator help
    ```
