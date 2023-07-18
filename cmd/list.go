package cmd

import (
	"fmt"

	common "github.com/spideyz0r/kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all contexts or namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please specify a subcommand. Use --help for more details.")

		},
	}
	namespaceListCmd = &cobra.Command{
		Use:   "namespace",
		Short: "list namespace",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			common.ListNamespaces()
		},
	}
	contextListCmd = &cobra.Command{
		Use:   "context",
		Short: "list a context",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			common.ListContexts()
		},
	}
)

func init() {
	listCmd.AddCommand(namespaceListCmd)
	listCmd.AddCommand(contextListCmd)
	namespaceListCmd.Flags().Bool("debug", false, "Enable debug mode")
	contextListCmd.Flags().Bool("debug", false, "Enable debug mode")
}
