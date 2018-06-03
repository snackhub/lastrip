package main

import (
	"github.com/snackhub/lastrip"
)

func main() {
	server := lastrip.NewServer(&lastrip.ServerConfigure{
		Addr: "0.0.0.0:37658",
	})
	server.Serve()
}
