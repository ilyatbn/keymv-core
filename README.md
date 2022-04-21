# core gRPC server

## Notes to self  

In order to update your request/response in the gRPC server, you need to run the following command to regenerate all the functions in the code from within the folder where the .proto file is located:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative params.proto 
```

Before you go get a private repo you need to set this shit:
export GOPRIVATE=github.com/ilyatbn/keymv-proto  
go get github.com/ilyatbn/keymv-proto/core



ToDo
- Implement SSL
-