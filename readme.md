# Simple calculator

Programming assignement consisting in building a simple calculator in Golang. 
It consists of a client and a server that communicate via gRPC.

## Installation

Clone this repository 

```bash
https://github.com/rmcaxoul/grpc_server.git
```

Make sure you have all the dependencies
```bash
cd grpc_server
go get
```

Start the server
```bash
go run server.go
```
Then you can run the client
```bash
go run client/client.go -a=[first_number] -b=[second_number] -method=[operator]
```
## Constraints

This calculator works only for integers (positive and negative).
Inputing anything else into a or b will result in an error.
Method has to be + - * or /.
All fields are mandatory.
Division will result in a floating point.
Works with int32:
```go
max int32   = 2147483647
min int32   = -2147483648
```
If the limit is reached within an operation, an error will be returned.