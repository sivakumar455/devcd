/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package cleancd

import (
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "to clean volumes, temp logs, rt jars etc... ",
	Long:  `to clean volumes, temp logs, rt jars etc... `,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("clean called")
	// },
}

func AddCleanSubCommands() {
	CleanCmd.AddCommand(rtjarCmd)
	CleanCmd.AddCommand(nxjarCmd)
	CleanCmd.AddCommand(logsCmd)
	CleanCmd.AddCommand(clncdCmd)

}

func init() {
	// rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	AddCleanSubCommands()
}
