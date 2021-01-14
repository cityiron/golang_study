FROM ironc/alpine-bash:0.0.1

RUN mkdir /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOARCH = amd64

ADD go.mod /app/
ADD go.sum /app/
ADD arg-bug/client/config /app/config/
ADD arg-bug/client/dg-client /app/

WORKDIR /app

ENV CONF_CONSUMER_FILE_PATH /app/config/client.yml
ENV APP_LOG_CONF_FILE /app/config/log.yml

CMD ["./dg-client"]