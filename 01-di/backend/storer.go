package backend

type Storer interface {
	Put(k []byte, v []byte) error
	Get(k []byte) ([]byte, error)
}

type Client struct {
	db Storer
}

func NewClient(db Storer) *Client {
	return &Client{db}
}

func (cli *Client) Insert(k string, v string) error {
	return cli.db.Put([]byte(k), []byte(v))
}

func (cli *Client) FindByID(k string) (string, error) {
	b, err := cli.db.Get([]byte(k))
	return string(b[:]), err
}
