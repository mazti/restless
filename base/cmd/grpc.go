package cmd

import (
	"database/sql"
	base "github.com/tiennv147/restless/base/pb/base"
	"github.com/tiennv147/restless/base/repository"
	"github.com/tiennv147/restless/base/service"
	"github.com/tiennv147/restless/base/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunGRPC(listener net.Listener) error {
	db, err := sql.Open("mysql", Config.Database.URL)
	CheckError(err)
	db.SetMaxOpenConns(Config.Database.MaxActive)
	db.SetMaxIdleConns(Config.Database.MaxIdle)

	baseRepository := repository.NewBaseRepository(db)
	baseService, err := service.NewBaseService(baseRepository, *Config)
	tableService, err := service.NewTableService(baseRepository, *Config)
	CheckError(err)

	s := grpc.NewServer()
	base.RegisterBaseServer(s, transport.NewBaseGRPCServer(baseService))
	base.RegisterTableServer(s, transport.NewTableGRPCServer(tableService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
