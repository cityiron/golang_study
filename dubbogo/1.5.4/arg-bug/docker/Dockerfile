#build env
FROM golang:1.14-alpine
RUN sed -i -e 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && \
    apk --update --no-cache add gcc musl-dev
LABEL maintainer="Tie Cheng, ironcity@apache.com"

RUN mkdir /build

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

ADD arg-bug/server /build/
ADD go.mod /build/
ADD go.sum /build/

RUN echo $(ls -l /build)

WORKDIR /build
RUN go build -o dg-server service.go

FROM ironc/alpine-bash:0.0.1
LABEL maintainer="Tie Cheng, ironcity@apache.com"

RUN sed -i -e 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && \
    apk --update --no-cache add git tzdata ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    unset GOPATH

RUN mkdir -p /app/

COPY --from=builder /build/dg-server /app/
COPY --from=builder /build/config /app/config/

RUN echo $(ls -l /app)

WORKDIR /app

ENV CONF_PROVIDER_FILE_PATH = /app/config/server.yml
ENV APP_LOG_CONF_FILE = /app/config/log.yml

CMD ["./dg-server"]