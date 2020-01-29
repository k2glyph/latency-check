package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/k2glyph/latency-check/commands/mysql"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "latency",
		Short: "Check latency of remote machine",
		Long:  `Check latency of remote database server like mysql or just remote servers`,
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.AddCommand(mysqlCommand())

	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			log.Println("config specified but unable to read it, using defaults")
		}
	}
}
