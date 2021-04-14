package deploy

import (
	// "fmt"
	"os"
	"tfd/util"
	tf "tfd/util/terraform"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Path string
var Action string

var validActions = []string{"init", "plan", "apply", "workspace"}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploys a terraform directory",
	Long: `This command executes Terraform commands to deploy infrastructure.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Trace("deploy called")
		log.Tracef("Action: %s", Action)
		log.Tracef("Path: %s", Path)

		if !util.SliceContains(validActions, Action) {
			log.Fatalf("Invalid action provided. Valid actions: %s", validActions)
		}

		if _, err := os.Stat(Path); os.IsNotExist(err) {
			log.Fatalf("Invalid path provided. '%s' does not exist!", Path)
		}

		if init := tf.Init(Path); init > 0 {
			log.Fatalf("Init returned non zero exit code: %v", init)
		}
	},
}

func init() {
	deployCmd.Flags().StringVarP(&Action, "action", "a", "plan", "Action you wish to execute in the path.")
	deployCmd.Flags().StringVarP(&Path, "path", "p", "", "Path to the directory you wish to deploy.")
	deployCmd.MarkFlagRequired("path")
}

func GetCmd() *cobra.Command {
	return deployCmd
}
