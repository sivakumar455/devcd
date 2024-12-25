/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package install

import (
	"devcd/logger"
	"devcd_ext/services"

	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Install full devc",
	Long:  `Install full devc`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("cd called")
		services.InstallFulldevc()
	},
}

func init() {

}
