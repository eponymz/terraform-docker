package cmd

import (
	"os"
	"strings"

	"tfd/cmd/validate"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func SetLogLevel(level string) {
	var logrusLevel logrus.Level
	insensitiveLevel := strings.ToLower(level)
	switch insensitiveLevel {
	case "trace":
		logrusLevel = logrus.TraceLevel
	case "debug":
		logrusLevel = logrus.DebugLevel
	case "warn":
		logrusLevel = logrus.WarnLevel
	default:
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)
}

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "tfd",
	Short: "tfd is a terraform pipeline tool.",
	Long:  `A Fast and Flexible tool for pipeline validation and application of Terraform.`,
}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	logrus.SetOutput(os.Stdout)
	cobra.OnInitialize(InitConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tfd.yaml)")
	RootCmd.PersistentFlags().StringP("verbosity", "v", "Info", "Verbosity in logging level. E.g. Info, Warn, Debug.")
	RootCmd.AddCommand(validate.GetCmd())
}

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatal(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".tfd")
	}

	viper.SetEnvPrefix("TFD")
	viper.SetDefault("LOGLEVEL", "Info")
	viper.BindPFlag("LOGLEVEL", RootCmd.PersistentFlags().Lookup("verbosity"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}
	SetLogLevel(string(viper.GetString("LOGLEVEL")))
	logrus.Debugf("Logging level: %s\n", logrus.GetLevel().String())
	logrus.Debugf("Viper keys are: %s", viper.AllKeys())
}
