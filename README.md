# Mago
User Service, responsible for creating user get user/profile information. Server has default TLS encryption by default. Sample TLS credentials, private/public keys, are available in `helpers/ssl` folder.

## Start Mago Server
To run gRPC server for user management you can use:
```shell script
go run mago_server/server.go
```

Server will start on gRPC default port, 50051, to give a different port number provide `--port` argument like below:
```shell script
go run mago_server/server.go --port 8080
```

On successful server startup you would see:
```shell script
$ go run mago_server/server.go --port 50052
Starting up Mago server on 0.0.0.0:50052...
Mago server is up & running ;)
```
