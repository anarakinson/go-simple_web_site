include .$(PWD)/.env

.PHONY: start_app
start_app:
	docker-compose up

.PHONY: build_app
build_app:
	docker-compose build

.PHONY: stop_app
stop_app:
	docker-compose down

.DEFAULT_GOAL := start_app
