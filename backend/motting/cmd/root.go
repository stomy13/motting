package cmd

import (
	"os"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/server"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port       string
	dbServer   string
	dbPort     string
	dbSchema   string
	dbLogin    string
	dbPassword string
	logpath    string
)

func NewCmdRoot(s server.Serve) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "motting",
		Short: "motting",
		Run: func(cmd *cobra.Command, args []string) {
			// avoid not used error
			var err error
			_ = err

			conargs := getConnectArgs()
			conargs.SetDefault()
			handler := server.NewHandler()
			err = s.RunServer(handler, port)
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&port, "port", ":3001", "Listen port")
	flags.StringVar(&dbServer, "db_server", "motting-db", "db server")
	flags.StringVar(&dbPort, "db_port", "3306", "db port")
	flags.StringVar(&dbSchema, "db_schema", "motting", "db schema")
	flags.StringVar(&dbLogin, "db_login", "motting", "db login")
	flags.StringVar(&dbPassword, "db_password", "motting", "db password")
	flags.StringVar(&logpath, "log", "-", "log file path")

	viper.SetEnvPrefix("MOTT")
	viper.AutomaticEnv()

	if viper.IsSet("log") {
		flags.Set("log", viper.GetString("log"))
	}
	if logpath != "-" {

		cmd.SetOutput(&lumberjack.Logger{
			Filename:   logpath,
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		})
		cmd.Println("log:" + logpath)

		// ログファイルを作成
		// err := lib.CreateFile(logpath)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	os.Exit(1)
		// }./
	} else {
		cmd.Println("log:(stdout)")
	}

	if viper.IsSet("port") {
		flags.Set("port", viper.GetString("port"))
	}
	cmd.Println("port:" + port)

	if viper.IsSet("db_server") {
		flags.Set("db_server", viper.GetString("db_server"))
	}
	cmd.Println("db_server:" + dbServer)

	if viper.IsSet("db_port") {
		flags.Set("db_port", viper.GetString("db_port"))
	}
	cmd.Println("db_port:" + dbPort)

	if viper.IsSet("db_schema") {
		flags.Set("db_schema", viper.GetString("db_schema"))
	}
	cmd.Println("db_schema:" + dbSchema)

	if viper.IsSet("db_login") {
		flags.Set("db_login", viper.GetString("db_login"))
	}
	cmd.Println("db_login:" + dbLogin)

	if viper.IsSet("db_password") {
		flags.Set("db_password", viper.GetString("db_password"))
	}
	cmd.Println("db_password:" + dbPassword)

	return cmd
}

func Execute() {
	server := server.NewServer()
	cmd := NewCmdRoot(server)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func getConnectArgs() *dbaccess.ConnectArgs {
	var conarg dbaccess.ConnectArgs

	conarg.Address = dbServer
	conarg.Port = dbPort
	conarg.DBName = dbSchema
	conarg.User = dbLogin
	conarg.Password = dbPassword

	return &conarg
}
