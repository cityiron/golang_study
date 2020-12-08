package main

import (
	"context"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/remoting/getty"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

var demoProvider = new(DemoProvider)

// 如果不需要配置文件，可以用下面的方式
func init() {
	// 需要在解析配置之前注册
	config.SetConsumerService(demoProvider)
	hessian.RegisterPOJO(&User{})

	// dubbogo comsumer config
	dgCfg := config.ConsumerConfig{
		Check:      new(bool),
		Registries: make(map[string]*config.RegistryConfig, 4),
		References: make(map[string]*config.ReferenceConfig, 4),
	}
	dgCfg.ApplicationConfig = &config.ApplicationConfig{
		Organization: "dubbo-go-proxy",
		Name:         "Dubbogo Proxy",
		Module:       "dubbogo proxy",
		Version:      "0.1.0",
		Owner:        "Dubbogo Proxy",
		Environment:  "dev",
	}
	dgCfg.Registries["demoZk"] = &config.RegistryConfig{
		Protocol:   "zookeeper",
		Address:    "127.0.0.1:2181",
		TimeoutStr: "3s",
	}
	dgCfg.References["DemoProvider"] = &config.ReferenceConfig{
		Registry:      "demoZk",
		Protocol:      "dubbo",
		InterfaceName: "com.funnycode.DemoService",
		Cluster:       "failover",
		Version:       "1.0.0",
		Group:         "tc",
	}
	config.SetConsumerConfig(dgCfg)
	getty.SetClientConf(getty.GetDefaultClientConfig())
	config.Load()
}

type User struct {
	Name string
	Age  int
}

func (User) JavaClassName() string {
	return "com.funnycode.User"
}

type DemoProvider struct {
	SayHello  func(ctx context.Context, name string) (string, error)             `dubbo:"sayHello"`
	SayHello2 func(ctx context.Context, user *User) (string, error)              `dubbo:"sayHello"`
	SayHello3 func(ctx context.Context, user *User, name string) (string, error) `dubbo:"sayHello"`
	SayHello4 func(ctx context.Context, user *User) (string, error)
	//SayHello4 func(ctx context.Context, user User) (string, error)
	SayHello5 func(ctx context.Context, user *User, name string) (string, error)
}

func (p *DemoProvider) Reference() string {
	return "DemoProvider"
}

func (p *DemoProvider) MethodMapper() map[string]string {
	return map[string]string{
		"SayHello": "SayHello",
	}
}
