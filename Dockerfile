#源镜像
FROM golang:latest
#设置工作目录
WORKDIR $GOPATH/src/github.com/mygohttp
#将服务器的go工程代码加入到docker容器中
COPY /usr/local/go/src/golang.org/x  $GOPATH/src/golang.org/x
COPY /usr/local/go/src/github.com/labstack/echo  $GOPATH/src/github.com/labstack/echo
ADD . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./mygohttp"]
