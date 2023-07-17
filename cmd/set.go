package cmd

import (
	"fmt"
	common "kubesw/pkg/common"

	"github.com/spf13/cobra"
)

var (
	debug = false
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set context or namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Manage namespaces")
		},
	}
	namespaceSetCmd = &cobra.Command{
		Use:   "namespace",
		Short: "set namespace",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Please specify a single namespace")
				return
			}
			new_kube_config_path, kubeconfig_kubesw_dir := common.InitialSetup()
			if debug {
				fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
			}
			kube_config := common.UpdateContext(kubeconfig_kubesw_dir, common.GetCurrent("context"), args[0])
			common.UpdateNamespace(kube_config, args[0])
			common.SpawnShell(kube_config)
		},
	}
	contextSetCmd = &cobra.Command{
		Use:   "context",
		Short: "set a context",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Please specify a single context")
				return
			}
			new_kube_config_path, kubeconfig_kubesw_dir := common.InitialSetup()
			if debug {
				fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
			}
			kube_config := common.UpdateContext(kubeconfig_kubesw_dir, args[0], "default")
			common.SpawnShell(kube_config)
		},
	}
)

func init() {
	setCmd.AddCommand(namespaceSetCmd)
	setCmd.AddCommand(contextSetCmd)
}