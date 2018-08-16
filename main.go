package main

import (
	"github.com/cyrusn/ssgo-server/cmd"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.Execute()
}
