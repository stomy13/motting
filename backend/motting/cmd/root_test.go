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

func (*mockServer) RunServer(port string, conargs *server.ConnectArgs) error {
	return nil
}
