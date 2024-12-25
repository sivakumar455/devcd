/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package install

import (
	"devcd_ext/couchbase"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "to install full devcd env or any pre setup needed for devcd env",
	Long:  `to install full devcd env or any pre setup needed for devcd env`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("install called")
	// },
}

func AddRunSubCommands() {
	InstallCmd.AddCommand(cbinitCmd)
	InstallCmd.AddCommand(cdCmd)
	InstallCmd.AddCommand(getmsCmd)
}

func init() {

	// Here you will define your flags and configuration settings.
	couchbase.LoadCbConfig()

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	AddRunSubCommands()
}
