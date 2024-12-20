package dispatcher

import (
	"slices"
	"strings"
	"sync"
	"time"
)

type ListenerCallback func(args map[string]any) (time.Time, error)

type Listener struct {
	Name          string
	UniqueName    string
	CallbackFunc  ListenerCallback
	CreatedAt     time.Time
	LastExecution *time.Time
}

type DispatcherSubscription = []struct {
	Name  string
	Uname string
	Cb    ListenerCallback
}

type Dispatcher struct {
	listeners []*Listener
	mu        sync.Mutex
}

func (dispatcher *Dispatcher) AddListener(name, uname string, cb ListenerCallback) {
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

func (dispatcher *Dispatcher) AddManyListener(data DispatcherSubscription) {
	for _, v := range data {
		dispatcher.AddListener(v.Name, v.Uname, v.Cb)
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

func (dispatcher *Dispatcher) Execute(uname string, args map[string]any, logfunc func(err error)) {
	listeners := dispatcher.listeners
	for _, v := range listeners {
		name := v.UniqueName
		if name == uname {
			now, err := v.CallbackFunc(args)
			v.LastExecution = &now
			logfunc(err)
		}
	}
}

func (dispatcher *Dispatcher) ExecuteGroup(name string, args map[string]any, logfunc func(err error)) {
	listeners := dispatcher.listeners
	for _, v := range listeners {
		if name == v.Name {
			now, err := v.CallbackFunc(args)
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
