package cmd

import (
	"log"

	"golang.org/x/net/context"
)

type ServerComputer struct {
}

func (s *ServerComputer) SayHello(ctx context.Context, in *Computer) (*Computer, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Computer{Body: "Hello from Computer Server " + in.Body + "!"}, nil
}
