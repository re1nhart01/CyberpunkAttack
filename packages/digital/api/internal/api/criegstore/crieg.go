package criegstore

import (
	"sync"

	"github.com/cyberpunkattack/pkg/crieg"
)

type CriegUserCredentials struct {
	Id       int    `json:"id"`
	UserHash string `json:"userHash"`
	Username string `json:"username"`
}

type CriegStore struct {
	Global   *crieg.CriegStorePath[crieg.MapStorage]
	Sessions *crieg.CriegStorePath[crieg.MapListStorage]
}

var CriegStored = crieg.New(&CriegStore{
	Global:   crieg.NewStorePath[crieg.MapStorage](),
	Sessions: crieg.NewStorePath[crieg.MapListStorage](),
}, &crieg.CriegConfig{})

func Crieg() *CriegStore {
	return CriegStored.Group
}

func StoreSelector[T any](collectionName string) T {
	switch collectionName {
	case "global":
		return any(Crieg().Global).(T)
	case "sessions":
		return any(Crieg().Sessions).(T)
	}
	return any(map[string]any{}).(T)
}

func withMutex(cb func() error, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	if err := cb(); err != nil {
		//TODO: log error
	}
}
