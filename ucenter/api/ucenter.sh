goctl rpc protoc register.proto --go_out=./types --go-grpc_out=./types --zrpc_out=./register --style go_zero
protoc login.proto --go_out=./types --go-grpc_out=./types

goctl model mysql datasource --url="root:root@tcp(127.0.0.1:13306)/mscoin" --table="member" -c --dir .