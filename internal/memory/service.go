package memory

import (
	"fmt"
	"github.com/4udesenko/golang-learning-facade/internal/entity"
)

// RAMProvider ...
type RAMProvider interface {
	Load() *entity.RAM
}

type ramProvider struct{}

// Load ...
func (r *ramProvider) Load() *entity.RAM {
	fmt.Println("RAM processing...")
	hex := "test string"
	return &entity.RAM{
		Address: fmt.Sprint(&hex),
	}
}

// NewRAMProvider ...
func NewRAMProvider() *ramProvider {
	return &ramProvider{}
}
