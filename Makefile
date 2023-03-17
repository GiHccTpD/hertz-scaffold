.PHONY: build run compile swag clean help

BINARY_NAME=ServerName

build:
	go build -o ${BINARY_NAME} main.go

run:
	./&=${BINARY_NAME}

compile:
	echo "Compiling for every OS and Platform"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o deploy/${BINARY_NAME}_darwin main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/${BINARY_NAME}_linux main.go
	 CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o deploy/${BINARY_NAME}_windows main.go

swag:
	swag init --parseDependency --parseInternal --parseDepth 5 --instanceName "swagger"

clean:
	go clean
	rm -rf deploy/*

help:
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make compile - 编译 Go 代码, 生成不同系统可执行的二进制文件"
	@echo "make swag - 生成swagger配置文件"
	@echo "make clean - 移除二进制文件和 vim swap files"
