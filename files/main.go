
package files

import (
	"os"
	"path/filepath"
	"github.com/da99/cli.go/run"
)

// Returns true if the argument exists on the filesystem.
func Is(file_path string) bool {
	_, err := os.Stat(file_path);
	return err == nil
}

// Returns first file that exists and has read access.
//   It takes a shortcut using os.Stat with no error,
//   instead of actually checking read access.
func First(f ...string) string {
	for _, fp := range f {
		_, err := os.Stat(fp);
		if err == nil { return fp }
	}
	return ""
}


// Returns a slice of file paths, w/o directories, 1 level deep.
func List_Shallow_Files(file_path string) []string {
	return run.Cmd_Args("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "f")
}

// The same as List_Shallow_Files, but with ext argument: "*.go.html", "*.js", etc
func List_Shallow_Files_Ext(file_path string, ext string) []string {
	return run.Cmd_Args("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "f", "-iname", ext)
}

// Returns a string slice of directory paths, 1 level deep, ignoring . directories.
func List_Shallow_Dirs(file_path string) []string {
	return run.Cmd_Args("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "d", "-not", "-name", ".*")
}

func Clean_Paths(paths ...string) []string {
	var new_paths []string
	for _, x := range paths {
		new_paths = append(new_paths, filepath.Clean(x))
	}
	return new_paths
}
