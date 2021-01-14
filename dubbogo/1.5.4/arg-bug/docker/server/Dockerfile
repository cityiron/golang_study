FROM ironc/alpine-bash:0.0.1

RUN mkdir /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOARCH = amd64

ADD go.mod /app/
ADD go.sum /app/
ADD arg-bug/server/config /app/config/
ADD arg-bug/server/dg-server /app/

WORKDIR /app

ENV CONF_PROVIDER_FILE_PATH /app/config/server.yml
ENV APP_LOG_CONF_FILE /app/config/log.yml

CMD ["./dg-server"]