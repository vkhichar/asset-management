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
# How to use Makefile 

1.to build
$ make build

2.to test 
$ make test

3.to formate your code use
$ make fmt

4.to migrate database migrations
$ make migrate

5.to rollback database migrations
$ make rollback

6.to start server
$ make startapp