package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/di"
	"github.com/HuJingwei/go-zero/pkg/conf/env"
	"github.com/HuJingwei/go-zero/pkg/conf/paladin"
	"github.com/HuJingwei/go-zero/pkg/log"
	"github.com/HuJingwei/go-zero/pkg/naming"
	"github.com/HuJingwei/go-zero/pkg/naming/discovery"
	xip "github.com/HuJingwei/go-zero/pkg/net/ip"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("admin.main.app start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	// start discovery register
	var cancel context.CancelFunc
	ip := xip.InternalIP()
	hn, _ := os.Hostname()
	dis := discovery.New(nil)
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    "admin.main.app",
		Hostname: hn,
		Addrs: []string{
			"http://" + ip + ":" + "8000",
			"grpc://" + ip + ":" + "9000",
		},
	}
	if cancel, err = dis.Register(context.Background(), ins); err != nil {
		panic(err)
	}
	// end discovery register

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}
			closeFunc()
			log.Info("admin.main.app exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
