# GRPCDump
Tool for capture and parse GRPC messages from ethernet traffic

## Installation / Getting started

```bash
go run cmd/grpcdump.go
```
### From Source

```bash
go get github.com/rmedvedev/grpcdump
go install github.com/rmedvedev/grpcdump
```

## Usage 

The next option explains usage doc
```bash
grpcdump -help
```
For example, to capture grpc traffic and parse grpc request and response you need to use this command:
```bash
grpcdump -i lo -p 50051 -proto-path ./grpc/protofiles -proto-files helloworld.proto 
```