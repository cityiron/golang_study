## 如何运行

### 安装 Zookeeper 服务

运行之前需要有一个 [zookeeper](https://zookeeper.apache.org/releases.html) 服务环境，你可以从官方下载后用命令启动：

```bash
zkServer start
```

或者当你有 docker 环境的话可以用 docker 启动:

```bash
docker run --name zookeeper -p2181:2181 -d zookeeper
```

### 启动 dubbogo server

```bash
docker pull ironc/gd-server:dgissue900
docker run --network host --name go-provider -it ironc/gd-server:dgissue900 
```

### 启动 dubbogo client

```bash
docker pull ironc/gd-client:dgissue900
docker run --network host --name go-consumer -it ironc/gd-client:dgissue900
```

输出：

```bash
[2021-01-14/13:29:19 main.main: client.go: 25] 


start to test dubbo
[2021-01-14/13:29:19 main.main: client.go: 32] response result: Hello tc

[2021-01-14/13:29:19 main.main: client.go: 58] response result: Hello tc You are 18

[2021-01-14/13:29:19 main.main: client.go: 65] response result: Hello tc You are 18
```

### 启动 dubbogo client 调用 java 复现错误

镜像方式：

```bash
docker pull ironc/gd-client-2-java:dgissue900
docker run --network host --name dubbogo-client-2-java -it ironc/gd-client-2-java:dgissue900
```

如果镜像方式无法启动，可以直接运行当前的 client/client.go 代码，配置 CONF_CONSUMER_FILE_PATH 变量

比如：

```bash
export CONF_CONSUMER_FILE_PATH=/XXX/golang_study/dubbogo/1.5.4/arg-bug/client/config/client.yml
go run client.go demo.go
```