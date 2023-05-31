FROM golang:1.19.1 as builder

WORKDIR /app

# 复制源代码和相关文件
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY src/ .

# 复制PAGES文件夹到容器中的/pages目录
COPY pages/ .

# 构建应用
RUN go build -o myapp ./src

# 设置容器启动命令
CMD ["./myapp"]
