package wstorify

import (
	"fmt"
	"slices"
)

type MapListStorage map[string][]*Account

func (storage MapListStorage) Create(category, id string, user *Account) error {
	if _, categoryExists := storage[category]; !categoryExists {
		storage[category] = []*Account{}
	}

	if exists := slices.ContainsFunc(storage[category], func(e *Account) bool {
		return e.Name == id
	}); exists {
		return fmt.Errorf("user with ID '%s' already exists in category '%s'", id, category)
	}
	storage[category] = append(storage[category], user)
	return nil
}

func (storage MapListStorage) Read(category, id string) (*Account, error) {
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

func (storage MapListStorage) Update(category, id string, updatedUser *Account) error {
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

func NewMapListStorage() *MapListStorage {
	return &MapListStorage{}
}
