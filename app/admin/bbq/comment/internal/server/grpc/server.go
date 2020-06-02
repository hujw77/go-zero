package grpc

import (
	pb "git.atmatrix.org/k12/zero/app/admin/bbq/comment/api"
	"git.atmatrix.org/k12/zero/app/admin/bbq/comment/internal/service"
	"git.atmatrix.org/k12/zero/library/net/rpc/warden"
	"git.atmatrix.org/k12/zero/library/conf/paladin"
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
