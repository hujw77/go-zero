// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/dao"
	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/service"
	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/server/grpc"
	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
