package wstorify

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type (
	Store[T any] *T

	NewClient struct {
		IntoChannel string
		User        *Account
		Payload     map[string]any
	}

	Factory[T any] struct {
		Config *Config
		Group  Store[T]
		Mutex  *sync.Mutex
	}

	IDispatch interface {
		Execute(uname string, args map[string]any, logfunc func(err error))
		ExecuteGroup(name string, args map[string]any, logfunc func(err error))
		AddListener(name, uname string, cb func(args map[string]any) (time.Time, error))
	}

	StorePath[T any] struct {
		Store       *T
		ChannelName string
		Broadcast   chan []byte
		Register    chan *NewClient
		Unregister  chan map[string]string
		Mutex       *sync.Mutex
		Dispatch    any
		Injections  any
	}

	Config struct{}

	Account struct {
		Name            string
		FromGroup       string
		Role            string
		WsConnection    *websocket.Conn
		UserCredentials any
		TTC             time.Time
	}

	SocketMessage struct {
		IsSended   bool           `json:"is_sended"`
		ReceivedAt time.Time      `json:"received_at"`
		FromUser   string         `json:"from_user"`
		Group      string         `json:"group"`
		Data       map[string]any `json:"payload"`
	}

	EssentialEvent struct {
		Event   string         `json:"event"`
		From    string         `json:"from"`
		Channel string         `json:"channel"`
		Data    map[string]any `json:"data"`
	}

	InferMethods[T comparable] interface {
		Listen(path *StorePath[T]) error
		ReadPump(client *NewClient, path *StorePath[T])
		handleBroadcast(message []byte, dp *IDispatch)
	}
)

func isMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func New[T any](store Store[T], config *Config) *Factory[T] {
	return &Factory[T]{
		Config: config,
		Group:  store,
		Mutex:  &sync.Mutex{},
	}
}

func NewStorePath[T any](name string, store *T, dispatcher any, injections any) *StorePath[T] {
	return &StorePath[T]{
		ChannelName: name,
		Store:       store,
		Dispatch:    dispatcher,
		Broadcast:   make(chan []byte),
		Register:    make(chan *NewClient),
		Unregister:  make(chan map[string]string),
		Mutex:       &sync.Mutex{},
		Injections:  injections,
	}
}

func IntoBytes[T comparable](data T) []byte {
	bytes, err := json.Marshal(&data)
	if err != nil {
		return make([]byte, 0)
	}
	return bytes
}

func NewEssentialEvent(eventName string, userHash string, channel string, data map[string]any) *EssentialEvent {
	return &EssentialEvent{
		Event:   eventName,
		From:    userHash,
		Channel: channel,
		Data:    data,
	}
}

func NewEchoSocketMessage(message []byte, username, group string) (echoMessage []byte, err error) {
	messageMap := map[string]any{}
	if err := json.Unmarshal(message, &messageMap); err != nil {
		return nil, err
	}

	socketMessage := &SocketMessage{
		IsSended:   true,
		ReceivedAt: time.Now(),
		FromUser:   username,
		Group:      group,
		Data:       messageMap,
	}

	if echoMessage, err = json.Marshal(&socketMessage); err != nil {
		return nil, err
	}

	return echoMessage, nil
}
