module github.com/williamjPriest/HTMXGO

go 1.21.2

replace github.com/williamjPriest/HTMXGO/database => ./database

replace github.com/williamjPriest/HTMXGO/middlewares => ./middlewares

replace github.com/williamjPriest/HTMXGO/models => ./models

require (
	github.com/joho/godotenv v1.5.1
	github.com/williamjPriest/HTMXGO/middlewares v0.0.0-00010101000000-000000000000
	github.com/williamjPriest/HTMXGO/routes v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/williamjPriest/HTMXGO/database v0.0.0-00010101000000-000000000000 // indirect
	github.com/williamjPriest/HTMXGO/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/williamjPriest/HTMXGO/utils v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/crypto v0.14.0 // indirect
)

replace github.com/williamjPriest/HTMXGO/utils => ./utils

replace github.com/WilliamJPriest/Go-WebScrapper/models => ./models

replace github.com/WilliamJPriest/Go-WebScrapper/database => ./database

replace github.com/williamjPriest/HTMXGO/routes => ./routes

replace github.com/williamjPriest/HTMXGO/controllers => ./controllers
