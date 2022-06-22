FROM golang:1.17.1-alpine3.14 as builder

# 指定构建过程中的工作目录
WORKDIR /goapi

# 将当前目录（dockerfile所在目录）下所有文件都拷贝到工作目录下
COPY . /goapi/

# 执行代码编译命令。操作系统参数为linux，编译后的二进制产物命名为main，并存放在当前目录下。
RUN GOOS=linux GOARCH=amd64 go build -o main .

# 选用运行时所用基础镜像（GO语言选择原则：尽量体积小、包含基础linux内容的基础镜像）
FROM alpine:3.13

# 指定运行时的工作目录
WORKDIR /goapi

# 将构建产物/goapi/main拷贝到运行时的工作目录中
COPY --from=builder /goapi/main /goapi/
COPY --from=builder /goapi/comm/config/server.conf /goapi/comm/config/

# 设置时区
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# 兼容开放接口服务
RUN apk add ca-certificates

# 设置release模式
ENV GIN_MODE release

# 执行启动命令
CMD ["/goapi/main"]
