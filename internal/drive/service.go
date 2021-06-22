package drive

import "fmt"

// SSDProvider ...
type SSDProvider interface {
	Read()
}

type ssdProvider struct{}

// Read ...
func (c *ssdProvider) Read() {
	fmt.Println("SSD processing...")
}

// NewSSDProvider ...
func NewSSDProvider() *ssdProvider {
	return &ssdProvider{}
}
