package cmd

import (
	"fmt"

	common "github.com/spideyz0r/kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Aliases: []string{"g", "current"},
		Short: "Get current context or namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please specify a subcommand. Use --help for more details.")
		},
	}
	namespaceGetCmd = &cobra.Command{
		Use:   "namespace",
		Aliases: []string{"ns", "namespaces"},
		Short: "get namespace",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			fmt.Printf("%s\n", common.GetCurrent("namespace"))
		},
	}
	contextGetCmd = &cobra.Command{
		Use:   "context",
		Aliases: []string{"ctx", "contexts"},
		Short: "get a context",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			fmt.Printf("%s\n", common.GetCurrent("context"))
		},
	}
)

func init() {
	getCmd.AddCommand(namespaceGetCmd)
	getCmd.AddCommand(contextGetCmd)
	namespaceGetCmd.Flags().Bool("debug", false, "Enable debug mode")
	contextGetCmd.Flags().Bool("debug", false, "Enable debug mode")
}
