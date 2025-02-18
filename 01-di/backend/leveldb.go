package backend

import (
	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/memdb"
)

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
