package main

import (
	"github.com/nlink-jp/markdown-viewer/cmd"
)

// Version is set at build time
var version = "dev"

func main() {
	cmd.Execute(version)
}
