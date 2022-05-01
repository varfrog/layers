run:
	DBUSER=app DBPASS=pass DBHOST=127.0.0.1 DBPORT=3306 go run ./cmd/main.go

test:
	go test -v ./...
