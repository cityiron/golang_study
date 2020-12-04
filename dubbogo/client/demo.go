package main

import (
	"context"
	hessian "github.com/apache/dubbo-go-hessian2"

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

func init() {
	config.SetConsumerService(demoProvider)
	hessian.RegisterPOJO(&User{})
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
