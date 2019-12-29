package main

import (
	"github.com/oklog/oklog/pkg/group"
	"github.com/tiennv147/restless/meta/cmd/servers"
	"net"
)

func main() {
	g := &group.Group{}

	grpcListener, err := net.Listen("tcp", servers.Config.GRPC.ListenAddr)
	servers.CheckError(err)
	g.Add(func() error {
		return servers.RunGRPC(grpcListener)
	}, func(error) {
		servers.CheckError(grpcListener.Close())
	})

	httpListener, err := net.Listen("tcp", servers.Config.HTTP.ListenAddr)
	servers.CheckError(err)
	g.Add(func() error {
		return servers.RunHTTP(httpListener)
	}, func(error) {
		servers.CheckError(httpListener.Close())
	})

	servers.CheckError(g.Run())
}
