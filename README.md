# What's this ?
gRPC sample

# Pre Install
## Using Go
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## Using Java
Please ref https://github.com/grpc/grpc-javac

# Genarating Server/Client Stub
```
protoc --go_out=plugins=grpc:. echo.proto 
```

# Generating docs
```
protoc --doc_out=./doc --doc_opt=html,index.html echo/*.proto  
```
