package config

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-Node/types"
)

type Config struct {
	RendezvousString string
	BootstrapPeers   types.AddrList
	ListenPort       int
}

const Port = 2233

const ProtocolID = "/chat/1.1.0"

const RendezvousString = "86c2cb14a20a3964e7194c9d5ae7cf0f" // md5 of string "RSS3-Node"

const DiscoveryInterval = time.Minute * 5

const RSS3Pointers = "RSS3Pointers"
