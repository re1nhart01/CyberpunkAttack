package dispatcher

import (
	"slices"
	"strings"
	"sync"
	"time"
)

type Listener struct {
	Name          string
	UniqueName    string
	CallbackFunc  func() (time.Time, error)
	CreatedAt     time.Time
	LastExecution *time.Time
}

type Dispatcher struct {
	listeners []*Listener
	mu        sync.Mutex
}

func (dispatcher *Dispatcher) AddListener(name, uname string, cb func() (time.Time, error)) {
	dispatcher.mu.Lock()
	defer dispatcher.mu.Unlock()
	listener := &Listener{
		Name:          name,
		UniqueName:    uname,
		CallbackFunc:  cb,
		CreatedAt:     time.Now(),
		LastExecution: nil,
	}

	dispatcher.listeners = append(dispatcher.listeners, listener)
}

func (dispatcher *Dispatcher) AddManyListener(data []struct {
	name  string
	uname string
	cb    func() (time.Time, error)
}) {
	for _, v := range data {
		dispatcher.AddListener(v.name, v.uname, v.cb)
	}
}

func (dispatcher *Dispatcher) RemoveGroupListener(name string) {
	dispatcher.listeners = filter(dispatcher.listeners, func(t *Listener, _ int) bool {
		return strings.TrimSpace(name) != strings.TrimSpace(t.Name)
	})
}

func (dispatcher *Dispatcher) RemoveListener(name, uname string) {
	indexOf := slices.IndexFunc(dispatcher.listeners, func(e *Listener) bool {
		return e.Name == name && e.UniqueName == uname
	})
	if indexOf != -1 {
		dispatcher.listeners = append(dispatcher.listeners[:indexOf], dispatcher.listeners[indexOf+1:]...)
	}
}

func (dispatcher *Dispatcher) Execute(uname string, logfunc func(err error)) {
	listeners := dispatcher.listeners
	for _, v := range listeners {
		name := v.UniqueName
		if name == uname {
			now, err := v.CallbackFunc()
			v.LastExecution = &now
			logfunc(err)
		}
	}
}

func (dispatcher *Dispatcher) ExecuteGroup(name string, logfunc func(err error)) {
	listeners := dispatcher.listeners
	for _, v := range listeners {
		if name == v.Name {
			now, err := v.CallbackFunc()
			v.LastExecution = &now
			logfunc(err)
		}
	}
}

func New() *Dispatcher {
	return &Dispatcher{
		listeners: []*Listener{},
	}
}
