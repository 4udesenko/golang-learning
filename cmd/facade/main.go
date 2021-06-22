package main

import (
	"github.com/4udesenko/golang-learning-facade/internal/drive"
	"github.com/4udesenko/golang-learning-facade/internal/facade"
	"github.com/4udesenko/golang-learning-facade/internal/memory"
	"github.com/4udesenko/golang-learning-facade/internal/processor"
)

func main() {
	computer := facade.NewComputerFacade(
		processor.NewCPUProvider(),
		memory.NewRAMProvider(),
		drive.NewSSDProvider(),
	)
	computer.Start()
}
