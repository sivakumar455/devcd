/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package install

import (
	"devcd/logger"
	"devcd_ext/extensions"

	"github.com/spf13/cobra"
)

// volumesCmd represents the volumes command
var cdvolCmd = &cobra.Command{
	Use:   "cdvol",
	Short: "clean cd volumes ",
	Long:  `clean cd volumes `,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("volumes called")
		extensions.CleanCdVolumes()
	},
}

func init() {

}
