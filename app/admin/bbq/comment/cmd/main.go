package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.atmatrix.org/k12/zero/app/admin/bbq/comment/internal/server/grpc"
	"git.atmatrix.org/k12/zero/app/admin/bbq/comment/internal/server/http"
	"git.atmatrix.org/k12/zero/app/admin/bbq/comment/internal/service"
	"git.atmatrix.org/k12/zero/library/conf/paladin"
	ecode "git.atmatrix.org/k12/zero/library/ecode/tip"
	"git.atmatrix.org/k12/zero/library/log"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("comment-admin start")
	ecode.Init(nil)
	svc := service.New()
	grpcSrv := grpc.New(svc)
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, _ := context.WithTimeout(context.Background(), 35*time.Second)
			grpcSrv.Shutdown(ctx)
			httpSrv.Shutdown(ctx)
			log.Info("comment-admin exit")
			svc.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
