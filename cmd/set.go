package cmd

import (
	"fmt"

	common "github.com/spideyz0r/kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	debug  = false
	setCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"s", "switch"},
		Short:   "Set context or namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please specify a subcommand. Use --help for more details.")
		},
	}
	namespaceSetCmd = &cobra.Command{
		Use:     "namespace",
		Aliases: []string{"ns", "namespaces"},
		Short:   "set namespace",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			new_kube_config_path, kubeconfig_kubesw_dir := common.InitialSetup()

			var namespace string
			if len(args) != 1 {
				namespace = common.FzfSelect(common.ListNamespaces())
			}
			if debug {
				fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
			}
			if namespace == "" {
				namespace = args[0]
			}
			kube_config := common.UpdateContext(kubeconfig_kubesw_dir, common.GetCurrent("context"), namespace)
			common.UpdateNamespace(kube_config, namespace)
			history := common.InjectShellHistory(cmd.CalledAs(), namespace)
			common.SpawnShell(kube_config, history)
		},
	}
	contextSetCmd = &cobra.Command{
		Use:     "context",
		Aliases: []string{"ctx", "contexts"},
		Short:   "set a context",
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			common.SetDebug(debug)
			new_kube_config_path, kubeconfig_kubesw_dir := common.InitialSetup()
			if debug {
				fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
			}

			var context string
			if len(args) != 1 {
				context = common.FzfSelect(common.ListContexts())
			}
			if context == "" {
				context = args[0]
			}
			kube_config := common.UpdateContext(kubeconfig_kubesw_dir, context, "default")
			history := common.InjectShellHistory(cmd.CalledAs(), context)
			common.SpawnShell(kube_config, history)
		},
	}
)

func init() {
	setCmd.AddCommand(namespaceSetCmd)
	setCmd.AddCommand(contextSetCmd)
	namespaceSetCmd.Flags().Bool("debug", false, "Enable debug mode")
	contextSetCmd.Flags().Bool("debug", false, "Enable debug mode")
}
