package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	maxConnections     int
	whitelistedClients string
	blacklistedIps     string
}

func configure(options Options) {
	if _, err := os.Stat(options.whitelistedClients); err == nil {
		fmt.Println("Loading whitelisted clients:", options.whitelistedClients)
	} else {
		fmt.Println("Skipping whitelisted clients")
	}
}

func main() {
	maxConnectionsPtr := flag.Int("max-connections", 20, "Max number of TCP connections.")
	whitelistedClientsPtr := flag.String("whitelisted-clients", "", "Whitelist of client types.")
	blacklistedIpsPtr := flag.String("blacklisted-ips", "", "Blacklisted IPs.")
	flag.Parse()

	options := Options{maxConnections: *maxConnectionsPtr,
		whitelistedClients: *whitelistedClientsPtr,
		blacklistedIps:     *blacklistedIpsPtr}

	configure(options)

	SetupRoutesAndStartServer()
}
