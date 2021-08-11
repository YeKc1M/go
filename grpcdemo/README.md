generate .pb.go file

`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/<protofile>.proto`

generate .pb.go & _grpc.pb.go

`protoc -I $SRC_DIR --go_out $DST_DIR --go_opt paths=source_relative --go-grpc_out $DST_DIR --go-grpc_opt paths=source_relative $SRC_DIR/<protofile>`

`protoc -I ./proto --go_out ./pb/template_engine --go_opt paths=source_relative --go-grpc_out ./pb/template_engine --go-grpc_opt paths=source_relative ./proto/test.proto`

generate .pb.gw.go

`protoc -I ./proto --grpc-gateway_out ./pb/template_engine --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative ./proto/test.proto`

# reference

[https://www.cnblogs.com/linguoguo/p/10148467.html](https://www.cnblogs.com/linguoguo/p/10148467.html)