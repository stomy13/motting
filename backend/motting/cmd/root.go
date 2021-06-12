package cmd

import (
	"os"

	"github.com/MasatoTokuse/motting/motting/infrastracture/persistence/mysql"
	"github.com/MasatoTokuse/motting/motting/presenter/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port       string
	dbHostName string
	dbPort     string
	dbName     string
	dbUserName string
	dbPassword string
)

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "motting",
		Short: "motting",
		RunE: func(cmd *cobra.Command, args []string) error {
			connection := mysql.OpenConnection(dbHostName, dbPort, dbName, dbUserName, dbPassword)
			server := server.NewServer(port, connection)
			return server.Run()
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&port, "port", ":3001", "Listen port")
	flags.StringVar(&dbHostName, "db_server", "motting-db", "db server")
	flags.StringVar(&dbPort, "db_port", "3306", "db port")
	flags.StringVar(&dbName, "db_schema", "motting", "db schema")
	flags.StringVar(&dbUserName, "db_login", "motting", "db login")
	flags.StringVar(&dbPassword, "db_password", "motting", "db password")

	viper.SetEnvPrefix("MOTT")
	viper.AutomaticEnv()

	if viper.IsSet("port") {
		flags.Set("port", viper.GetString("port"))
	}
	cmd.Println("port:" + port)

	if viper.IsSet("db_server") {
		flags.Set("db_server", viper.GetString("db_server"))
	}
	cmd.Println("db_server:" + dbHostName)

	if viper.IsSet("db_port") {
		flags.Set("db_port", viper.GetString("db_port"))
	}
	cmd.Println("db_port:" + dbPort)

	if viper.IsSet("db_schema") {
		flags.Set("db_schema", viper.GetString("db_schema"))
	}
	cmd.Println("db_schema:" + dbName)

	if viper.IsSet("db_login") {
		flags.Set("db_login", viper.GetString("db_login"))
	}
	cmd.Println("db_login:" + dbUserName)

	if viper.IsSet("db_password") {
		flags.Set("db_password", viper.GetString("db_password"))
	}
	cmd.Println("db_password:" + dbPassword)

	return cmd
}

func Execute() {
	cmd := NewCmdRoot()

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
