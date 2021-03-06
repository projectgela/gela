// Copyright 2019 The Gela Authors
// This file is part of the Core Gela infrastructure
// https://gela.io
// Package gelxDAO provides an interface to work with gelx database, including leveldb for masternode and mongodb for SDK node
package gelxDAO

import (
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/ethdb"
)

const defaultCacheLimit = 1024

type GelXDAO interface {
	// for both leveldb and mongodb
	IsEmptyKey(key []byte) bool
	Close()

	// mongodb methods
	HasObject(hash common.Hash, val interface{}) (bool, error)
	GetObject(hash common.Hash, val interface{}) (interface{}, error)
	PutObject(hash common.Hash, val interface{}) error
	DeleteObject(hash common.Hash, val interface{}) error // won't return error if key not found
	GetListItemByTxHash(txhash common.Hash, val interface{}) interface{}
	GetListItemByHashes(hashes []string, val interface{}) interface{}
	DeleteItemByTxHash(txhash common.Hash, val interface{})

	// basic gelx
	InitBulk()
	CommitBulk() error

	// gelx lending
	InitLendingBulk()
	CommitLendingBulk() error

	// leveldb methods
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	NewBatch() ethdb.Batch
}

// use alloc to prevent reference manipulation
func EmptyKey() []byte {
	key := make([]byte, common.HashLength)
	return key
}
