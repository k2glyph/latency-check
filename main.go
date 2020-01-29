package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

func dbConn() (db *sql.DB) {
	dbUser := viper.GetString("mysql.username")
	dbPass := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.database")
	dbHost := viper.GetString("mysql.hostname")
	dbPort := viper.GetString("mysql.port")
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func dbExec(file []byte) {
	db := dbConn() //open connection
	requests := strings.Split(string(file), ";")
	requests = requests[:len(requests)-1]
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
}
func mysqlCommand() *cobra.Command {
	var mysqlLatency = &cobra.Command{
		Use:   "mysql-check",
		Short: "Check latency of mysql remote server",
		Long:  `Check latency of mysql remote server with running sql file`,
		Run: func(cmd *cobra.Command, args []string) {
			startTime := time.Now()
			if file, err := ioutil.ReadFile(viper.GetString("mysql.runsqlfile")); err != nil {
				log.Fatal("No run sql file found")
			} else {
				dbExec(file)
			}
			if file, err := ioutil.ReadFile(viper.GetString("mysql.drainsqlfile")); err != nil {
				log.Fatal("No drain sql file found")
			} else {
				dbExec(file)
			}
			endTime := time.Now()
			fmt.Println("Total time took:", endTime.Sub(startTime))
		},
	}
	// mysqlLatency.Flags().StringVar(&hostname, "hostname", "", "hostname of remote mysql server")
	return mysqlLatency
}
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
