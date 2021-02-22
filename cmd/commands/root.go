package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "Centic",
	Short: "User Service",
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "config.json", "Configuration file to use.")

	viper.SetEnvPrefix("centic")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	RootCmd.RunE = serverCmdF
}
