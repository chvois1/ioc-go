package backend

import (
	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/memdb"
)

type Client struct {
	mdb *memdb.DB
}

func NewClient() *Client {
	return &Client{
		mdb: memdb.New(comparer.DefaultComparer, 0),
	}
}

func (cli *Client) Insert(k []byte, v []byte) error {
	return cli.mdb.Put(k, v)
}

func (cli *Client) FindByID(k []byte) ([]byte, error) {
	return cli.mdb.Get(k)
}
