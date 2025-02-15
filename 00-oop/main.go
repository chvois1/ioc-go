package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/memdb"
)

// main is the main entry point
// - init a client API
// - insert a value
// - retreive that value
func main() {

	const key = "key-001"
	const value = "value-001"
	cli := NewClient()
	if err := cli.Insert([]byte(key), []byte(value)); err != nil {
		fmt.Printf("cannot insert value: [%s]\n", err.Error())
		return
	}
	v, err := cli.FindByID([]byte(key))
	if err != nil {
		fmt.Printf("cannot retreive value from key [%s]: [%s]\n", key, err.Error())
		return
	}
	fmt.Printf("found value [%s] from key [%s]\n", string(v), key)

}

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
