#源镜像
FROM golang:latest
#设置工作目录
WORKDIR $GOPATH/src/github.com/mygohttp
#将服务器的go工程代码加入到docker容器中
ADD https://codeload.github.com/golang/crypto/zip/master  $GOPATH/src/golang.org/x/crypto
ADD https://codeload.github.com/golang/net/zip/master  $GOPATH/src/golang.org/x/net
RUN sudo apt-get update &&\
    sudo apt-get install -y unzip  &&\
    unzip $GOPATH/src/golang.org/x/crypto &&\
    mv $GOPATH/src/golang.org/x/crypto-master $GOPATH/src/golang.org/x/crypto &&\
    ls $GOPATH/src/golang.org/x -trl &&\
    go get github.com/labstack/echo/...
ADD  . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./mygohttp"]



