package deploy

import (
	"os"
	"tfd/util"
	tf "tfd/util/terraform"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploys a terraform directory",
	Long:  `This command executes Terraform commands to deploy infrastructure.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			//used for flags
			Action       string = viper.GetString("ACTION")
			Path         string = viper.GetString("PATH")
			Workspace    string = viper.GetString("WORKSPACE")
			validActions        = []string{"init", "plan", "apply"}
		)

		logrus.Trace("deploy called")
		logrus.Tracef("Action: %s", Action)
		logrus.Tracef("Path: %s", Path)
		logrus.Tracef("Workspace: %s", Workspace)

		if !util.SliceContains(validActions, Action) {
			logrus.Fatalf("Invalid action provided. Valid actions: %s", validActions)
		}

		if _, err := os.Stat(Path); os.IsNotExist(err) {
			logrus.Fatalf("Invalid path provided. '%s' does not exist!", Path)
		}

		if init := tf.Init(Path); init > 0 {
			logrus.Fatalf("Init returned non zero exit code: %v", init)
		}

		switch Action {
		case "init":
			break
		case "apply":
			if apply := tf.Apply(Path, Workspace); apply > 0 {
				logrus.Fatalf("Apply returned non zero exit code: %v", apply)
			}
			break
		default:
			logrus.Debugf("Defaulting action to 'plan'. Action provided: %s", Action)
			if plan := tf.Plan(Path, Workspace); plan > 0 {
				logrus.Fatalf("Plan returned non zero exit code: %v", plan)
			}
			break
		}
	},
}

func init() {
	deployCmd.Flags().StringP("action", "a", "plan", "Action you wish to execute in the path.")
	deployCmd.Flags().StringP("path", "p", "", "Path to the directory you wish to deploy.")
	deployCmd.Flags().StringP("workspace", "w", "", "Workspace/Environment you wish to deploy.")
	deployCmd.Flags().Bool("auto-apply", false, "Whether running in pipeline or not.")
	viper.BindPFlag("AUTOAPPLY", deployCmd.Flags().Lookup("auto-apply"))
	viper.BindPFlag("ACTION", deployCmd.Flags().Lookup("action"))
	viper.BindPFlag("PATH", deployCmd.Flags().Lookup("path"))
	viper.BindPFlag("WORKSPACE", deployCmd.Flags().Lookup("workspace"))
	viper.SetDefault("AUTOAPPLY", false)
	viper.SetDefault("ACTION", "plan")
	deployCmd.MarkFlagRequired("path")
	deployCmd.MarkFlagRequired("workspace")
}

func GetCmd() *cobra.Command {
	return deployCmd
}
