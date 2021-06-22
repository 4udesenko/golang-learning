package facade

import (
	"fmt"
	"github.com/4udesenko/golang-learning-facade/internal/entity"
)

type cpuProvider interface {
	Process()
}

type ramProvider interface {
	Load() *entity.RAM
}

type ssdProvider interface {
	Read()
}

// ComputerFacade ...
type ComputerFacade interface {
	Start()
}

type computerFacade struct {
	cpuProvider cpuProvider
	ramProvider ramProvider
	ssdProvider ssdProvider
}

// Start ...
func (c *computerFacade) Start() {
	c.cpuProvider.Process()
	ram := c.ramProvider.Load()
	c.ssdProvider.Read()

	fmt.Println(fmt.Sprintf("Computer started! Memory address is %s", ram.Address))
}

// NewComputerFacade ...
func NewComputerFacade(cpu cpuProvider, ram ramProvider, ssd ssdProvider) *computerFacade {
	return &computerFacade{
		cpuProvider: cpu,
		ramProvider: ram,
		ssdProvider: ssd,
	}
}
