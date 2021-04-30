package validate

import (
	"fmt"
	"strings"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tfdocCmd = &cobra.Command{
	Use:   "tfdoc",
	Short: "validates a terraform directory recursively with terraform-docs",
	Long: `This subcommand recursively validates a terraform directory
using terraform-docs`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Trace("tfdoc cobra command called")
		logrus.Tracef("Arguments: %s", args)

		if len(args) < 1 {
			fmt.Println("You must pass a directory to validate tfdoc command")
			cmd.Help()
		} else {
			except := strings.Split(viper.GetString("IGNORE"), " ")
			tfdoc := util.ExecExceptRCompare(except, "README.md", "terraform-docs", args[0], "markdown", "--sort-by-required=true")
			fmt.Print(tfdoc)
			if strings.Contains(tfdoc, "returned differences") {
				fmt.Println("Validation Failed!")
			}
		}
	},
}
