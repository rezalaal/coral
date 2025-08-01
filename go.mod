module coral

go 1.24.5

require github.com/lib/pq v1.10.9 // برای ارتباط با PostgreSQL

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.40.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
