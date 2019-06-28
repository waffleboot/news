
.PHONY: client storage messaging

client:	cmd/client/* client/* messaging/*
	docker build -f cmd/client/Dockerfile -t yangand/news/client .

storage: cmd/storage/* messaging/*
	docker build -f cmd/storage/Dockerfile -t yangand/news/storage .

all: client storage

up:
	docker-compose up -d
	
down:
	docker-compose down --remove-orphans