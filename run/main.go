
package run

import (
	"github.com/da99/cli.go/exit"
	"os"
	"os/exec"
	"strings"
	"bytes"
	"bufio"
)


func One_Line_Script(cmd_str string) []string {
	raw := exec.Command("sh", "-c", cmd_str)
	output, o_err := raw.Output()
	exit.PrintError(o_err)

	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

// Returns a string slice of output from a command executed via exec.Command.
func Cmd_Args(cmd string, args ...string) []string {
	os_cmd := exec.Command(cmd, args...)
	var out bytes.Buffer
	os_cmd.Stdout = &out
	os_cmd.Stderr = os.Stderr;
	err := os_cmd.Run()
	if err != nil {
		exit.PrintError(err)
	}
	return Bytes_To_Lines(&out)
}

func Bytes_To_Lines(buf *bytes.Buffer) []string {
	scanner := bufio.NewScanner(buf);
	fin := []string{}
	for scanner.Scan() {
		fin = append(fin, scanner.Text())
	}
	return fin
}
