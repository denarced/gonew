package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	projectName = kingpin.Arg("name", "Project and directory name").Required().String()
)

func main() {
	kingpin.Parse()
	if pathExists(*projectName) {
		fmt.Fprintf(os.Stderr, "Exists so quitting: %v\n", *projectName)
		os.Exit(1)
	}

	dirs := generateDirPaths(*projectName)
	for _, each := range dirs {
		err := os.MkdirAll(each, os.FileMode(0775))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create dir %v: %v\n", each, err)
			os.Exit(2)
		}
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func generateDirPaths(root string) []string {
	paths := []string{}
	name := filepath.Base(root)
	paths = append(paths, filepath.Join(root, "cmd", name))
	paths = append(paths, filepath.Join(root, "configs"))
	paths = append(paths, filepath.Join(root, "internal", "app", name))
	paths = append(paths, filepath.Join(root, "internal", "pkg"))
	paths = append(paths, filepath.Join(root, "pkg"))
	paths = append(paths, filepath.Join(root, "scripts"))
	return paths
}
