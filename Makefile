postgres:
	docker run --name pgtestdb -p 5500:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it pgtestdb createdb --username=postgres --owner=postgres bariodu 
dropdb:
	docker exec -it pgtestdb dropdb --username=postgres --owner=postgres bariodu 
migrate-up:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5500/bariodu?sslmode=disable" -verbose up

migrate-down:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5500/bariodu?sslmode=disable" -verbose down

sqlc-generate:
	@sqlc -f ./src/db/sqlc.yaml generate
	@echo "Successfully generated."

create-migration:
	migrate create -ext sql -dir src/db/migrations -seq $(migration_name)

test:
	go test -v -cover ./...

test-no-cache:
	go test -v -cover -count=1 ./...

.PHONY: migrate-up sqlc-generate __migrate-down postgres

