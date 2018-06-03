package main

import (
	"time"

	"github.com/snackhub/lastrip"
)

func main() {
	client, err := lastrip.Dial("0.0.0.0:37658")
	if err != nil {
		panic(err)
	}
	client.CreateTask("test:topic", time.Now().Unix())
}
