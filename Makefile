# 伪目标，如果未指定终极目标，将默认使用Makefile里的第一个目标，
# 即Makefile里最靠前的规则（本Makefile是all这个目标）
# 因此，当执行make时，会默认第一个目标为终极目标，所以all后面的依赖都会执行下去
# 因此，可以将all后面放置一些必须编译的默认选项，在只使用make时；当然也可以使用make指定终极目标进行构建
.PHONY: openssl format build build_linux build_win clean swag docker-build help format test run
all: openssl format build

# 声明编译项目的文件名
BUILD_NAME=web_app

# swagger接口文档初始化
swag:
	@swag init

# 在./config目录，签发自建的tls证书
# 或者使用go标准库：go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
openssl:
	@openssl genrsa -out ./config/key.pem 2048;openssl req -new -x509 -key ./config/key.pem -out ./config/cert.pem -days 3650

# 使用Dockerfile对项目打包编译出镜像
docker-build: swag
	@docker build -t ${BUILD_NAME}:1.0

# 格式化项目
format:
	@go fmt ./
	@go vet ./

# 测试代码
test: swag
	@go test -v #回归测试

# 直接运行项目根目录下已经编译好的二进制文件
run:
	./${BUILD_NAME}

# 默认编译
build: test
	@go build -o ${BUILD_NAME} ${SOURCE}

# 交叉编译--适应linux系统
build_linux: test
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_NAME} .

# 交叉编译--适应windows系统
build_win: test
    @CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BUILD_NAME} .

# 清除编译文件
clean:
	@go clean

# 帮助命令
help:
	@echo "make - 格式化 Go 代码、更新swagger文档、生成tls证书、测试代码、编译生成二进制文件"
	@echo "make docker-build - 构建本项目的Docker镜像"
	@echo "make build - 编译 Go 代码, 生成当前环境默认的二进制文件"
	@echo "make build_linux - 编译 Go 代码, 生成linux环境二进制文件"
	@echo "make build_win - 编译 Go 代码, 生成windows环境二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make format - 运行 Go 工具 'fmt' and 'vet'"