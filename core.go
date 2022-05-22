package main

import (
	"errors"
	"sync"
)

var ErrorNoSuchKey = errors.New("no such key")

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

// Put function takes two parameters of type string as inputs  and returns an error type
func Put(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

// Get function takes one parameter of type string as input and returns error and string types
func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}
