package validate

import (
	"fmt"
	"os"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates a terraform directory recursively",
	Long: `This command recursively validates a terraform directory
using terraform-docs, terraform fmt, tflint, and tfsec`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Trace("validate called")
		logrus.Tracef("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate command or call a subcommand.")
			cmd.Help()
		} else {
			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				logrus.Errorf("Directory %s does not exist!", args[0])
			} else {
				tfdocCmd.Run(cmd, args)
				tffmtCmd.Run(cmd, args)
				tflintCmd.Run(cmd, args)
				tfsecCmd.Run(cmd, args)
			}
		}
	},
}

func init() {
	validateCmd.AddCommand(tflintCmd)
	validateCmd.AddCommand(tfsecCmd)
	validateCmd.AddCommand(tffmtCmd)
	validateCmd.AddCommand(tfdocCmd)
}

func GetCmd() *cobra.Command {
	return validateCmd
}

func GettfdocCmd() *cobra.Command {
	return tfdocCmd
}

func GettffmtCmd() *cobra.Command {
	return tffmtCmd
}

func GettflintCmd() *cobra.Command {
	return tflintCmd
}

func GettfsecCmd() *cobra.Command {
	return tfsecCmd
}
