package validate

import (
	"fmt"
	"os/exec"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var tflintCmd = &cobra.Command{
	Use:   "tflint",
	Short: "validates a terraform directory recursively with tflint",
	Long: `This subcommand recursively validates a terraform directory
using tflint`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("tflint called")
		logrus.Debug("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tflint command")
			cmd.Help()
		} else {
			tflint := exec.Command("tflint", args[0])
			out, _ := tflint.CombinedOutput()
			fmt.Print(string(out))
		}
	},
}

var tfsecCmd = &cobra.Command{
	Use:   "tfsec",
	Short: "validates a terraform directory recursively with tfsec",
	Long: `This subcommand recursively validates a terraform directory
using tfsec`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("tfsec called")
		logrus.Debug("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tfsec command")
			cmd.Help()
		} else {
			tfsec := exec.Command("tfsec", args[0])
			out, _ := tfsec.CombinedOutput()
			fmt.Print(string(out))
		}
	},
}

var tffmtCmd = &cobra.Command{
	Use:   "tffmt",
	Short: "validates a terraform directory recursively with tffmt",
	Long: `This subcommand recursively validates a terraform directory
using tffmt`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("tffmt called")
		logrus.Debug("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tffmt command")
			cmd.Help()
		} else {
			tffmt := exec.Command("terraform", "fmt", args[0])
			out, _ := tffmt.CombinedOutput()
			fmt.Print(string(out))
		}
	},
}

var tfdocCmd = &cobra.Command{
	Use:   "tfdoc",
	Short: "validates a terraform directory recursively with terraform-docs",
	Long: `This subcommand recursively validates a terraform directory
using terraform-docs`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("tfdoc called")
		logrus.Debug("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tfdoc command")
			cmd.Help()
		} else {
			tfdoc := exec.Command("terraform-docs", "markdown", "--sort-by-required=true", args[0])
			out, _ := tfdoc.CombinedOutput()
			fmt.Print(string(out))
		}
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates a terraform directory recursively",
	Long: `This subcommand recursively validates a terraform directory
using terraform-docs, terraform fmt, tflint, and tfsec`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("validate called")
		logrus.Debug("Arguments: %s\n", args)
		logrus.Debug("Flag toggle: %s\n", cmd.Flag("toggle").Value)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate command")
			cmd.Help()
		} else {
			tflintCmd.Run(cmd, args)
			tfsecCmd.Run(cmd, args)
			tfdocCmd.Run(cmd, args)
			tffmtCmd.Run(cmd, args)
		}
	},
}

func init() {
	validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	validateCmd.AddCommand(tflintCmd)
	validateCmd.AddCommand(tfsecCmd)
	validateCmd.AddCommand(tffmtCmd)
	validateCmd.AddCommand(tfdocCmd)
}

func GetCmd() *cobra.Command {
	return validateCmd
}
