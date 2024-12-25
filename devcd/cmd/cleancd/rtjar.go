/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package cleancd

import (
	"devcd/logger"
	"devcd/services"

	"github.com/spf13/cobra"
)

// runtimeCmd represents the runtime command
var rtjarCmd = &cobra.Command{
	Use:   "rtjar",
	Short: "runtime jar clean up",
	Long:  `runtime jar clean up `,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("rtjar called")
		services.CleanRuntimeJars()
	},
}

func init() {

}
