package cmd

import (
	"fmt"
	"github.com/mazti/restless/newmeta/config"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"time"
)

var (
	Config, _ = config.New()
	MetaCmd   = &cobra.Command{
		Use:   "meta",
		Short: "RESTLESS - " + *Config.Name,
		Long:  "The Great RESTLESS - " + *Config.Name,
		Run: func(cmd *cobra.Command, args []string) {
			grpcOnly, _ := cmd.Flags().GetBool("grpc")
			g := &group.Group{}

			grpcListener, err := net.Listen("tcp", Config.GRPC.ListenAddr)
			CheckError(err)
			g.Add(func() error {
				log.Println("GRPC server", "started...", Config.GRPC.ListenAddr)
				return RunGRPC(grpcListener)
			}, func(error) {
				log.Println("GRPC server", "stopped...", Config.GRPC.ListenAddr)
				CheckError(grpcListener.Close())
			})

			if !grpcOnly {
				httpListener, err := net.Listen("tcp", Config.HTTP.ListenAddr)
				CheckError(err)
				g.Add(func() error {
					log.Println("HTTP server", "started...", Config.HTTP.ListenAddr)
					return RunHTTP(httpListener)
				}, func(error) {
					log.Println("HTTP server", "stopped...", Config.HTTP.ListenAddr)
					CheckError(httpListener.Close())
				})
			}
			CheckError(g.Run())
		},
	}
)

// Execute executes the root command.
func Execute() error {
	MetaCmd = &cobra.Command{
		Use:   "meta",
		Short: "RESTLESS - " + *Config.Name,
		Long:  "The Great RESTLESS - " + *Config.Name,
		Run: func(cmd *cobra.Command, args []string) {
			grpcOnly, _ := cmd.Flags().GetBool("grpc")
			g := &group.Group{}

			grpcListener, err := net.Listen("tcp", Config.GRPC.ListenAddr)
			CheckError(err)
			g.Add(func() error {
				log.Println("GRPC server", "started...", Config.GRPC.ListenAddr)
				return RunGRPC(grpcListener)
			}, func(error) {
				log.Println("GRPC server", "stopped...", Config.GRPC.ListenAddr)
				CheckError(grpcListener.Close())
			})

			if !grpcOnly {
				httpListener, err := net.Listen("tcp", Config.HTTP.ListenAddr)
				CheckError(err)
				g.Add(func() error {
					log.Println("HTTP server", "started...", Config.HTTP.ListenAddr)
					return RunHTTP(httpListener)
				}, func(error) {
					log.Println("HTTP server", "stopped...", Config.HTTP.ListenAddr)
					CheckError(httpListener.Close())
				})
			}
			CheckError(g.Run())
		},
	}
	return MetaCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	MetaCmd.Flags().BoolP("grpc", "g", false, "Only run GRPC server")
}

func initConfig() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{name: *Config.Name})
}

type logWriter struct {
	name string
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05-0700") + " [" + writer.name + "] " + string(bytes))
}

func CheckError(err error) {
	if err != nil {
		log.Println("ERROR", err.Error())
		os.Exit(-1)
	}
}
