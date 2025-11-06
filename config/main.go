package config

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/da99/cli.go/files"
)

func must_exist(str_path string) bool {
	if !files.Is(str_path) {
		fmt.Fprintf(os.Stderr, "Does not exist: %v\n", str_path)
		os.Exit(1)
	}
	return true
}

func Get_Config_Bytes(raw_files ...string) ([]byte, error) {
	file_path := files.First(raw_files...)
	if file_path == "" {
		return nil, nil
	}
	contents, read_err := os.ReadFile(file_path)
	if read_err != nil { return nil, read_err }
	return contents, nil
}

func Get_Config() (map[string]interface{}, error) {
	fin := make(map[string]interface{})

	contents, config_err := Get_Config_Bytes("config.json", "config/main.json")
	if config_err != nil {
		return fin, config_err
	}

	if contents == nil { return fin, nil }

	j_err := json.Unmarshal(contents, &fin)
	if j_err != nil {
		return fin, j_err
	}

	return fin, nil
}



