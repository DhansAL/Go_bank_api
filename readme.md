migrate -path db/migration -database "postgres://postgres:1234@localhost:5432/simplebank?sslmode=disable" -verbose up

go test -v -cover ./...