package main

import (
	"github.com/spf13/cobra"
	meta "github.com/tiennv147/restless/meta/cmd"
	base "github.com/tiennv147/restless/base/cmd"
)

var (
	RestlessCmd = &cobra.Command{
		Use:   "restless",
		Short: "RESTLESS",
		Long:  `Great RESTLESS`,
	}
)

func main() {
	RestlessCmd.AddCommand(base.BaseCmd)
	RestlessCmd.AddCommand(meta.MetaCmd)
	RestlessCmd.Execute()
}
