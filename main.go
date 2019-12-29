package main

import (
	"github.com/spf13/cobra"
	meta "github.com/tiennv147/restless/meta/cmd"
)

var (
	RestlessCmd = &cobra.Command{
		Use:   "restless",
		Short: "RESTLESS",
		Long:  `Great RESTLESS`,
	}
)

func main() {
	RestlessCmd.AddCommand(meta.MetaCmd)
	RestlessCmd.Execute()
}
