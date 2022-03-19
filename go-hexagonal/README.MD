## Run - Docker
    - docker-compose up -d
    - docker exec -it appproduct bash
### commands
    - go mod init github.com/mrpsousa/go-hexagonal

### installs
    - go get -u github.com/stretchr/testify

### command line
    - mockgen -destination=application/mocks/application.go -source=application/product.go

### tests
    - go test ./...
