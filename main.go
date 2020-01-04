package main

import (
	base "github.com/mazti/restless/base/cmd"
	"github.com/spf13/cobra"
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
	RestlessCmd.Execute()
}
