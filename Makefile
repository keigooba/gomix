#Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gomix
# BINARY_UINX=$(BINARY_NAME)_unix

all: run
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	# rm -f $(BINARY_UINX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/fukata/golang-stats-api-handler
	$(GOGET) github.com/smartystreets/goconvey
	$(GOGET) gopkg.in/ini.v1

# クロスコンパイル 実行環境にdocker使用
# build-linux:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UINX) -v
# docker-build:
# 	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/butbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UINX)" -v
