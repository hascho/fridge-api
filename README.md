# Go Fridge Inventory API

A simple REST API to track items in your fridge, built with Go, Gin, and PostgreSQL.  
This project is for learning purposes and as a portfolio piece.

---

## Features

- Manage categories (Dairy, Fruits, Vegetables, etc.)
- Track items with quantity, unit, expiry date, and notes
- Swagger documentation available at `/swagger/index.html`
- Modular design following service + repository + handler pattern

---

## Prerequisites

- [Go 1.25+](https://golang.org/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/compose/install/)
- `.env` file for configuration (see `.env.example`)

---

## Setup

1. Clone the repository:

```bash
git clone https://github.com/hascho/go-fridge.git
cd go-fridge
````

2. Copy the environment file:

```bash
cp .env.example .env
```

3. Build and run the app:

```bash
make run
```

4. The API should now be running at:

```bash
http://localhost:8080
```

## API Endpoints

### Categories

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | /api/categories | Create a new category |
| GET    | /api/categories | List all categories |
| GET    | /api/categories/:id | Get a single category |
| PUT    | /api/categories/:id | Update a category |
| DELETE | /api/categories/:id | Delete a category |

### Items

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | /api/items | Create a new item |
| GET    | /api/items | List all items |
| GET    | /api/items/:id | Get a single item |
| PUT    | /api/items/:id | Update an item |
| DELETE | /api/items/:id | Delete an item |

**Filters:**  
You can filter items using query parameters:

- `category_id` — filter by category  
- `expired` — `true` or `false`  
- `expiring_within` — items expiring within N days  

Example request:

`GET /api/items?category_id=1&expired=false`

---

## Swagger Documentation

Swagger UI is available at:

`http://localhost:8080/swagger/index.html`

---

## Makefile Commands

- `make run` — build & run the app with Docker Compose  
- `make down` — stop and remove containers  
- `make tidy` — tidy Go modules  
- `make swag` — regenerate Swagger docs

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
