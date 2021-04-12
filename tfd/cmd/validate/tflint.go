package validate

import (
	"fmt"
	"strings"
	"tfd/util"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tflintCmd = &cobra.Command{
	Use:   "tflint",
	Short: "validates a terraform directory recursively with tflint",
	Long: `This subcommand recursively validates a terraform directory
using tflint`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Trace("tflint cobra command called")
		logrus.Tracef("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tflint command")
			cmd.Help()
		} else {
			except := strings.Split(viper.GetString("IGNORE"), " ")
			tflint := util.ExecExceptR(except, "tflint", args[0])
			fmt.Println(tflint)
		}
	},
}
