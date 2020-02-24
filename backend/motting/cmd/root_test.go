package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

func TestGetConnectArgs(t *testing.T) {
	cases := []struct {
		dbServer   string
		dbPort     string
		dbSchema   string
		dbLogin    string
		dbPassword string
		want       dbaccess.ConnectArgs
	}{
		{
			dbServer:   "2",
			dbPort:     "3",
			dbSchema:   "4",
			dbLogin:    "5",
			dbPassword: "6",
			want:       dbaccess.ConnectArgs{Address: "2", Port: "3", DBName: "4", User: "5", Password: "6"},
		},
	}

	for _, c := range cases {
		dbServer = c.dbServer
		dbPort = c.dbPort
		dbSchema = c.dbSchema
		dbLogin = c.dbLogin
		dbPassword = c.dbPassword

		conargs := *getConnectArgs()

		if c.want != conargs {
			t.Errorf("unexpected response: want:%+v,	get:%+v", c.want, conargs)
		}
	}
}

func TestNewCmdRoot(t *testing.T) {
	cases := []struct {
		command string
		want    dbaccess.ConnectArgs
	}{
		{command: "motting", want: dbaccess.ConnectArgs{Address: "2", Port: "3", DBName: "4", User: "5", Password: "6"}},
	}

	for _, c := range cases {

		os.Setenv("MOTT_LOG", "-")
		os.Setenv("MOTT_PORT", "1")
		os.Setenv("MOTT_DB_SERVER", "2")
		os.Setenv("MOTT_DB_PORT", "3")
		os.Setenv("MOTT_DB_SCHEMA", "4")
		os.Setenv("MOTT_DB_LOGIN", "5")
		os.Setenv("MOTT_DB_PASSWORD", "6")

		server := newMockServer()
		NewCmdRoot(server)
		conargs := *getConnectArgs()

		if c.want != conargs {
			t.Errorf("unexpected response: want:%+v,	get:%+v", c.want, conargs)
		}
	}
}

func TestExecute(t *testing.T) {
	cases := []struct {
		command string
		want    error
	}{
		{command: "motting", want: nil},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		server := newMockServer()
		cmd := NewCmdRoot(server)
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		err := cmd.Execute()

		if c.want != err {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, err)
		}
	}
}

type mockServer struct{}

func newMockServer() *mockServer {
	return &mockServer{}
}

func (*mockServer) RunServer(port string, conargs *dbaccess.ConnectArgs) error {
	return nil
}
