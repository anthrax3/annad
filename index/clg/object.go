package clg

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *index) GetID() spec.ObjectID {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.ID
}

func (i *index) GetType() spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Type
}