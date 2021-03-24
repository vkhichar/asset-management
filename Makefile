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
	


