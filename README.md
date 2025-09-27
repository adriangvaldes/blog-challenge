# Go Blog API

A simple and lightweight RESTful API built with Go for managing blog posts. This project serves as a practical example of building a backend service using Go's standard library and the popular `gorilla/mux` router.

## Features

- **CRUD Operations:** Full support for creating, reading, updating, and deleting blog posts.
- **RESTful Endpoints:** Clean and predictable API endpoints.
- **In-Memory Storage:** Currently uses an in-memory slice for data storage (database integration is a planned feature).
- **JSON Communication:** All requests and responses are in JSON format.

## Technologies Used

- [Go](https://golang.org/)
- [gorilla/mux](https://github.com/gorilla/mux) for HTTP routing

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) version 1.18 or higher.

### Installation & Running

1.  **Clone the repository (example):**

    ```bash
    git clone [https://github.com/your-username/blog-api.git](https://github.com/your-username/blog-api.git)
    cd blog-api
    ```

2.  **Install dependencies:**
    The project uses Go Modules. The `gorilla/mux` dependency will be downloaded automatically when you build or run the project. You can also install it manually:

    ```bash
    go get [github.com/gorilla/mux](https://github.com/gorilla/mux)
    ```

3.  **Run the server:**
    ```bash
    go run main.go
    ```
    The server will start and listen on `http://localhost:8080`.

## API Endpoints

The base URL for all endpoints is `http://localhost:8080`.

| Method   | Endpoint      | Description                    | Request Body (JSON)                                  | Success Response            |
| :------- | :------------ | :----------------------------- | :--------------------------------------------------- | :-------------------------- |
| `GET`    | `/posts`      | Retrieves a list of all posts. | `None`                                               | `200 OK` with post array    |
| `POST`   | `/posts`      | Creates a new post.            | `{"title": "...", "content": "...", "author_id": 1}` | `201 Created` with new post |
| `GET`    | `/posts/{id}` | Retrieves a single post by ID. | `None`                                               | `200 OK` with single post   |
| `PUT`    | `/posts/{id}` | Updates an existing post.      | `{"title": "...", "content": "..."}`                 | `200 OK` with updated post  |
| `DELETE` | `/posts/{id}` | Deletes a post by ID.          | `None`                                               | `204 No Content`            |

### Example cURL Commands

- **Get all posts:** `curl http://localhost:8080/posts`
- **Create a post:** `curl -X POST -H "Content-Type: application/json" -d '{"title":"New Post", "content":"Content here"}' http://localhost:8080/posts`
- **Get post with ID 1:** `curl http://localhost:8080/posts/1`
- **Update post with ID 1:** `curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Title"}' http://localhost:8080/posts/1`
- **Delete post with ID 1:** `curl -X DELETE http://localhost:8080/posts/1`

## Future Improvements

- [ ] Replace in-memory storage with a persistent SQL database (e.g., PostgreSQL or SQLite).
- [ ] Add user registration and JWT-based authentication.
- [ ] Implement unit and integration tests.
