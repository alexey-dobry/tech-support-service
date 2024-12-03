module github.com/alexey-dobry/tech-support-platform/internal/services/req_user_service

replace github.com/alexey-dobry/tech-support-platform/internal/pkg => ../../pkg

go 1.23.2

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/alexey-dobry/tech-support-platform/internal/pkg v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
)
