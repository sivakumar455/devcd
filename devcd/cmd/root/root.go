/*
Copyright © 2024 Siva Kumar <EMAIL ADDRESS>
*/
package root

import (
	"devcd/cmd/cleancd"
	runcd "devcd/cmd/run"
	"devcd/config"
	"devcd/logger"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "devc",
	Short: "A CLI tool to run and manage microservice containers in your PC",
	Long: `
	devcd is a CLI tool that helps to run microservice containers in your PC
	through docker or any other container engine`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Name: Siva Kumar")
	},
	Version: "v1.0.1",
	// DisableFlagsInUseLine: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	RootCmd.InitDefaultHelpCmd()
	RootCmd.InitDefaultHelpFlag()

	oldHelpFunc := RootCmd.HelpFunc()
	RootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		oldHelpFunc(cmd, args)
		printDevInfo()
	})

	err := RootCmd.Execute()
	if err != nil {
		logger.Error("Error running devcd", "error", err)
		os.Exit(1)
	}
}

func AddSubCommands() {
	RootCmd.AddCommand(runcd.RunCmd)
	RootCmd.AddCommand(cleancd.CleanCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	err := config.LoadGlobalConfig()
	if err != nil {
		logger.Error("Error loading config")
		os.Exit(1)
	}
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devcd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	AddSubCommands()
}

func printDevInfo() {
	fmt.Printf("\nContact @Siva Kumar Padala for any issues. \n\n")
}
