package aerospike_bench

import (
	"fmt"
	. "github.com/aerospike/aerospike-client-go"
	)

func panicOnError(err error) {
	if err != nil {
	panic(err)
	}
}

func PutAndGet(ns string, set string, pk string) {
	// define a client to connect to
	client, err := NewClient("node01", 3000)
	panicOnError(err)

	key, err := NewKey(ns, set, pk)
	panicOnError(err)

	// define some bins with data
	bins := BinMap{
	"bin1": 42,
	"bin2": "An elephant is a mouse with an operating system",
	"bin3": []interface{}{"Go", 2009},
	}

	// write the bins
	err = client.Put(nil, key, bins)
	panicOnError(err)

	// read it back!
	rec, err := client.Get(nil, key)
	panicOnError(err)

	fmt.Printf("Got records by key=%v:\n%#v\n", *key, *rec)

	// delete the key, and check if key exists
	existed, err := client.Delete(nil, key)
	panicOnError(err)
	fmt.Printf("Record existed before delete? %v\n", existed)
}

