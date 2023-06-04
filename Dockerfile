FROM golang:1.19.2 as builder

ENV MYPATH /usr/local
WORKDIR $MYPATH


ADD main.go ./
ADD ./pages ./
ADD ./src ./
ADD index.html ./
RUN go mod init ./main
RUN go build -o ./main ./main.go

EXPOSE 8080

CMD ./main
