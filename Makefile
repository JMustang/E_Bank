postgres:
	docker run --name docker -p 5432:5432 -e POSTGRES_USER=docker -e POSTGRES_PASSWORD=docker -d postgres:12-alpine

createdb:
	docker exec -it docker createdb --username=docker --owner=docker Simple_Bank

dropdb:
	docker exec -it Simple_Bank

migrateup:
	migrate -path db/migration -database "postgresql://docker:docker@localhost:5433/Simple_Bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://docker:docker@localhost:5433/Simple_Bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/JMustang/E_Bank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock