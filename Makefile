SHELL := /bin/bash

PROJECT_SRC=${GOPATH}/src/github.com/yuwenhui/gin-starter
GO_SRC=${GOPATH}/src

pbgen:
	protoc -I ${PROJECT_SRC} ${PROJECT_SRC}/proto/protobuf/common.proto --go_out=${GO_SRC}
	./inline_replace.py proto/common.pb.go
	protoc -I ${PROJECT_SRC} ${PROJECT_SRC}/proto/protobuf/task.proto --go_out=${GO_SRC}
	./inline_replace.py proto/task.pb.go