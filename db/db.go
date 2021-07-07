package db

import (
	"context"

	"github.com/NaturalSelectionLabs/RSS3-Node/validator"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type DBHandler struct {
	Messages   chan []byte
	PeerEvents chan *pubsub.PeerEvent

	ctx context.Context

	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription
	evh   *pubsub.TopicEventHandler
	self  host.Host

	TableName string

	UpdateValidator *validator.Validator
}

func JoinDistributedDB(ctx context.Context, ps *pubsub.PubSub, h host.Host, tableName string, validator *validator.Validator) (*DBHandler, error) {
	topic, err := ps.Join(tableName)
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	evh, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	dbh := &DBHandler{
		ctx:             ctx,
		ps:              ps,
		topic:           topic,
		sub:             sub,
		evh:             evh,
		self:            h,
		TableName:       tableName,
		Messages:        make(chan []byte),
		PeerEvents:      make(chan *pubsub.PeerEvent),
		UpdateValidator: validator,
	}

	// TODO:
	// 1. fetch the table from other nodes,
	// 2. listen on the incoming data and peerevents,
	// 3. for mutations, validate and forward

	return dbh, nil
}

// Get returns the value corresponding to the key in dbh
func (dbh *DBHandler) Get(key []byte) ([]byte, error) {
	return nil, nil
}

// Set sets the value to the key in dbh and returns the value
func (dbh *DBHandler) Set(key []byte) ([]byte, error) {
	return nil, nil
}

// Del deletes the entry corresponding to the key in dbh
// and returns that value to it
func (dbh *DBHandler) Del(key []byte) ([]byte, error) {
	return nil, nil
}

func (dbh *DBHandler) listenMessage() {
	// when receiving the new message, it will
}
