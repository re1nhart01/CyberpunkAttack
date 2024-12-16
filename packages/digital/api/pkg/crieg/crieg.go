package crieg

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/net/websocket"
)

type CriegStore[T any] *T

// global: {
//   "name": 0xf13dd
// }

type CriegFactory[T any] struct {
	Config *CriegConfig
	Group  CriegStore[T]
}

type MapStorage map[string]*CriegUser
type MapListStorage map[string]map[string]*CriegUser

type CriegConfig struct{}

type CriegUser struct {
	Name            string
	FromGroup       string
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
		storage[category] = make(map[string]*CriegUser)
	}
	if _, exists := storage[category][id]; exists {
		return fmt.Errorf("user with ID '%s' already exists in category '%s'", id, category)
	}
	storage[category][id] = user
	return nil
}

func (storage MapListStorage) Read(category, id string) (*CriegUser, error) {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		if user, exists := categoryUsers[id]; exists {
			return user, nil
		}
		return nil, fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return nil, fmt.Errorf("category '%s' does not exist", category)
}

func (storage MapListStorage) Update(category, id string, updatedUser *CriegUser) error {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		if _, exists := categoryUsers[id]; exists {
			categoryUsers[id] = updatedUser
			return nil
		}
		return fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return fmt.Errorf("category '%s' does not exist", category)
}

func (storage MapListStorage) Delete(category, id string) error {
	if categoryUsers, categoryExists := storage[category]; categoryExists {
		if _, exists := categoryUsers[id]; exists {
			delete(categoryUsers, id)
			// Remove category if empty
			if len(categoryUsers) == 0 {
				delete(storage, category)
			}
			return nil
		}
		return fmt.Errorf("user with ID '%s' does not exist in category '%s'", id, category)
	}
	return fmt.Errorf("category '%s' does not exist", category)
}

func New[T any](store CriegStore[T], config *CriegConfig) *CriegFactory[T] {
	return &CriegFactory[T]{
		Config: config,
		Group:  store,
	}
}
