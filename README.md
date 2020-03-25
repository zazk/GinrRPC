# Go Gin gRPC Tutorial

### Basic API using gRPC and protobufs in Go with Go Gin

### Server

```
cd server/ && go run main.go
```

### Client

```
cd client/ && go run main.go
```

### You can test both `GET` requests

### `/mult/:a/:b`

```
http://localhost:8080/mult/210/204
```

Response:

```
{
    "result": "42840"
}
```

### `/add/:a/:b`

```
http://localhost:8080/add/210/204

```

Response:

```
{
    "result": "42840"
}
```

### Tested with standalone client

![Image of Testing](http://g.recordit.co/svCipAEVn0.gif)

Client: https://github.com/uw-labs/bloomrpc
