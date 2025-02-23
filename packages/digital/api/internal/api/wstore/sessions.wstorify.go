package wstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"gameplay"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/pkg/dispatcher"
	"github.com/gorilla/websocket"
	"time"

	"github.com/cyberpunkattack/pkg/wstorify"
)

func ListenSessions(path *wstorify.StorePath[wstorify.MapListStorage]) error {
	store := path.Store
	dp, ok := path.Dispatch.(*dispatcher.Dispatcher)
	if !ok {
		return errors.New("no dispatcher inside")
	}

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
			withMutex(func() error {
				handleSessionBroadcast(message, dp)
				return nil
			}, path.Mutex)
			break
		}
	}
}

func ReadSessionPump(client *wstorify.NewClient, path *wstorify.StorePath[wstorify.MapListStorage]) {
	defer func() {
		path.Unregister <- map[string]string{"name": client.User.Name, "session_id": client.Payload["session_id"].(string)}
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

func handleSessionBroadcast(message []byte, dp *dispatcher.Dispatcher) {
	event := wstorify.EssentialEvent{}
	err := json.Unmarshal(message, &event)

	if err != nil {
		//TODO: log it and revalidate
	}

	executionObject := gameplay.ExecutionObject{
		Creator:   "",
		SessionId: "",
		Session:   models.SessionIM{},
		Now:       time.Now(),
		Action: gameplay.ActionObject{
			Creator:   "",
			SessionId: "",
			Deck:      "",
			Card:      "",
			To:        "",
			From:      "",
			Payload:   nil,
		},
	}

	handler := gameplay.NewHandler(executionObject)
	err := handler.HandleGameplay()
}
