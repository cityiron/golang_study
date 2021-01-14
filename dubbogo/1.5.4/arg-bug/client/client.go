package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	gxlog "github.com/dubbogo/gost/log"
)

var (
	survivalTimeout int = 10e9
)

// they are necessary:
// 		export CONF_CONSUMER_FILE_PATH="xxx"
// 		export APP_LOG_CONF_FILE="xxx"
func main() {
	config.Load()
	gxlog.CInfo("\n\n\nstart to test dubbo")

	res, err := demoProvider.SayHello(context.TODO(), "tc")
	if err != nil {
		panic(err)
	}

	gxlog.CInfo("response result: %v\n", res)

	user := &User{
		Name: "tc",
		Age:  18,
	}

	res, err = demoProvider.SayHello2(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	gxlog.CInfo("response result: %v\n", res)

	res, err = demoProvider.SayHello3(context.TODO(), user, "tc")
	if err != nil {
		panic(err)
	}

	gxlog.CInfo("response result: %v\n", res)

	//res, err = demoProvider.SayHello4(context.TODO(), user)
	//if err != nil {
	//	panic(err)
	//}
	//
	//gxlog.CInfo("response result: %v\n", res)
	//
	//res, err = demoProvider.SayHello5(context.TODO(), user, "tc")
	//if err != nil {
	//	panic(err)
	//}
	//
	//gxlog.CInfo("response result: %v\n", res)

	initSignal()
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP,
		syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("app exit now...")
			return
		}
	}
}
