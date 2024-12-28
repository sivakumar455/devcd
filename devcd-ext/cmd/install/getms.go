/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package install

import (
	"devcd/logger"
	"devcd_ext/extensions"

	"github.com/spf13/cobra"
)

// getmsCmd represents the getms command
var getmsCmd = &cobra.Command{
	Use:   "getms",
	Short: "to extract ms jars and save it under runtime",
	Long:  `to extract ms jars and save it under runtime`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("getms called")
		extensions.NewExtractMs()
	},
}

func init() {

}
