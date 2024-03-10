module github.com/Jain2003/tri-nit-2024/backend

go 1.21.5

require (
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/cors v1.2.1
)

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/joho/godotenv v1.5.1
	internal/database v1.0.0
)

replace internal/database => ./internal/database
