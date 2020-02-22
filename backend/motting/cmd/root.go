package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/MasatoTokuse/motting/motting/server"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "motting",
		Short: "motting",
		Run: func(cmd *cobra.Command, args []string) {
			// avoid not used error
			var err error
			_ = err

			conarg := getConnectArgs()
			err = server.RunServer(port, conarg)
		},
	}
	port       string
	dbServer   string
	dbPort     string
	dbSchema   string
	dbLogin    string
	dbPassword string
	logpath    string
)

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVar(&port, "port", ":3000", "Listen port")
	flags.StringVar(&dbServer, "db_server", "localhost", "db server")
	flags.StringVar(&dbPort, "db_port", "3306", "db port")
	flags.StringVar(&dbSchema, "db_schema", "motting", "db schema")
	flags.StringVar(&dbLogin, "db_login", "motting", "db login")
	flags.StringVar(&dbPassword, "db_password", "motting", "db password")
	flags.StringVar(&logpath, "log", "-", "log file path")
}

func Execute() {

	flags := rootCmd.PersistentFlags()

	viper.SetEnvPrefix("MOTT")
	viper.AutomaticEnv()

	if viper.IsSet("log") {
		flags.Set("log", viper.GetString("log"))
	}
	if logpath != "-" {
		log.Println("log:" + logpath)
		log.SetOutput(&lumberjack.Logger{
			Filename:   logpath,
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		})

		// ログファイルを作成
		// err := lib.CreateFile(logpath)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	os.Exit(1)
		// }
	} else {
		log.Println("log:(stdout)")
	}

	if viper.IsSet("port") {
		flags.Set("port", viper.GetString("port"))
	}
	log.Println("port:" + port)

	if viper.IsSet("db_server") {
		flags.Set("db_server", viper.GetString("db_server"))
	}
	log.Println("db_server:" + dbServer)

	if viper.IsSet("db_port") {
		flags.Set("db_port", viper.GetString("db_port"))
	}
	log.Println("db_port:" + dbPort)

	if viper.IsSet("db_schema") {
		flags.Set("db_schema", viper.GetString("db_schema"))
	}
	log.Println("db_schema:" + dbSchema)

	if viper.IsSet("db_login") {
		flags.Set("db_login", viper.GetString("db_login"))
	}
	log.Println("db_login:" + dbLogin)

	if viper.IsSet("db_password") {
		flags.Set("db_password", viper.GetString("db_password"))
	}
	log.Println("db_password:" + dbPassword)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConnectArgs() *server.ConnectArgs {
	var conarg server.ConnectArgs

	conarg.Address = dbServer
	conarg.Port = dbPort
	conarg.DBName = dbSchema
	conarg.User = dbLogin
	conarg.Password = dbPassword

	return &conarg
}
