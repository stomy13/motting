package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/server"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "motting", want: "finished\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		server := newMockServer()
		cmd := NewCmdRoot(server)
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

type mockServer struct{}

func newMockServer() *mockServer {
	return &mockServer{}
}

func (*mockServer) RunServer(port string, conargs *server.ConnectArgs) error {
	return nil
}
