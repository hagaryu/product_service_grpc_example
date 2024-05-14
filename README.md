# product_service_grpc_example
created a very simple example for grpc server and client to show how the basics can work

# generating proto piles for golang:
from root directory of the project run:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative */*/*.proto
```

# How to run:
run/debug the example_server/main.go first. 
Once it writes: "server listening at 50051" run/debug the example_client/main.go
