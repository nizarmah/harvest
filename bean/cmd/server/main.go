package main

import (
	marketingService "harvest/bean/internal/service/marketing"

	serverDriver "harvest/bean/internal/driver/server"
)

func main() {
	marketing := marketingService.Init()

	server := serverDriver.Init()
	server.Route("/", marketing.Handlers.Landing)
	server.Listen(":8080")
}
