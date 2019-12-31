package cmd

import (
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"github.com/tiennv147/restless/meta/config"
	"net"
	"os"
)

var (
	Config *config.Config
	Logger log.Logger

	MetaCmd = &cobra.Command{
		Use:   "meta",
		Short: "RESTLESS - META",
		Long:  `The Great RESTLESS - META`,
		Run: func(cmd *cobra.Command, args []string) {
			grpcOnly, _ := cmd.Flags().GetBool("grpc")
			g := &group.Group{}

			grpcListener, err := net.Listen("tcp", Config.GRPC.ListenAddr)
			CheckError(err)
			g.Add(func() error {
				Logger.Log("RESTLESS - META", "GRPC server", "started...", Config.GRPC.ListenAddr)
				return RunGRPC(grpcListener)
			}, func(error) {
				Logger.Log("RESTLESS - META", "GRPC server", "stopped...", Config.GRPC.ListenAddr)
				CheckError(grpcListener.Close())
			})

			if !grpcOnly {
				httpListener, err := net.Listen("tcp", Config.HTTP.ListenAddr)
				CheckError(err)
				g.Add(func() error {
					Logger.Log("RESTLESS - META", "HTTP server", "started...", Config.HTTP.ListenAddr)
					return RunHTTP(httpListener)
				}, func(error) {
					Logger.Log("RESTLESS - META", "HTTP server", "stopped...", Config.HTTP.ListenAddr)
					CheckError(httpListener.Close())
				})
			}
			CheckError(g.Run())
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return MetaCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	MetaCmd.Flags().BoolP("grpc", "g", false, "Only run GRPC server")
}

func initConfig() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)

	var err error
	Config, err = config.New()
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		Logger.Log("ERROR", err.Error())
		os.Exit(-1)
	}
}
