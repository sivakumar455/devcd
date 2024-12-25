/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package runcd

import (
	"devcd/services"
	"devcd/utils"

	"github.com/spf13/cobra"
)

// testmsCmd represents the testms command
var TestmsCmd = &cobra.Command{
	Use:   "testms",
	Short: "to start testms, to start ot stop testing ms which can be configured in global config",
	Long:  `to start testms, to start ot stop testing ms which can be configured in global config`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("testms called")
	// },
}

func init() {

	cmdAppStart, cmdAppStop := utils.GenerateCommandsv2("testms", services.RunService)
	TestmsCmd.AddCommand(cmdAppStart, cmdAppStop)
}
