GOCMD=go 
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
BINARY_NAME=asset-management
File_Path=/home/josh/go1.15.8.linux-amd64/go/src/github.com/vkhichar/asset-management


build:
	$(GOBUILD) -o $(BINARY_NAME) -v


test:
	$(GOTEST)  ./service/
	$(GOTEST)  ./handler/
	$(GOTEST)  ./repository/

testeach:
	$(GOTEST) -v ./service/
	$(GOTEST) -v ./handler/
	$(GOTEST) -v ./repository/

clean:
	$(GOCLEAN)

fmt:
	$(GOFMT) $(File_Path)
	


