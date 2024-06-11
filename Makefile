
# Makefile

# 指定生成代码的目标路径
GO_GEN_PATH=./api

# 你的 .proto 文件路径
PROTO_PATH=./api

# proto 文件名
PROTO_FILE=binance.proto

# 指定 protoc 命令所在的路径，如果已经在 PATH 中则不必修改
PROTOC=protoc

# 确保 protoc-gen-go 和 protoc-gen-go-grpc 插件安装在 $PATH 中
# 您可以通过如下命令安装它们：
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: all gen clean

all: gen

# 编译 .proto 文件生成 Go 源代码
gen:
	$(PROTOC) -I$(PROTO_PATH) --go_out=$(GO_GEN_PATH) --go_opt=paths=source_relative \
	--go-grpc_out=$(GO_GEN_PATH) --go-grpc_opt=paths=source_relative \
	$(PROTO_PATH)/$(PROTO_FILE)

# 清除生成的文件
clean:
	rm -f $(GO_GEN_PATH)/*.go

# install golang 1.21 on ubuntu 20.04
go-init:
	wget https://go.dev/dl/go1.20.14.linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf go1.20.14.linux-amd64.tar.gz
	echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
	source ~/.bashrc
	go version