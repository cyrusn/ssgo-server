package main

import (
	"github.com/cyrusn/ssgo-server/cmd"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.Execute()
}
