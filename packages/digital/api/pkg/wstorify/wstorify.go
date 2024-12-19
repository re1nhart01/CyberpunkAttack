package wstorify

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Store[T any] *T

type NewClient struct {
	IntoChannel string
	User        *Account
	Payload     map[string]any
}

type Factory[T any] struct {
	Config *Config
	Group  Store[T]
	Mutex  *sync.Mutex
}

type StorePath[T any] struct {
	Store      *T
	Broadcast  chan []byte
	Register   chan *NewClient
	Unregister chan map[string]string
	Mutex      *sync.Mutex
}

type Config struct{}

type Account struct {
	Name            string
	FromGroup       string
	Role            string
	WsConnection    *websocket.Conn
	UserCredentials any
	TTC             time.Time
}

type SocketMessage struct {
	IsSended   bool           `json:"is_sended"`
	ReceivedAt time.Time      `json:"received_at"`
	FromUser   string         `json:"from_user"`
	Group      string         `json:"group"`
	Data       map[string]any `json:"payload"`
}

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

func NewStorePath[T any](store *T) *StorePath[T] {
	return &StorePath[T]{
		Store:      store,
		Broadcast:  make(chan []byte),
		Register:   make(chan *NewClient),
		Unregister: make(chan map[string]string),
		Mutex:      &sync.Mutex{},
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
