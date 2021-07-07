package storage

import (
	"io"
	"types"

	"github.com/NaturalSelectionLabs/RSS3-Node/types"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

type IPFSLayer struct {
	API icore.CoreAPI
}

// NewIPFSLayer returns an instance to handle the ipfs
func NewIPFSLayer(peers types.AddrList) (*IPFSLayer, error) {
	// get a ipfs node running and connects to peernodes
	return nil, nil
}

// Get returns the file corresponding to the cid from ipfs
func (*IPFSLayer) Get(cid []byte) (io.Reader, error) {
	return nil, nil
}

// Add adds a new file corresponding to the cid
func (*IPFSLayer) Add(in io.Reader) error {
	return nil
}

// Delete deletes the file corresponding to the cid
func (*IPFSLayer) Delete(cid []byte) error {
	return nil
}
