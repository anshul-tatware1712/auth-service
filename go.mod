module auth-service

go 1.26.0

require github.com/joho/godotenv v1.5.1

require (
	github.com/go-chi/chi/v5 v5.5.0
	github.com/go-sql-driver/mysql v1.10.1
	github.com/golang-jwt/jwt/v5 v5.3.1
	golang.org/x/crypto v0.58.0
)

require filippo.io/edwards25519 v1.2.0 // indirect
