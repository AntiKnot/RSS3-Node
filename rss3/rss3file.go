package rss3

import (
	"io"

	"github.com/NaturalSelectionLabs/RSS3-Node/db"
	"github.com/NaturalSelectionLabs/RSS3-Node/storage"
	"github.com/NaturalSelectionLabs/RSS3-Node/types"
)

type RSS3File struct {
	DbHdl   *db.DBHandler
	IpfsHdl *storage.IPFSLayer
}

// NewRSS3File returns the RSS3File handler
func NewRSS3File() (*RSS3File, error) {
	return nil, nil
}

func (*RSS3File) Get(user types.UserID) (io.Reader, error) {
	return nil, nil
}

func (*RSS3File) Set(user types.UserID, f io.Reader) error {
	// update the db and the ipfs
	return nil
}

// Validate checks if the newCid can be upsert to the table
// func Validate(user types.UserID, newCid []byte) (bool, error) {
// 	return true, nil
// }
