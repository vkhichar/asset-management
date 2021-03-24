FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 
    
ARG APP_PORT 
ARG DB_HOST 
ARG DB_PORT 
ARG DB_USERNAME 
ARG DB_PASSWORD 
ARG DB_NAME 
ARG EVENT_SERVICE_URL  
ARG EVENT_API_TIMEOUT 
ARG JWT_SECRET 
ARG TOKEN_EXPIRY

WORKDIR /asset_management

COPY go.mod .
COPY go.sum .
RUN go mod download

#copy my code from where I am right now (i.e. current dir /github.com/vkhichar/asset_management)
#to working dir specified in workdir command
COPY . .  

#run go build. write output in main
#run all packages (specified by .) 
RUN go build -o main .

#now dir structure inside docker image becomes /asset_management/main (binary is in main)


EXPOSE 9000

CMD ["/asset_management/main"]