include .$(PWD)/.env

.PHONY: db
db:
	docker run --name simple_website_db -e MYSQL_ROOT_PASSWORD=${MYSQL_PASSWORD} -e MYSQL_DATABASE=simple_website -p 3307:3306 -d --rm mysql:latest

.PHONY: migrations
migrations:
	migrate -path ./db_schema -database 'mysql://root:${MYSQL_PASSWORD}@tcp(127.0.0.1:3307)/simple_website?query' up

.PHONY: build
build:
	go build -v

.PHONY: start
start:
	go run main.go

.PHONY: stop
stop:
	docker stop simple_website_db

.DEFAULT_GOAL := start
