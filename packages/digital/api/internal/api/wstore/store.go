package wstore

import (
	"fmt"
	"sync"

	"github.com/cyberpunkattack/api/service"
	"github.com/cyberpunkattack/pkg/dispatcher"
	"github.com/cyberpunkattack/pkg/wstorify"
)

type UserCredentials struct {
	Id       int    `json:"id"`
	UserHash string `json:"userHash"`
	Username string `json:"username"`
}

type SocketEvent struct {
	Event   string         `json:"event"`
	From    string         `json:"from"`
	Channel string         `json:"channel"`
	Data    map[string]any `json:"data"`
}

type WStorifyStore struct {
	Global   *wstorify.StorePath[wstorify.MapStorage]
	Sessions *wstorify.StorePath[wstorify.MapListStorage]
}

var AllocatedWsStore = wstorify.New(&WStorifyStore{
	Global: wstorify.NewStorePath(
		wstorify.NewMapStore(),
		dispatcher.New(),
		service.InjectableServices{
			Gateway: service.NewGatewayService(),
		}),
	Sessions: wstorify.NewStorePath(
		wstorify.NewMapListStorage(),
		dispatcher.New(),
		service.InjectableServices{},
	),
}, &wstorify.Config{})

func Store() *WStorifyStore {
	return AllocatedWsStore.Group
}

func StoreSelector[T any](collectionName string) T {
	switch collectionName {
	case "global":
		return any(Store().Global).(T)
	case "sessions":
		return any(Store().Sessions).(T)
	}
	return any(map[string]any{}).(T)
}

func withMutex(cb func() error, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	if err := cb(); err != nil {
		//TODO: log error
		fmt.Println("withMutex", err)
	}
}
