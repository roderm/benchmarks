GOGO_PATH=${GOPATH}/src/github.com/gogo/protobuf
.PHONY: gogo
gogo: 
	# check how to do it with GO111MODULE=on
	IMPORT_PATH=
	$(eval IMPORT_PATH=${IMPORT_PATH}:${GOPATH}/src/)
	$(eval IMPORT_PATH=${IMPORT_PATH}:${GOGO_PATH}/)
	$(eval IMPORT_PATH=${IMPORT_PATH}:${GOGO_PATH}/protobuf/)
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc;
	go install github.com/gogo/protobuf/protoc-gen-gogo;
	find ./ -type f -name *.proto -exec \
		protoc \
			--proto_path=${IMPORT_PATH}:. \
			--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
		{} \;

.PHONY: proto
proto:
	protoc \
		--proto_path=${GOPATH}/src/:${GOPATH}/src/github.com/gogo/protobuf/:${GOPATH}/src/github.com/gogo/protobuf/protobuf:. \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp:. \
		--go-sqlmap_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp:. \
		company.proto