package cli

import (
	"flag"
	"fmt"
)

var (
	// Port store the port string for server
	Port string
)

func Start() {
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

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f)
	})
	flag.Parse()
}
