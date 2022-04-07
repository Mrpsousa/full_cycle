## Run - Docker
    - docker-compose up -d
    - docker exec -it appproduct bash
### commands
    - go mod init github.com/mrpsousa/go-hexagonal

### installs
    - go get -u github.com/stretchr/testify

### command line
    - mockgen -destination=application/mocks/application.go -source=application/product.go
    - sqlite3 sqlite.db
        - create table products(id string, name string, price float, status string);
        - .tables

### tests
    - go test ./...

### Run with cobra 
    - up container ...
    - docker exec -it appproduct bash
    - cobra init
    - go mod tidy
    - cobra add cli
    - go run main.go cli -a=create -n=Product -p=25.0
    