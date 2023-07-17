package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getnamespaceCmd = &cobra.Command{
		Use:   "get",
		Short: "Manage namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Manage namespaces")
		},
	}
	namespaceGetCmd = &cobra.Command{
		Use:   "namespace",
		Short: "get namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting the current namespace")
		},
	}
	contextGetCmd = &cobra.Command{
		Use:   "context",
		Short: "get a context",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting the current context")
		},
	}
)

func init() {
	getnamespaceCmd.AddCommand(namespaceGetCmd)
	getnamespaceCmd.AddCommand(contextGetCmd)
}