package db

import (
	. "types"
)

type Package struct {
	Container *Container
	Client    *ClientType
	RespChan  chan []byte
}

func (this *DbCore) Processor() {
	for {
		pkg := <- this.Input
		pkg.RespChan <- ProcessQuery(pkg.Client, pkg.Container)
	}
}
