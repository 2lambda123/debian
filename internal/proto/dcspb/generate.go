package dcspb

//go:generate protoc -I.. -I. dcs.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
