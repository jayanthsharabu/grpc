protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
       	   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ${path}


build:
	go build -o $(output) $(path)

run:
	go run $(path)