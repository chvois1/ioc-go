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
	store := NewBackend()
	fake := NewTestBackend()
	cli := NewClient(store)
	if err := cli.backend.Insert([]byte(key), []byte(value)); err != nil {
		fmt.Printf("cannot insert value: [%s]\n", err.Error())
		return
	}
	v, err := cli.backend.FindByID([]byte(key))
	if err != nil {
		fmt.Printf("cannot retreive value from key [%s]: [%s]\n", key, err.Error())
		return
	}
	fmt.Printf("found value [%s] from key [%s]\n", string(v), key)
	cli = NewClient(fake)
	if err := cli.backend.Insert([]byte(key), []byte(value)); err != nil {
		fmt.Printf("cannot insert value: [%s]\n", err.Error())
		return
	}
	v, err = cli.backend.FindByID([]byte(key))
	if err != nil {
		fmt.Printf("cannot retreive value from key [%s]: [%s]\n", key, err.Error())
		return
	}
	fmt.Printf("found value [%s] from key [%s]\n", string(v), key)
}

type Storer interface {
	Insert(k []byte, v []byte) error
	FindByID(k []byte) ([]byte, error)
}

type Backend struct {
	mdb *memdb.DB
}

func NewBackend() *Backend {
	return &Backend{
		mdb: memdb.New(comparer.DefaultComparer, 0),
	}
}

func (be *Backend) Insert(k []byte, v []byte) error {
	return be.mdb.Put(k, v)
}

func (be *Backend) FindByID(k []byte) ([]byte, error) {
	return be.mdb.Get(k)
}

type TestBackend struct {
	kv map[string][]byte
}

func NewTestBackend() *TestBackend {
	return &TestBackend{
		kv: make(map[string][]byte),
	}
}

func (t *TestBackend) Insert(k []byte, v []byte) error {
	t.kv[string(k)] = v
	return nil
}

func (t *TestBackend) FindByID(k []byte) ([]byte, error) {
	return t.kv[string(k)], nil
}

type Client struct {
	backend Storer
}

func NewClient(db Storer) *Client {
	return &Client{
		backend: db,
	}
}
