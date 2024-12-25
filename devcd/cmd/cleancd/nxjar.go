/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package cleancd

import (
	"devcd/logger"
	"devcd/services"

	"github.com/spf13/cobra"
)

// nexusCmd represents the nexus command
var nxjarCmd = &cobra.Command{
	Use:   "nxjar",
	Short: "clean nexus jars folder",
	Long:  `clean nexus jars folder`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("nxjar called")
		services.CleanNexusJars()
	},
}

func init() {

}
