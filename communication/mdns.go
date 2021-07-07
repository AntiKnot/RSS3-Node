package communication

import (
	"context"

	"github.com/NaturalSelectionLabs/RSS3-Node/config"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	logger.Info("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		logger.Warn("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

// SetupMdnsDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func SetupMdnsDiscovery(ctx context.Context, h host.Host, rendezvous string) error {
	disc, err := discovery.NewMdnsService(ctx, h, config.DiscoveryInterval, rendezvous)
	if err != nil {
		return err
	}

	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)
	return nil
}
