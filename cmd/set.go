package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	setnamespaceCmd = &cobra.Command{
		Use:   "set",
		Short: "Manage namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Manage namespaces")
		},
	}
	namespaceSetCmd = &cobra.Command{
		Use:   "namespace",
		Short: "set namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Setting the current namespace")
		},
	}
	contextSetCmd = &cobra.Command{
		Use:   "context",
		Short: "set a context",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Setting the current context")
		},
	}
)

func init() {
	setnamespaceCmd.AddCommand(namespaceSetCmd)
	setnamespaceCmd.AddCommand(contextSetCmd)
}