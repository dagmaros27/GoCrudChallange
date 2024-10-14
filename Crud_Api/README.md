# Go CRUD API with Gin

This is a simple CRUD API built using the Gin framework in Go. The API manages "Person" resources in an in-memory database and is structured following the MVC (Model-View-Controller) pattern. It supports common CRUD operations: creating, reading, updating, and deleting records, and is CORS-enabled for access by frontend applications hosted on different domains.

## Table of Contents

- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Testing the API](#testing-the-api)
- [CORS Support](#cors-support)
- [License](#license)

## Project Structure

```
crud_api/
│
├── main.go
├── controller/
│   └── controller.go
├── model/
│   └── person.go
└── route/
    └── route.go
```

- **`main.go`**: Initializes the server and sets up routes.
- **`controller/`**: Contains controller functions to handle API requests.
- **`model/`**: Defines the `Person` data model and data management functions.
- **`route/`**: Sets up routing and middleware (e.g., CORS).

## Endpoints

| Method | Endpoint      | Description                        |
| ------ | ------------- | ---------------------------------- |
| GET    | `/person`     | Retrieves all persons.             |
| GET    | `/person/:id` | Retrieves a specific person by ID. |
| POST   | `/person`     | Creates a new person.              |
| PUT    | `/person/:id` | Updates an existing person by ID.  |
| DELETE | `/person/:id` | Deletes an existing person by ID.  |

### Request and Response Examples

- **GET /person**

  ```json
  [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "John Doe",
      "age": 30,
      "hobbies": ["reading", "swimming"]
    }
  ]
  ```

- **POST /person**

  ```json
  // Request Body
  {
    "name": "Jane Doe",
    "age": 25,
    "hobbies": ["running", "painting"]
  }

  // Response
  {
    "id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Jane Doe",
    "age": 25,
    "hobbies": ["running", "painting"]
  }
  ```

- **PUT /person/:id**

  ```json
  // Request Body
  {
    "name": "Jane Doe Updated",
    "age": 26,
    "hobbies": ["yoga", "reading"]
  }

  // Response
  {
    "id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Jane Doe Updated",
    "age": 26,
    "hobbies": ["yoga", "reading"]
  }
  ```

- **DELETE /person/:id**
  - Response: `204 No Content`

### Error Responses

- **404 Not Found**: Returned when a person is not found or an endpoint does not exist.
- **400 Bad Request**: Returned when the request body is invalid.
- **500 Internal Server Error**: Returned for unexpected server errors.

## Installation

1. Make sure you have [Go](https://golang.org/doc/install) installed.
2. Install the Gin framework:
   ```bash
   go get -u github.com/gin-gonic/gin
   ```

## Running the Application

To run the API server, navigate to the project root and run:

```bash
go run main.go
```

The server will be accessible at `http://localhost:8080`.

## Testing the API

You can test the API using tools like [Postman](https://www.postman.com/) or `curl`.

Examples:

- **GET all persons**

  ```bash
  curl -X GET http://localhost:8080/person
  ```

- **POST a new person**

  ```bash
  curl -X POST http://localhost:8080/person \
    -H "Content-Type: application/json" \
    -d '{"name": "John Smith", "age": 40, "hobbies": ["hiking", "gardening"]}'
  ```

- **PUT to update a person**

  ```bash
  curl -X PUT http://localhost:8080/person/123e4567-e89b-12d3-a456-426614174000 \
    -H "Content-Type: application/json" \
    -d '{"name": "John Smith Updated", "age": 41, "hobbies": ["cycling"]}'
  ```

- **DELETE a person**
  ```bash
  curl -X DELETE http://localhost:8080/person/123e4567-e89b-12d3-a456-426614174000
  ```

## CORS Support

CORS is enabled to allow frontend applications hosted on different domains to access the API. The following headers are set:

- `Access-Control-Allow-Origin`: `*`
- `Access-Control-Allow-Methods`: `GET, POST, PUT, DELETE, OPTIONS`
- `Access-Control-Allow-Headers`: `Content-Type`
