package handlerregistry

import (
	"fmt"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
)

var Registry = New()

type Handler interface {
	Handle(g *game.Game, c *command.Command, outputCh chan string)
}

type handlerRegistry struct {
	Registry map[string]Handler
	Name     string
}

// New creates a new registry for the given name.
func New() handlerRegistry {
	return handlerRegistry{
		Registry: make(map[string]Handler),
		Name:     "Handler Registry",
	}
}

// Add adds a Handler to the registry for the command if it doesn't
// already have one.
func (r *handlerRegistry) Add(command string, handler Handler) error {
	if _, ok := r.Registry[command]; ok == true {
		return fmt.Errorf("%s: command \"%s\" already has a handler", r.Name, command)
	}
	r.Registry[command] = handler
	return nil
}

// Remove removes a Handler from the registry.
func (r *handlerRegistry) Remove(command string) error {
	if _, ok := r.Registry[command]; !ok {
		return fmt.Errorf("%s: command \"%s\" is not registered", r.Name, command)
	}
	delete(r.Registry, command)
	return nil
}

// Find retrieves a Handler listed under the command.
func (r *handlerRegistry) Find(command string) (Handler, error) {
	if factory, ok := r.Registry[command]; ok {
		return factory, nil
	}
	return nil, fmt.Errorf("%s: command \"%s\" not found", r.Name, command)
}
