module simple_website

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
)

require internal/handlers v1.0.0
replace internal/handlers => ./internal/handlers
require internal/database v1.0.0
replace internal/database => ./internal/database
require internal/articles v1.0.0
replace internal/articles => ./internal/articles
