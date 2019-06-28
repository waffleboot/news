
.PHONY: client storage

client:	cmd/client/* client/*
	docker build -f cmd/client/Dockerfile -t yangand/news/client .

storage: cmd/storage/* storage/*
	docker build -f cmd/storage/Dockerfile -t yangand/news/storage .

all: client storage

up:
	docker-compose up -d
	
down:
	docker-compose down --remove-orphans