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
var BsCmd = &cobra.Command{
	Use:   "bs",
	Short: "to start backing services kafka, zookeeper, couchbase etc... containers",
	Long:  `to start backing services kafka, zookeeper, couchbase etc... containers `,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("bs called")
	// 	utils.TestUtil()

	// },
}

func init() {

	cmdAppStart, cmdAppStop := utils.GenerateCommandsv2("bs", services.RunService)
	BsCmd.AddCommand(cmdAppStart, cmdAppStop)

}
