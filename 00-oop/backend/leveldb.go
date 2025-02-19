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

func (cli *Client) Insert(k string, v string) error {
	return cli.mdb.Put([]byte(k), []byte(v))
}

func (cli *Client) FindByID(k string) (string, error) {
	b, err := cli.mdb.Get([]byte(k))
	return string(b[:]), err
}
