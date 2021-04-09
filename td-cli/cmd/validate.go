package cmd

import (
	"bytes"
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
			var out bytes.Buffer
			tflint.Stdout = &out
			tflint.Run()
			fmt.Print(out.String())
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
			tflint := exec.Command("tfsec", args[0])
			var out bytes.Buffer
			tflint.Stdout = &out
			tflint.Run()
			fmt.Print(out.String())
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
			tflint := exec.Command("terraform", "fmt", args[0])
			var out bytes.Buffer
			tflint.Stdout = &out
			tflint.Run()
			fmt.Print(out.String())
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
			tflint := exec.Command("terraform-docs", "markdown", "--sort-by-required=true", args[0])
			var out bytes.Buffer
			tflint.Stdout = &out
			tflint.Run()
			fmt.Print(out.String())
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
	RootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	validateCmd.AddCommand(tflintCmd)
	validateCmd.AddCommand(tfsecCmd)
	validateCmd.AddCommand(tffmtCmd)
	validateCmd.AddCommand(tfdocCmd)
}
