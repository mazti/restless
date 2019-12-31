package cmd

import (
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"github.com/tiennv147/restless/base/config"
	"net"
	"os"
)

var (
	Config *config.Config
	Logger log.Logger

	BaseCmd = &cobra.Command{
		Use:   "base",
		Short: "RESTLESS - BASE",
		Long:  `The Great RESTLESS - BASE`,
		Run: func(cmd *cobra.Command, args []string) {
			grpcOnly, _ := cmd.Flags().GetBool("grpc")
			g := &group.Group{}

			Logger.Log(Config)
			grpcListener, err := net.Listen("tcp", Config.GRPC.ListenAddr)
			CheckError(err)
			g.Add(func() error {
				Logger.Log("RESTLESS - BASE", "GRPC server", "started...", Config.GRPC.ListenAddr)
				return RunGRPC(grpcListener)
			}, func(error) {
				Logger.Log("RESTLESS - BASE", "GRPC server", "stopped...", Config.GRPC.ListenAddr)
				CheckError(grpcListener.Close())
			})

			if !grpcOnly {
				httpListener, err := net.Listen("tcp", Config.HTTP.ListenAddr)
				CheckError(err)
				g.Add(func() error {
					Logger.Log("RESTLESS - BASE", "HTTP server", "started...", Config.HTTP.ListenAddr)
					return RunHTTP(httpListener)
				}, func(error) {
					Logger.Log("RESTLESS - BASE", "HTTP server", "stopped...", Config.HTTP.ListenAddr)
					CheckError(httpListener.Close())
				})
			}
			CheckError(g.Run())
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return BaseCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	BaseCmd.Flags().BoolP("grpc", "g", false, "Only run GRPC server")
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
