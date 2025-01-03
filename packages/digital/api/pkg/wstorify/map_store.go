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

func (storage MapStorage) BroadcastSpecific(
	listOfUsers []string,
	eventName,
	sender,
	channel string,
	data map[string]any,
	onNonExistUser func(event *EssentialEvent) error) error {
	event := NewEssentialEvent(eventName, sender, channel, data)
	eventBytes := IntoBytes(event)
	for _, v := range listOfUsers {
		user, exists := storage[v]
		if exists {
			err := user.WsConnection.WriteMessage(1, eventBytes)
			if err != nil {
				//TODO: handle and log error
				return err
			}
		} else {
			err := onNonExistUser(event)
			if err != nil {
				//TODO: handle and log error
				return err
			}
		}
	}

	return nil
}

func (storage MapStorage) BroadcastAll(
	eventName,
	sender,
	channel string,
	data map[string]any) error {
	event := NewEssentialEvent(eventName, sender, channel, data)
	eventBytes := IntoBytes(event)
	for _, user := range storage {
		err := user.WsConnection.WriteMessage(1, eventBytes)
		if err != nil {
			//TODO: handle and log error
			return err
		}
	}

	return nil
}

func NewMapStore() *MapStorage {
	return &MapStorage{}
}
