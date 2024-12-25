/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package runcd

import (
	"devcd/services"
	"devcd/utils"

	"github.com/spf13/cobra"
)

// msCmd represents the ms command
var MsCmd = &cobra.Command{
	Use:   "ms",
	Short: "to start all microservices",
	Long:  `to start all microservices`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("ms called")
	// },
}

func init() {

	cmdAppStart, cmdAppStop := utils.GenerateCommandsv2("ms", services.RunService)
	MsCmd.AddCommand(cmdAppStart, cmdAppStop)
}
