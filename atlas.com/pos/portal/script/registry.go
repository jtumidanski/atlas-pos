package script

import (
	"errors"
)

type Registry struct {
	registry map[string]Script
}

var registry *Registry

func GetRegistry() *Registry {
	return registry
}

func InitRegistry(scripts ...PortalScript) {
	registry = &Registry{make(map[string]Script)}
	for _, s := range scripts {
		registry.addPortalScript(s)
	}
}

func (s *Registry) GetScript(name string) (*Script, error) {
	if val, ok := s.registry[name]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *Registry) addPortalScript(handler PortalScript) {
	s.registry[handler.Name()] = NewScript(handler)
}
