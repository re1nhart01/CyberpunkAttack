package criegstore

import (
	"fmt"

	"github.com/cyberpunkattack/pkg/crieg"
)

func ListenGlobal(path *crieg.CriegStorePath[crieg.MapStorage]) error {
	store := path.Store
	for {
		select {
		case client := <-path.Register:
			withMutex(func() error {
				return store.Create(client.User.Name, client.User)
			}, path.Mutex)
			break
		case client := <-path.Unregister:
			withMutex(func() error {
				return store.Delete(client["name"])
			}, path.Mutex)
			break
		case message := <-path.Broadcast:
			fmt.Println(message)
			break
		}
	}
}

func ReadGlobalPump() {

}

func WriteGlobalPump() {

}
