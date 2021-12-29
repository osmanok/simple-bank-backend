postgres: docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER='postgres' -e POSTGRES_PASSWORD='password' -d postgres:12-alpine

createdb: docker exec -it postgres12 createdb -U postgres simplebank

dropdb: docker exec -it postgres12 psql -U postgres -c "DROP DATABASE IF EXISTS simplebank;"

migrateup: migrate -path db/migration -database "postgresql://postgres:@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown: migrate -path db/migration -database "postgresql://postgres:@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc: sqlc generate

test: go test -v -cover ./...

server: go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test main