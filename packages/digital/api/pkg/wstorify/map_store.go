package wstorify

import "fmt"

type MapStorage map[string]*Account

func (storage MapStorage) Create(id string, user *Account) error {
	if _, exists := storage[id]; exists {
		return fmt.Errorf("user with ID '%s' already exists", id)
	}
	storage[id] = user
	return nil
}

func (storage MapStorage) Read(id string) (*Account, error) {
	if user, exists := storage[id]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user with ID '%s' does not exist", id)
}

func (storage MapStorage) Update(id string, updatedUser *Account) error {
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

func NewMapStore() *MapStorage {
	return &MapStorage{}
}
