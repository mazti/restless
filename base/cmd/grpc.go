package cmd

import (
	"context"
	"database/sql"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/pb"
	"github.com/mazti/restless/base/repository"
	"github.com/mazti/restless/base/service"
	"github.com/mazti/restless/base/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RunGRPC(listener net.Listener) error {
	service.Init(*Config.IDSalt)

	db, err := sql.Open("mysql", Config.Database.URL)
	CheckError(err)
	db.SetMaxOpenConns(Config.Database.MaxActive)
	db.SetMaxIdleConns(Config.Database.MaxIdle)
	baseRepository := repository.NewBaseRepository(db)


	client, err := ent.Open("mysql", Config.MetaDatabase.URL)
	CheckError(err)
	defer client.Close()
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	metaRepo := repository.NewMetaRepository(client)

	baseService, err := service.NewBaseService(baseRepository, metaRepo)
	tableService, err := service.NewTableService(baseRepository, metaRepo)
	CheckError(err)

	s := grpc.NewServer()
	pb.RegisterBaseServer(s, transport.NewBaseGRPCServer(baseService))
	pb.RegisterTableServer(s, transport.NewTableGRPCServer(tableService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
