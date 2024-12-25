/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package runcd

import (
	"devcd/services"
	"devcd/utils"

	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var CdCmd = &cobra.Command{
	Use:   "cd",
	Short: "to run full devcd ",
	Long:  `to run full devcd `,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("cd called")
	// },
}

func init() {
	cmdAppStart, cmdAppStop := utils.GenerateCommandsv2("cd", services.RunService)
	CdCmd.AddCommand(cmdAppStart, cmdAppStop)
}
