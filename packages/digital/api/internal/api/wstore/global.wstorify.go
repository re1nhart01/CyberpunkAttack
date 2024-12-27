package wstore

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cyberpunkattack/api/service"
	"github.com/cyberpunkattack/pkg/dispatcher"
	"github.com/cyberpunkattack/pkg/wstorify"
	"github.com/gorilla/websocket"
)

func ListenGlobal(path *wstorify.StorePath[wstorify.MapStorage]) error {
	store := path.Store
	services, ok := path.Injections.(service.InjectableServices)
	if !ok {
		return errors.New("no injection inside")
	}

	list := services.Gateway.GetAllSubscriptions()
	dp, ok := path.Dispatch.(*dispatcher.Dispatcher)
	if !ok {
		return errors.New("no dispatcher inside")
	}
	dp.AddManyListener(list)

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
			handleBroadcast(message, dp)
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
		fmt.Println(message)
		path.Broadcast <- message

		msg, err := wstorify.NewEchoSocketMessage(message, client.User.Name, client.IntoChannel)
		client.User.WsConnection.WriteMessage(websocket.TextMessage, msg)
	}
}

func handleBroadcast(message []byte, dp *dispatcher.Dispatcher) {
	event := SocketEvent{}
	fmt.Println("zxczx")
	err := json.Unmarshal(message, &event)
	if err != nil {
		//TODO: log it and revalidate
	}

	dp.ExecuteGroup(event.Channel, event.Event, map[string]any{
		"event": event,
	}, func(err error) {
		fmt.Println("a")
	})

}
