# 使用适当的基础镜像
FROM golang:1.19.1 as builder

# 创建一个工作目录
WORKDIR /app

# 创建一个工作目录
RUN go mod init myapp

# 复制源代码到容器中
COPY main.go ./
COPY pages ./pages

# 构建应用程序
RUN go build -o myapp

# 设置环境变量
ENV PORT=8080

# 指定容器启动时要执行的命令
CMD ["./myapp"]



