// Package main will build the binary for roadie, the lifter and shifter
// of data to temporary locations.
package main

import (
	"github.com/cloudcloud/roadie/pkg/config"
	"github.com/cloudcloud/roadie/pkg/server"
)

func main() {
	server.New(config.New()).Start()
}
