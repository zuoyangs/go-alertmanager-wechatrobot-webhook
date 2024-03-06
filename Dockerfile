# 第一阶段：构建阶段
FROM registry.cn-hangzhou.aliyuncs.com/zuoyang/golang:1.22.0-alpine3.19 AS builder

ENV PATH /usr/local/bin:$PATH
ENV LANG C.UTF-8  \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN set -ex \
    && echo "export GO111MODULE=on" >> /root/profile \
    && echo "export GOPROXY=https://goproxy.cn" >> /root/profile \
    && . /root/profile \
    && go build -ldflags "-w" -o /app/go-webchat .

# 第二阶段：最终阶段
FROM registry.cn-hangzhou.aliyuncs.com/zuoyang/alpine:3.19

ENV TZ=Asia/Shanghai
ENV PATH /usr/local/bin:$PATH

RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache bash tzdata curl wget procps net-tools \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone  \
    && echo "alias ll='ls -l --color=auto'" >> /root/profile \
    && . /root/profile

COPY --from=builder /app/go-webchat /opt/go-webchat

EXPOSE 8080

CMD ["/opt/go-webchat", "--RobotKey=577aa9e0-816f-49f4-80d7-40e0237c77c5"]
