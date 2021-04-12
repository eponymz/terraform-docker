package validate

import (
	"fmt"
	"strings"
	"tfd/util"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tfsecCmd = &cobra.Command{
	Use:   "tfsec",
	Short: "validates a terraform directory recursively with tfsec",
	Long: `This subcommand recursively validates a terraform directory
using tfsec`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Trace("tfsec cobra command called")
		logrus.Tracef("Arguments: %s\n", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tfsec command")
			cmd.Help()
		} else {
			except := strings.Split(viper.GetString("IGNORE"), " ")
			tfsec := util.ExecExceptR(except, "tfsec", args[0])
			fmt.Print(tfsec)
		}
	},
}
