package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snackhub/lastrip"
)

func main() {
	// client, err := lastrip.Dial("0.0.0.0:37658")
	// if err != nil {
	// 	panic(err)
	// }
	// client.CreateTask("test:topic", "{\"a\":0, \"b\":true, \"c\":\"\r\t\"}", time.Now().Unix())
	client, err := lastrip.Dial("0.0.0.0:37658")
	if err != nil {
		panic(err)
	}
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(os.Stdout, "> ")
		data, _, err := stdin.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stdout, "error read from stdin: %v", err)
			continue
		}
		if len(data) == 0 {
			continue
		}
		params, err := lastrip.ParseSliceToSlices(data)
		fmt.Fprintf(os.Stdout, "[debug] params: %v\n", params)
		if err != nil {
			fmt.Fprintf(os.Stdout, "invalid input: %v\n", err)
			continue
		}
		cmd := lastrip.NewCommand(params[0])
		err = cmd.Initialize(params[1:])
		if err != nil {
			fmt.Fprintf(os.Stdout, "cmd: %s, wrong parameters: %v\n", params[0], params[1:])
			continue
		}
		client.DoCommand(cmd)
		// client.CreateTask(params[0], params[1], v)
	}
}
