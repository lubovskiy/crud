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

proto:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
			-I${GOPATH}/pkg/mod/google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55/googleapis \
     --go_out=. --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/*.proto


.PHONY: test engine clean docker run stop proto
