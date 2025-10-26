package args

import (
	"fmt"
	"os"
	"strings"
)

var ARGS = os.Args[1:]
var ARGS_STRING = strings.Join(ARGS, " ")
var CAPTURE []string

func Fail() {
	fmt.Printf("!!! Unknown arguments: %v\n", ARGS_STRING)
	os.Exit(1)
}
func IsMatch(cmd string, x int) bool {
	if strings.Contains(ARGS_STRING, cmd) {
		pieces := strings.Split(cmd, " ")
		cmd_count := len(pieces)
		arg_count := len(os.Args) - 1
		if (arg_count - cmd_count) == x {
			CAPTURE = os.Args[(cmd_count + 1):]
			return true
		}
		return false

	}

	return false
}
