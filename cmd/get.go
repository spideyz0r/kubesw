package cmd

import (
	"fmt"
	common "kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get current context or namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Manage namespaces")
		},
	}
	namespaceGetCmd = &cobra.Command{
		Use:   "namespace",
		Short: "get namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", common.GetCurrent("namespace"))

		},
	}
	contextGetCmd = &cobra.Command{
		Use:   "context",
		Short: "get a context",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", common.GetCurrent("context"))
		},
	}
)

func init() {
	getCmd.AddCommand(namespaceGetCmd)
	getCmd.AddCommand(contextGetCmd)
}