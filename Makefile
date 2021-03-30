GOCMD=go 
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt



build:
	$(GOBUILD) 
test:
	$(GOTEST)  ./service/
	$(GOTEST)  ./handler/
	$(GOTEST)  ./repository/
fmt:
	$(GOFMT) ./...
migrate:
	go run main.go migrate
startapp:
	go run main.go startapp
rollback:
	go run main.go rollback



	


