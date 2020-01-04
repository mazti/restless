package main

import (
	base "github.com/mazti/restless/base/cmd"
	meta "github.com/mazti/restless/meta/cmd"
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
	RestlessCmd.AddCommand(meta.MetaCmd)
	RestlessCmd.Execute()
}
