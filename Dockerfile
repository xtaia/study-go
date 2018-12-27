#源镜像
FROM golang:latest
#设置工作目录
#WORKDIR $GOPATH/src/github.com/mygohttp
#将服务器的go工程代码加入到docker容器中
RUN mkdir lib && \
    cd lib && \
    git clone https://github.com/golang/crypto.git && \
    git clone https://github.com/golang/net.git && \ 
    cd ..
RUN ls lib && \
    pwd
RUN pwd   
COPY /go/lib/crypto  $GOPATH/src/golang.org/x/crypto
COPY /go/lib/net  $GOPATH/src/golang.org/x/net
RUN  ls $GOPATH/src/golang.org/x 
ADD  . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./mygohttp"]



