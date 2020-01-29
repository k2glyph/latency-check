package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func mysqlCommand() *cobra.Command {
	// var hostname string
	var mysqlLatency = &cobra.Command{
		Use:   "mysql",
		Short: "Check latency of mysql remote server",
		Long:  `Check latency of mysql remote server with running sql file`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viber.GetString("hostname"))
			// if viper.GetString("hostname") || hostname == "" {
			// 	log.Fatal("hostname not defined")
			// }
			// fmt.Println(hostname)
		},
	}
	// mysqlLatency.Flags().StringVar(&hostname, "hostname", "", "hostname of remote mysql server")
	return mysqlLatency
}
