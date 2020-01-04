package main

import (
	base "github.com/mazti/restless/base/cmd"
	meta "github.com/mazti/restless/meta/cmd"
	record "github.com/mazti/restless/record/cmd"
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
	RestlessCmd.AddCommand(record.RecordCmd)
	RestlessCmd.Execute()
}
