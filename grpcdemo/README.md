generate .pb.go file

`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/<protofile>.proto`

generate .pb.go & _grpc.pb.go

`protoc -I $SRC_DIR --go_out $DST_DIR --go_opt paths=source_relative --go-grpc_out $DST_DIR --go-grpc_opt paths=source_relative $SRC_DIR/<protofile>`

`protoc -I ./proto --go_out ./pb/server --go_opt paths=source_relative --go-grpc_out ./pb/server --go-grpc_opt paths=source_relative ./proto/test.proto`

generate _grpc.pb.go

`protoc -I ./proto --grpc-gateway_out ./pb/server --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative ./proto/test.proto`
