# 🎬 Go Movie Project

This is a simple **Go REST API** built **without using a database**.  
It uses **structs** and **slices** to store, manage, and manipulate movie data in memory.

---

## 📌 Features

- CRUD operations on movies
- Simple routing using [Gorilla Mux](https://github.com/gorilla/mux)
- No database required — everything is stored in memory
- Easy to test using Postman

---

## 🛣️ API Routes

| Method | Route                | Description                           |
|--------|----------------------|---------------------------------------|
| GET    | `/`                  | Returns a basic welcome message       |
| GET    | `/movies`            | Returns a list of all movies          |
| GET    | `/movie/{id}`        | Returns a single movie by its ID      |
| DELETE | `/movie/{id}`        | Deletes a movie by its ID             |
| POST   | `/movie`             | Creates a new movie                   |
| PUT    | `/movie/{id}`        | Updates an existing movie by its ID   |

---

## 🚀 Project Setup

### ✅ Requirements

- **Go version 1.24.1** (or higher)

> You can check your Go version using:
> ```bash
> go version
> ```

---

### 📦 Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/Go-Movie-Project.git
   cd Go-Movie-Project
   
2. **Build the Project**:
   
   ```bash
   go build -o Go-Movie-Project ./cmd
   
3. **Run the Project**:
   
   ```bash
   ./Go-Movie-Project

4. **You should see**:

   ```bash
   Started server on port 3000

## 🔧 Testing the API

You can use [Postman](https://www.postman.com/) to test the API endpoints.

> Base URL: http://localhost:3000

Try sending requests to the above routes to create, fetch, delete, or update movies.

## 🧠 Notes

- All data is stored in-memory, meaning it resets when the server restarts.
- This project is ideal for learning Go fundamentals, HTTP routing, and struct/slice usage.

## 📂 Project Structure

```Go-Movie-Project/
├── cmd/             # Entry point of the application
│   └── main.go      # Main application file
├── internal/        # Internal packages
│   ├── handler/     # HTTP handlers for routing and request handling
│   ├── service/     # Business logic (e.g., ID generation, utilities)
│   └── model/       # Structs and types representing the data models
├── go.mod           # Go module file
├── go.sum           # Dependency checksum file
```

## 📬 Contact

Feel free to reach out. Happy coding! 😄

