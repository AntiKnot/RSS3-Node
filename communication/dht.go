package communication

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"

	"github.com/NaturalSelectionLabs/RSS3-Node/types"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

func SetupDhtDiscovery(ctx context.Context, h host.Host, rendezvous string, bootstrapPeers types.AddrList) error {
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		return err
	}
	logger.Info("Logger Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, peerAddr := range bootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				logger.Warning(err)
			} else {
				logger.Info("Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()

	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, rendezvous)

	peerChan, err := routingDiscovery.FindPeers(ctx, rendezvous)
	if err != nil {
		return err
	}

	for peerNode := range peerChan {
		if peerNode.ID == h.ID() {
			continue
		}
		logger.Debug("Found peer:", peerNode)

		logger.Debug("Connecting to:", peerNode)

		if err := h.Connect(ctx, peerNode); err != nil {
			logger.Error("Connection failed:", err)
		} else {
			logger.Info("Connected to:", peerNode)
		}
	}

	return nil
}
