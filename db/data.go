package db

import "google.golang.org/protobuf/proto"

// Sign an outgoing proto message
func (db *DBHandler) SignProtoMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return db.signData(data)
}

// sign binary data using the local node's private key
func (db *DBHandler) signData(data []byte) ([]byte, error) {
	key := db.self.Peerstore().PrivKey(db.self.ID())
	res, err := key.Sign(data)
	return res, err
}

// Authenticate an incoming proto message
func (db *DBHandler) AuthenticateMessage(message proto.Message, data []byte) bool {
	//TODO
	return true
}
