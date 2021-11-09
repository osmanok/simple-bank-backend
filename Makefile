postgres: 
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=YOURUSERNAME -e POSTGRES_PASSWORD=YOURPASSWORD -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb YOURDBNAME

dropdb:
    docker exec -it postgres12 psql -U YOURUSERNAME -c "DROP DATABASE IF EXISTS simplebank;"

migrateup: 
    migrate -path db/migration -database "postgresql://YOURUSERNAME:YOURPASSWORD@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown: 
    migrate -path db/migration -database "postgresql://YOURUSERNAME:YOURPASSWORD@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
    sqlc generate

test:
    go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test