.PHONY: all
all: up

.PHONY: default up down

default: deploy

up:
	docker-compose up -d --build

down:
	docker-compose down