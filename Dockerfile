# 使用官方的Go镜像作为基础镜像
FROM golang:1.23.1-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录为/workspace/myblog
WORKDIR /myblog

# 将本地的源代码复制到容器中
COPY . .

# 下载依赖包
RUN go mod download

# 编译Go程序，生成可执行文件
RUN go build -o main .

FROM debian:bullseye-slim

COPY ./wait-for.sh /

COPY ./configs /configs
COPY ./view /view
COPY ./uploads /uploads

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /myblog/main /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
	--no-install-recommends \
	netcat; \
    chmod 755 wait-for.sh