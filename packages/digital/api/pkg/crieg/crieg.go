package crieg

import (
	"fmt"
	"reflect"
	"slices"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type CriegStore[T any] *T

type CriegNewClient struct {
	IntoChannel string
	User        *CriegUser
	Payload     map[string]any
}

type CriegFactory[T any] struct {
	Config *CriegConfig
	Group  CriegStore[T]
	Mutex  *sync.Mutex
}

type CriegStorePath[T any] struct {
	Store      T
	Broadcast  chan any
	Register   chan *CriegNewClient
	Unregister chan map[string]string
	Mutex      *sync.Mutex
}

type MapStorage map[string]*CriegUser
type MapListStorage map[string][]*CriegUser

type CriegConfig struct{}

type CriegUser struct {
	Name            string
	FromGroup       string
	Role            string
	WsConnection    *websocket.Conn
	UserCredentials any
	TTC             time.Time
}

func isMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func (storage MapStorage) Create(id string, user *CriegUser) error {
	if _, exists := storage[id]; exists {
		return fmt.Errorf("user with ID '%s' already exists", id)
	}
	storage[id] = user
	return nil
}

func (storage MapStorage) Read(id string) (*CriegUser, error) {
	if user, exists := storage[id]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user with ID '%s' does not exist", id)
}

func (storage MapStorage) Update(id string, updatedUser *CriegUser) error {
	if _, exists := storage[id]; !exists {
		return fmt.Errorf("user with ID '%s' does not exist", id)
	}
	storage[id] = updatedUser
	return nil
}

func (storage MapStorage) Delete(id string) error {
	if _, exists := storage[id]; !exists {
		return fmt.Errorf("user with ID '%s' does not exist", id)
	}
	delete(storage, id)
	return nil
}

func (storage MapListStorage) Create(category, id string, user *CriegUser) error {
	if _, categoryExists := storage[category]; !categoryExists {
		storage[category] = []*CriegUser{}
	}

	if exists := slices.ContainsFunc(storage[category], func(e *CriegUser) bool {
		return e.Name == id
	}); exists {
		return fmt.Errorf("user with ID '%s' already exists in category '%s'", id, category)
	}
	storage[category] = append(storage[category], user)
	return nil
}

func (storage MapListStorage) Read(category, id string) (*CriegUser, error) {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		for _, user := range categoryUsers {
			if user.Name == id {
				return user, nil
			}
		}
		return nil, fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return nil, fmt.Errorf("category '%s' does not exist", category)
}

func (storage MapListStorage) Update(category, id string, updatedUser *CriegUser) error {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		for i, user := range categoryUsers {
			if user.Name == id {
				categoryUsers[i] = updatedUser
				return nil
			}
		}
		return fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return fmt.Errorf("category '%s' does not exist", category)
}

func (storage MapListStorage) Delete(category, id string) error {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		for i, user := range categoryUsers {
			if user.Name == id { // Assuming CriegUser has an ID field
				storage[category] = append(categoryUsers[:i], categoryUsers[i+1:]...)
				if len(storage[category]) == 0 {
					delete(storage, category)
				}
				return nil
			}
		}
		return fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return fmt.Errorf("category '%s' does not exist", category)
}

func New[T any](store CriegStore[T], config *CriegConfig) *CriegFactory[T] {
	return &CriegFactory[T]{
		Config: config,
		Group:  store,
		Mutex:  &sync.Mutex{},
	}
}

func NewStorePath[T any]() *CriegStorePath[T] {
	return &CriegStorePath[T]{
		Store:      *new(T),
		Broadcast:  make(chan any),
		Register:   make(chan *CriegNewClient),
		Unregister: make(chan map[string]string),
		Mutex:      &sync.Mutex{},
	}
}
