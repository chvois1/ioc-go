package main

import (
	"fmt"
	"ioc/00-oop/backend"
)

// main is the main entry point
// - init a client API
// - insert a value
// - retreive that value
func main() {
	const key = "key-001"
	const value = "value-001"

	cli := backend.NewClient()
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
