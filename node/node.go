package node

import (
	"context"
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-Node/communication"
	"github.com/NaturalSelectionLabs/RSS3-Node/config"
	"github.com/NaturalSelectionLabs/RSS3-Node/db"
	"github.com/NaturalSelectionLabs/RSS3-Node/validator"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/ipfs/go-log/v2"
)

var logger = log.Logger("rss3-node")

type RSS3Node struct {
	host host.Host
	ctx  context.Context
	cfg  *config.Config
}

// Create a RSS3 node with its specified configs
func NewRSS3Node(cfg *config.Config) (*RSS3Node, error) {
	ctx := context.Background()
	sourceMultiAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", cfg.ListenPort))
	if err != nil {
		return nil, err
	}

	priv, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		return nil, err
	}

	host, err := libp2p.New(ctx,
		libp2p.Identity(priv),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.ForceReachabilityPublic(),
		libp2p.EnableRelay(),
	)
	if err != nil {
		return nil, err
	}

	logger.Info("Host created. We are:", host.ID())
	logger.Info(host.Addrs())

	rn := &RSS3Node{
		host: host,
		ctx:  ctx,
		cfg:  cfg,
	}

	return rn, nil
}

func (rn *RSS3Node) Run() error {
	// 1. Listen commands from the user

	// 2. Set up ipfs node running

	// 3. Set up distributed pointers table
	ps, err := pubsub.NewGossipSub(rn.ctx, rn.host)
	if err != nil {
		return err
	}

	var validator validator.Validator //TODO
	_, err = db.JoinDistributedDB(rn.ctx, ps, rn.host, config.RSS3Pointers, &validator)
	if err != nil {
		return err
	}

	// 4. Set up the communication
	err = communication.SetupMdnsDiscovery(rn.ctx, rn.host, rn.cfg.RendezvousString)
	if err != nil {
		return err
	}

	err = communication.SetupDhtDiscovery(rn.ctx, rn.host, rn.cfg.RendezvousString, rn.cfg.BootstrapPeers)
	if err != nil {
		return err
	}

	//TODO: more communication methods supported

	select {}
}

func (rn *RSS3Node) Stop() error {
	return nil
}
