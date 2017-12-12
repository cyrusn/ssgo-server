package cli

import (
	"flag"
	"fmt"
)

var (
	// Port store the port string for server
	Port string
)

func init() {
	flag.Usage = func() {
		const welcomeText = "Subject Selection server for LPSS.\nUsage:"
		fmt.Println(welcomeText)
		flag.PrintDefaults()
	}

	const (
		defaultPort = ":5050"
		usagePort   = "server port"
	)

	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+" shorthand")
}

// Parse parses the flag for server
func Parse() {
	flag.Parse()
}
