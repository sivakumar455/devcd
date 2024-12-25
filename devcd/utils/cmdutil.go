package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Function to generate start and stop commands
func GenerateCommands(serviceName string, utilFunc func(string)) (*cobra.Command, *cobra.Command) {
	var cmdStart = &cobra.Command{
		Use:   "start",
		Short: fmt.Sprintf("Starts the %s service", serviceName),
		Run: func(cmd *cobra.Command, args []string) {
			utilFunc("start")
		},
	}

	var cmdStop = &cobra.Command{
		Use:   "stop",
		Short: fmt.Sprintf("Stops the %s service", serviceName),
		Run: func(cmd *cobra.Command, args []string) {
			utilFunc("stop")
		},
	}

	return cmdStart, cmdStop
}

func GenerateCommandsv2(serviceName string, utilFunc func(string, string)) (*cobra.Command, *cobra.Command) {
	var cmdStart = &cobra.Command{
		Use:   "start",
		Short: fmt.Sprintf("Starts the %s service", serviceName),
		Run: func(cmd *cobra.Command, args []string) {
			utilFunc(serviceName, "start")
		},
	}

	var cmdStop = &cobra.Command{
		Use:   "stop",
		Short: fmt.Sprintf("Stops the %s service", serviceName),
		Run: func(cmd *cobra.Command, args []string) {
			utilFunc(serviceName, "stop")
		},
	}

	return cmdStart, cmdStop
}
