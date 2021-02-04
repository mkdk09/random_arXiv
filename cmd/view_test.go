package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "random_arXiv view", want: "show called: optint: 0, optstr: default"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := rootCmd
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		// get := buf.String()
		// if c.want != get {
		// 	t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		// }
	}
}
