package main

import (
	"context"
	"strconv"

	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/config"
)

var (
	userProvider = new(DemoProvider)
)

func init() {
	config.SetProviderService(userProvider)
	hessian.RegisterPOJO(&User{})
}

type User struct {
	Name string
	Age  int
}

func (User) JavaClassName() string {
	return "com.funnycode.User"
}

type DemoProvider struct{}

func (p *DemoProvider) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello " + name, nil
}

func (p *DemoProvider) SayHello4(ctx context.Context, user *User) (string, error) {
	return "Hello " + user.Name + " You are " + strconv.Itoa(user.Age), nil
}

func (p *DemoProvider) SayHello5(ctx context.Context, user *User, name string) (string, error) {
	return "Hello " + name + " You are " + strconv.Itoa(user.Age), nil
}

func (p *DemoProvider) Reference() string {
	return "DemoProvider"
}

func (p *DemoProvider) MethodMapper() map[string]string {
	return map[string]string{
		"SayHello": "sayHello",
	}
}
