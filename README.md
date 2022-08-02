## Message Delivery System

This is a Go implementation for a client-server message system using bidirectional streaming RPCs

## Run Testing

#### Start the server
`cd` to server directory, run `go run ./server.go`

Example output
```
2022/08/01 21:15:21 Starting Server On Port::8080
```

#### Send identity message
`cd` to client directory, run `go run ./client.go identity`

Example output
```
2022/08/01 21:17:21 Message Type: identity, Content: 
2022/08/01 21:17:21 received: Users:1659403041  Message:"register!"
```

#### Send list message
`cd` to client directory, run `go run ./client.go list`

Example output
```
2022/08/01 21:18:01 Message Type: list, Content: 
2022/08/01 21:18:01 received: Users:1659403041
```

#### Send relay message
`cd` to client directory, run `go run ./client.go relay 1659403041 test-message`

Client id: 1659403041, receives following:
```
2022/08/01 21:19:12 received: Message:"test-message"
```

Send relay message to multiple client, use comma saperate user id, e.g. 

`go run ./client.go relay 1659403041,1659403042 test-message`
