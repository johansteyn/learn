MongoDB in Docker
=================

Install:
  % docker pull mongodb/mongodb-community-server

Run the container:
  % docker run --name mongodb -d mongodb/mongodb-community-server

To specify a port:
  % docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server
  
To persist data outside of the container:
  % docker run --name mongodb -d -p 27017:27017 -v $(pwd)/data:/data/db mongodb/mongodb-community-server

To automatically create a user with root permissions, specify a username and password using environment variables:
  % docker run --name mongodb -d -p 27017:27017 -v $(pwd)/data:/data/db -e MONGO_INITDB_ROOT_USERNAME=johan -e MONGO_INITDB_ROOT_PASSWORD=password mongodb/mongodb-community-server

Stop and remove the container:
  % docker stop mongodb
  % docker rm mongodb

MongoDB Golang Driver
=====================

With MongoDB up and running in Docker, looked at Golang libraries for MongoDB...
    https://www.mongodb.com/docs/drivers/go/current
    https://github.com/mongodb/mongo-go-driver
    https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
    https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo/options

Tried the v2 driver first, but ran into issue with go mod tidy:
  go.mongodb.org/mongo-driver/v2/options: module go.mongodb.org/mongo-driver/v2@latest found (v2.0.0), but does not contain package go.mongodb.org/mongo-driver/v2/options
Running these did not help:
  % go get go.mongodb.org/mongo-driver/v2/mongo
  % go get go.mongodb.org/mongo-driver/v2/mongo/options
So, I'll stick with the v1 driver for now...

