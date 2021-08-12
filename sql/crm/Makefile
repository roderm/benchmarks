.PHONY: proto
proto:
	protoc \
		--proto_path=${GOPATH}/src/:${GOPATH}/src/github.com/gogo/protobuf/:${GOPATH}/src/github.com/gogo/protobuf/protobuf:. \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp:. \
		--go-sqlmap_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp:. \
		crm.proto