/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package runcd

import (
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "to run full devcd or with options cd, ms, testms, bs, testbs, etc... ",
	Long:  `to run full devcd or with options cd, ms, testms, bs, testbs, etc...`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("run called")

	// },
}

func AddRunSubCommands() {
	RunCmd.AddCommand(CdCmd)
	RunCmd.AddCommand(BsCmd)
	RunCmd.AddCommand(MsCmd)
	RunCmd.AddCommand(TestmsCmd)
	RunCmd.AddCommand(TestBsCmd)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	AddRunSubCommands()
}
