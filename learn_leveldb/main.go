package main

import (
	"log"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/ipfs/go-ds-leveldb"
)

func main() {
	// err := RunPut()
	err := RunGet()
	if err != nil {
		panic(err)
	}
}

func RunPut() error {
	ds, err := leveldb.NewDatastore("./data", nil)
	if err != nil {
		return err
	}
	err = ds.Put(datastore.NewKey("/key/A"), []byte("val_A"))
	if err != nil {
		return err
	}
	err = ds.Put(datastore.NewKey("/key/B"), []byte("val_B"))
	if err != nil {
		return err
	}
	err = ds.Put(datastore.NewKey("/key/C"), []byte("val_C"))
	if err != nil {
		return err
	}
	err = ds.Put(datastore.NewKey("/else/test"), []byte("val_test"))
	if err != nil {
		return err
	}
	return nil
}

func RunGet() error {
	ds, err := leveldb.NewDatastore("./data", nil)
	if err != nil {
		return err
	}
	val, err := ds.Get(datastore.NewKey("/key/A"))
	if err != nil {
		return err
	}
	log.Println(string(val))
	res, err := ds.Query(query.Query{
		Prefix: "/key",
	})
	if err != nil {
		return err
	}
	for {
		t, ok := res.NextSync()
		if !ok {
			log.Println("not ok")
			break
		}
		log.Println(t.Key, string(t.Value))
	}
	return nil
}
