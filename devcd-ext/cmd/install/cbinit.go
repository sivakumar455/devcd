/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package install

import (
	"devcd/logger"
	"devcd_ext/couchbase"

	"github.com/spf13/cobra"
)

// cbinitCmd represents the cbinit command
var cbinitCmd = &cobra.Command{
	Use:   "cbinit",
	Short: "Initializing Couchbase cluster",
	Long:  `cb init`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("cbinit called")
		couchbase.CouchbaseInit()
	},
}

func init() {

}
