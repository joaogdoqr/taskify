include .env
export $(shell sed 's/=.*//' .env)

migration_run:
	tern migrate --migrations ./store/migrations --config ./store/migrations/tern.conf

migration_create:
	cd ./store/migrations/ && tern new  $(name)

sqlc_generate:
	sqlc generate -f ./store/sqlc.yml

start:
	podman compose up -d
	air

test:
	go test ./...
