package api

import "github.com/cyberpunkattack/pkg/crieg"

type CriegUserCredentials struct {
	Id       int    `json:"id"`
	UserHash string `json:"userHash"`
	Username string `json:"username"`
}

type CriegStore struct {
	global   crieg.MapStorage
	sessions crieg.MapListStorage
}

var CriegStored = crieg.New(&CriegStore{}, &crieg.CriegConfig{})

func Crieg() *CriegStore {
	return CriegStored.Group
}

func (store *CriegStore) Global() crieg.MapStorage {
	return store.global
}

func (store *CriegStore) Sessions() crieg.MapListStorage {
	return store.sessions
}
