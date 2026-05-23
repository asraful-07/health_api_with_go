**README.md**
**Go** is a programming language designed for building simple, fast, and reliable software. It was created at Google in 2007 and has since become one of the most popular programming languages in the world.The Go programming language is known for its simplicity, efficiency, and strong support for concurrent programming. It has a clean syntax and a powerful standard library that makes it easy to build a wide range of applications, from web servers to command-line tools.

# Project Structure

```bash
project-root/
│
├── cmd/
│   └── main.go
│
├── config/
│   └── db.go
│
├── internal/
│   │
│   ├── user/
│   │   │
│   │   ├── handler/
│   │   │   └── user_handler.go
│   │   │
│   │   ├── service/
│   │   │   └── user_service.go
│   │   │
│   │   ├── repository/
│   │   │   └── user_repository.go
│   │   │
│   │   ├── model/
│   │   │   └── user_model.go
│   │   │
│   │   └── routes/
│   │       └── user_routes.go
│   │
│   └── middleware/
│       └── middleware.go
│
├── pkg/
│   └── utils/
│       └── response.go
│
├── .env
├── go.mod
├── go.sum
└── README.md


# Recommended Structure two
my-app/
├── cmd/
│   └── main.go              ← app entry point, wires everything together
├── internal/
│   ├── domain/
│   │   ├── user.go          ← User struct (entity)
│   │   ├── repository.go    ← UserRepository interface
│   │   └── usecase.go       ← UserUsecase interface
│   ├── repository/
│   │   └── postgres_user_repo.go  ← implements UserRepository using pgx
│   ├── usecase/
│   │   └── user_usecase.go  ← business logic, implements UserUsecase
│   └── delivery/
│       └── http/
│           ├── user_handler.go  ← HTTP handlers (your current handlers)
│           └── router.go        ← mux setup, route registration
├── pkg/
│   └── db/
│       └── postgres.go      ← connectDB() lives here
├── .env
├── go.mod
└── go.sum
```
