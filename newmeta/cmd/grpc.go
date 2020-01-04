package cmd

import (
	"context"
	"github.com/mazti/restless/newmeta/ent"
	"github.com/mazti/restless/newmeta/pb"
	"github.com/mazti/restless/newmeta/repository"
	"github.com/mazti/restless/newmeta/service"
	"github.com/mazti/restless/newmeta/transport"
	"github.com/speps/go-hashids"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RunGRPC(listener net.Listener) error {
	client, err := ent.Open("mysql", Config.Database.URL)
	CheckError(err)
	defer client.Close()
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	hash, _ := hashids.NewWithData(hd)
	schemaRepo := repository.NewMetaRepository(client)
	schemaService := service.NewMetaService(schemaRepo, hash)

	s := grpc.NewServer()
	pb.RegisterMetaServer(s, transport.NewGRPCServer(schemaService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
