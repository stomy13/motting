package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/webpush/dbaccess"
	"github.com/MasatoTokuse/motting/webpush/message"
	"github.com/MasatoTokuse/motting/webpush/server"
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

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "webpush",
		Short: "webpush",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			_ = err
			conargs := getConnectArgs()
			conargs.SetDefault()

			pushAt := time.Now().Format("15:04")

			// Fetch users receiving message
			respPushTo, err := http.Get("http://localhost:3001/admin/api/v1/pushTo?pushAt=" + pushAt)
			if err != nil {
				return
			}
			defer respPushTo.Body.Close()

			bytes, err := ioutil.ReadAll(respPushTo.Body)
			if err != nil {
				return
			}

			var pushtimes []model.PushTime
			err = json.Unmarshal(bytes, &pushtimes)
			if err != nil {
				return
			}

			for _, pushtime := range pushtimes {

				// Fetch user's pushed phrase
				respPhrases, err := http.Get("http://localhost:3001/api/v1/phrase?userid=" + pushtime.UserID)
				if err != nil {
					log.Println(err)
					continue
				}
				defer respPhrases.Body.Close()

				bytes, err = ioutil.ReadAll(respPhrases.Body)
				if err != nil {
					log.Println(err)
					continue
				}

				var phrases []model.Phrase
				err = json.Unmarshal(bytes, &phrases)
				if err != nil {
					log.Println(err)
					continue
				}

				// no pushed phrase
				if len(phrases) < 1 {
					continue
				}

				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(phrases))
				phrase := phrases[randIndex]

				message := message.NewMessage(phrase.Author, phrase.Text)
				err = message.Push(pushtime.UserID)
				if err != nil {
					log.Println(err)
				}
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&port, "port", ":3002", "Listen port")
	flags.StringVar(&dbServer, "db_server", "motting-db", "db server")
	flags.StringVar(&dbPort, "db_port", "33306", "db port")
	flags.StringVar(&dbSchema, "db_schema", "webpush", "db schema")
	flags.StringVar(&dbLogin, "db_login", "webpush", "db login")
	flags.StringVar(&dbPassword, "db_password", "webpush", "db password")
	flags.StringVar(&logpath, "log", "-", "log file path")

	viper.SetEnvPrefix("PUSH")
	viper.AutomaticEnv()

	if viper.IsSet("log") {
		flags.Set("log", viper.GetString("log"))
	}
	if logpath != "-" {

		log.SetOutput(&lumberjack.Logger{
			Filename:   logpath,
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		})
		log.Println("log:" + logpath)

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

	return cmd
}

func Execute() {
	server := server.NewServer()
	cmd := NewCmdRoot()
	cmd.AddCommand(NewCmdAuth(server))

	if err := cmd.Execute(); err != nil {
		log.Println(err)
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
