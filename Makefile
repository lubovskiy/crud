BINARY=engine
test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} cmd/service/main.go

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

build:
	docker build -t go-clean-arch .

deploy:
	docker-compose up --build -d

delete:
	docker-compose down

migrage:
	migrate -path ./internal/migrations -database "postgres://postgres:postgres@localhost:5432/phonebook?sslmode=disable" up

.PHONY: test engine clean docker run stop proto migrate
