package wstore

import (
	"fmt"

	"github.com/cyberpunkattack/pkg/wstorify"
	"github.com/gorilla/websocket"
)

func ListenGlobal(path *wstorify.StorePath[wstorify.MapStorage]) error {
	store := path.Store
	for {
		select {
		case client := <-path.Register:
			withMutex(func() error {
				return store.Create(client.User.Name, client.User)
			}, path.Mutex)
			store.Read(client.User.Name)
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

func ReadGlobalPump(client *wstorify.NewClient, path *wstorify.StorePath[wstorify.MapStorage]) {
	defer func() {
		path.Unregister <- map[string]string{"name": client.User.Name}
		client.User.WsConnection.Close()
	}()
	for {
		_, message, err := client.User.WsConnection.ReadMessage()
		if err != nil {
			fmt.Println("ReadGlobalPump error", err)
			break
		}
		path.Broadcast <- message

		msg, err := wstorify.NewEchoSocketMessage(message, client.User.Name, client.IntoChannel)
		client.User.WsConnection.WriteMessage(websocket.TextMessage, msg)
	}
}
