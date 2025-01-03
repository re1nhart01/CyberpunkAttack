package wstore

import (
	"fmt"

	"github.com/cyberpunkattack/pkg/wstorify"
)

func ListenSessions(path *wstorify.StorePath[wstorify.MapListStorage]) error {
	store := path.Store
	for {
		select {
		case client := <-path.Register:
			withMutex(func() error {
				return store.Create(client.Payload["session_id"].(string), client.User.Name, client.User)
			}, path.Mutex)
			break
		case client := <-path.Unregister:
			withMutex(func() error {
				return store.Delete(client["session_id"], client["name"])
			}, path.Mutex)
			break
		case message := <-path.Broadcast:
			fmt.Println(message)
			break
		}
	}
}

func ReadSessionPump() {

}

func handleSessionBroadcast() {

}
