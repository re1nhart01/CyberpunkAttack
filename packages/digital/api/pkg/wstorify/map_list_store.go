package wstorify

import (
	"errors"
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

func (storage MapListStorage) BroadcastSpecific(
	category string,
	listOfUsers []string,
	eventName,
	sender,
	channel string,
	data map[string]any,
	onNonExistUser func(event *EssentialEvent) error) error {
	event := NewEssentialEvent(eventName, sender, channel, data)
	eventBytes := IntoBytes(event)
	if storage[category] == nil {
		return errors.New("category is not exists")
	}

	userList := storage[category]
	userMap := make(map[string]*Account, len(userList))
	for _, account := range userList {
		userMap[account.Name] = account
	}

	for _, username := range listOfUsers {
		if user, found := userMap[username]; found {
			if err := user.WsConnection.WriteMessage(1, eventBytes); err != nil {
				// TODO: handle and log error
				return err
			}
		} else {
			if err := onNonExistUser(event); err != nil {
				// TODO: handle and log error
				return err
			}
		}
	}

	return nil
}

func (storage MapListStorage) BroadcastAll(
	category,
	eventName,
	sender,
	channel string,
	data map[string]any) error {
	event := NewEssentialEvent(eventName, sender, channel, data)
	eventBytes := IntoBytes(event)
	if storage[category] == nil {
		return errors.New("category is not exists")
	}
	userList := storage[category]
	for _, user := range userList {
		err := user.WsConnection.WriteMessage(1, eventBytes)
		if err != nil {
			//TODO: handle and log error
			return err
		}
	}

	return nil
}

func NewMapListStorage() *MapListStorage {
	return &MapListStorage{}
}
