package main

import (
	"grpcdemo/pkg/gateway"
	"grpcdemo/pkg/rpc"
	"log"
)

func main() {
	log.Println("hello world!")

	go rpc.Main()
	gateway.Main()
}
