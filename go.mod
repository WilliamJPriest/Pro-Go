module github.com/williamjPriest/HTMXGO

go 1.21.2

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.5.1
	github.com/williamjPriest/HTMXGO/database v0.0.0-00010101000000-000000000000
	github.com/williamjPriest/HTMXGO/middlewares v0.0.0-00010101000000-000000000000
	github.com/williamjPriest/HTMXGO/models v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.14.0
)

require github.com/lib/pq v1.10.9 // indirect

replace github.com/williamjPriest/HTMXGO/database => ./database

replace github.com/williamjPriest/HTMXGO/middlewares => ./middlewares

replace github.com/williamjPriest/HTMXGO/models => ./models

replace github.com/williamjPriest/HTMXGO/routes => ./routes

replace github.com/williamjPriest/HTMXGO/controllers => ./controllers
