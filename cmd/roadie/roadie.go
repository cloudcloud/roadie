// Package main will build the binary for roadie, the lifter and shifter
// of data to temporary locations.
package main

import (
	"github.com/cloudcloud/roadie/pkg/config"
	"github.com/cloudcloud/roadie/pkg/server"
)

func main() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	s := server.New(c)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
