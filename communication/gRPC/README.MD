run:
    protoc --proto_path=proto proto/*.proto --go_out=pb
    protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
    go run cmd/server/server.go
    go run cmd/client/client.go
dependecias: 
    sudo apt install golang-goprotobuf-dev
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    go get -u google.golang.org/grpc