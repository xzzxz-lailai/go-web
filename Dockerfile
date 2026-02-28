# 编译阶段
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 先复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 再复制全部源码
COPY . .

# ENV CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct
# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags "-s -w" -o app .
    # 默认是开启（1），写 CGO_ENABLED=0 就是关闭关闭 C 语言依赖，强制纯 Go + 静态二进制
    # -trimpath 去掉源码路径信息减小 + 安全
    # -ldflags "-s -w"  去掉符号表 + 调试信息减小
    #  -o app .          把当前目录(.)的 main package 编译成可执行文件,输出名叫 app



# 运行阶段
FROM alpine:3.21
WORKDIR /app

# 设置时区为上海
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# 从 builder 阶段的 /app/app 文件，只复制这个二进制文件到当前阶段的当前目录（也就是 /app）
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
