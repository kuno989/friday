package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "friday",
		Short: "friday",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cfgFile != "" {
				viper.SetConfigFile(cfgFile)
			} else {
				viper.AddConfigPath(".")
				viper.SetConfigName("config")
			}
			viper.AutomaticEnv()
			if err := viper.ReadInConfig(); err == nil {
				log.WithField("location", viper.ConfigFileUsed()).Debug("loaded config file")
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yml)")
	rootCmd.AddCommand(runBackend)
	rootCmd.AddCommand(runFridayEngine)
}
