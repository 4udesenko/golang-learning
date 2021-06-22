package processor

import "fmt"

// CPUProvider ...
type CPUProvider interface {
	Process()
}

type cpuProvider struct{}

// Process ...
func (c *cpuProvider) Process() {
	fmt.Println("CPU processing...")
}

// NewCPUProvider ...
func NewCPUProvider() *cpuProvider {
	return &cpuProvider{}
}
