/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package cleancd

import (
	"devcd/services"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "clean logs folder under runtime",
	Long:  `clean logs folder under runtime`,
	Run: func(cmd *cobra.Command, args []string) {
		services.CleanLogs()
	},
}

func init() {

}
