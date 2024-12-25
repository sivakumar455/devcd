/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package cleancd

import (
	"devcd/services"

	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var clncdCmd = &cobra.Command{
	Use:   "cd",
	Short: "clean cd",
	Long:  `clean cd `,
	Run: func(cmd *cobra.Command, args []string) {
		services.CleanFullCd()
	},
}

func init() {

}
