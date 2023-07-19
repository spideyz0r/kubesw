package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getVersion = &cobra.Command{
		Use:   "version",
		Short: "Get version of kubesw",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.0.10")
		},
	}
)

func init() {
	getCmd.AddCommand()
}
