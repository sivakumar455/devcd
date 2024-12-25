/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package runcd

import (
	"devcd/services"
	"devcd/utils"

	"github.com/spf13/cobra"
)

// vappCmd represents the vapp command
var TestBsCmd = &cobra.Command{
	Use:   "testbs",
	Short: "to start backing services kafka, zookeeper, couchbase etc... containers",
	Long:  `to start backing services kafka, zookeeper, couchbase etc... containers `,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("vapp called")
	// 	utils.TestUtil()

	// },
}

func init() {

	cmdAppStart, cmdAppStop := utils.GenerateCommandsv2("testbs", services.RunService)
	TestBsCmd.AddCommand(cmdAppStart, cmdAppStop)

}
