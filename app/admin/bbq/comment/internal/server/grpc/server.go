package grpc

import (
	pb "github.com/HuJingwei/go-zero/app/admin/bbq/comment/api"
	"github.com/HuJingwei/go-zero/app/admin/bbq/comment/internal/service"
	"github.com/HuJingwei/go-zero/library/net/rpc/warden"
	"github.com/HuJingwei/go-zero/library/conf/paladin"
)

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
