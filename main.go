package main

import (
	"gobus/di"
)

func main() {
	server := di.Init()
	server.StartServer()
}
