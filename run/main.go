
package run

import (
	"github.com/da99/cli.go/exit"
	"os/exec"
	"strings"
)


func Lines(cmd_str string) []string {
	raw := exec.Command("sh", "-c", cmd_str)

	output, o_err := raw.Output()
	exit.PrintError(o_err)

	return strings.Split(strings.TrimSpace(string(output)), "\n")
}
