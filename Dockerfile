FROM golang:latest


ENV GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/SIMS
COPY . $GOPATH/src/SIMS

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
#RUN go mod init
#RUN go mod tidy
RUN go build .


EXPOSE 8080

# 运行golang程序的命令
ENTRYPOINT ["./SIMS"]