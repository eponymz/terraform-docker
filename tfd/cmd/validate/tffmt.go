package validate

import (
	"fmt"
	"strings"
	"tfd/util"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tffmtCmd = &cobra.Command{
	Use:   "tffmt",
	Short: "validates a terraform directory recursively with tffmt",
	Long: `This subcommand recursively validates a terraform directory
using tffmt`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Trace("tffmt cobra command called")
		logrus.Tracef("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tffmt command")
			cmd.Help()
		} else {
			except := strings.Split(viper.GetString("IGNORE"), " ")
			tffmt := util.ExecExceptR(except, "terraform fmt", args[0])
			fmt.Print(tffmt)
			if strings.Contains(tffmt, ".tf") {
				fmt.Println("Validation Failed!")
			}
		}
	},
}
