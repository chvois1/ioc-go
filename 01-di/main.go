package main

import (
	"fmt"
	"ioc/01-di/backend"
	"ioc/01-di/backend_test"
)

// main is the main entry point
// - init a client API
// - insert a value
// - retreive that value
func main() {
	const key = "key-001"
	const value = "value-001"
	db := backend.NewBackend()
	fake := backend_test.NewTestBackend()

	cli := backend.NewClient(db)
	if err := cli.Insert(key, value); err != nil {
		fmt.Printf("cannot insert value: [%s]\n", err.Error())
		return
	}
	v, err := cli.FindByID(key)
	if err != nil {
		fmt.Printf("cannot retreive value from key [%s]: [%s]\n", key, err.Error())
		return
	}
	fmt.Printf("found value [%s] from key [%s]\n", string(v), key)
	cli = backend.NewClient(fake)
	if err := cli.Insert(key, value); err != nil {
		fmt.Printf("cannot insert value: [%s]\n", err.Error())
		return
	}
	v, err = cli.FindByID(key)
	if err != nil {
		fmt.Printf("cannot retreive value from key [%s]: [%s]\n", key, err.Error())
		return
	}
	fmt.Printf("found value [%s] from key [%s]\n", string(v), key)
}
