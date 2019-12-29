package servers

import (
	"github.com/go-kit/kit/log"
	"github.com/tiennv147/restless/meta/config"
	"os"
)

var (
	Config *config.Config
	Logger log.Logger
)

func init() {
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
