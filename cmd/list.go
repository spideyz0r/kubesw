package cmd

import (
	"fmt"

	common "github.com/spideyz0r/kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Short:   "List all contexts or namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please specify a subcommand. Use --help for more details.")

		},
	}
	namespaceListCmd = &cobra.Command{
		Use:     "namespace",
		Aliases: []string{"ns", "namespaces"},
		Short:   "list namespace",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			common.PrintNamespaces()
		},
	}
	contextListCmd = &cobra.Command{
		Use:     "context",
		Aliases: []string{"ctx", "contexts"},
		Short:   "list a context",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			common.PrintContexts()
		},
	}
)

func init() {
	listCmd.AddCommand(namespaceListCmd)
	listCmd.AddCommand(contextListCmd)
	namespaceListCmd.Flags().Bool("debug", false, "Enable debug mode")
	contextListCmd.Flags().Bool("debug", false, "Enable debug mode")
}
