package backend_test

type TestBackend struct {
	kv map[string][]byte
}

func NewTestBackend() *TestBackend {
	return &TestBackend{
		kv: make(map[string][]byte),
	}
}

func (t *TestBackend) Put(k []byte, v []byte) error {
	t.kv[string(k[:])] = v
	return nil
}

func (t *TestBackend) Get(k []byte) ([]byte, error) {
	return t.kv[string(k)], nil
}
