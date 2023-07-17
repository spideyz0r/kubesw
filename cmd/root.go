package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubesw",
	Short: "Switch between kubernetes contexts and namespaces on terminal basis",
	Long:  "Kubeswitch is a dynamic tool designed to enhance your kubernetes workflow by enabling isolated context and namespace management for each terminal",
	Run: func(cmd *cobra.Command, args []string) {
		// Run when no subcommand is specified
		fmt.Println("Please specify a subcommand. Use --help for more details.")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() error {
	// new_kube_config_path, kubeconfig_kubesw_dir := common.InitialSetup()
	// fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
	// fmt.Printf("KUBESWCONFIG: %s\n", kubeconfig_kubesw_dir)
	return rootCmd.Execute()
}