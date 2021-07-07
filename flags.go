package main

import (
	"flag"

	cfg "github.com/NaturalSelectionLabs/RSS3-Node/config"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

func ParseFlags() *cfg.Config {
	c := &cfg.Config{}

	flag.Var(&c.BootstrapPeers, "peer", "Adds a peer multiaddress to the bootstrap list")
	flag.StringVar(&c.RendezvousString, "rendezvous", cfg.RendezvousString, "Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.IntVar(&c.ListenPort, "port", cfg.Port, "node listen port")

	flag.Parse()

	if len(c.BootstrapPeers) == 0 {
		c.BootstrapPeers = dht.DefaultBootstrapPeers
	}

	//TODO: read config from file

	return c
}
