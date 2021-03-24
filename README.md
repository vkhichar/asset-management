# asset-management
Asset management system for an organisation/business

# How to build, run
```
1. Install vendor dependencies
$ go mod vendor

2. Build the binary
$ go build

3. Export environment variables
Check the 'application.yml.sample' file and export the value of envs
for example:
$ export APP_PORT=9000
$ export DB_HOST=localhost

4. Create DB migrations
Create a database with the following name: 'asset_management'
Check migrations folder for creating db tables/indexes

5. Run the shell script for inserting data in table:
$ chmod +x seed.sh
$ ./seed.sh postgres 1234 asset_management

6. Run the server
$ ./asset-management
```

# How to build a docker image and run it
```

1. Building a docker image
$ docker build . -t <image-name>

2. Run the image
$ docker run -e APP_PORT=9000 -e DB_PORT=5432 -e EVENT_SERVICE_URL=http://localhost:9035 -e DB_HOST=localhost -e DB_USERNAME=postgres -e DB_PASSWORD=12345 -e DB_NAME=asset_management -e EVENT_API_TIMEOUT=3 -e JWT_SECRET=secret -e TOKEN_EXPIRY=5 -d --network=host -p 9000:9000 <image-name>

3. To check if image is running
$ docker ps -a

4. To read the logs
$ docker logs <container-id>

5. To stop the image 
$ docker stop <container-name>

```