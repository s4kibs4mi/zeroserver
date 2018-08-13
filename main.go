package main

import (
	"github.com/s4kibs4mi/zeroserver/cmd"
	"github.com/s4kibs4mi/zeroserver/log"
)

func main() {
	log.Init()
	cmd.ServeHttp()
}
