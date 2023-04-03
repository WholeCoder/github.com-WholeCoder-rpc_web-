package cmd

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Hacker) (*Hacker, error) {
	log.Printf("Receiver hacker body from client: %s", in.Name)

	h, ok := hackers[in.Name]

	if !ok {
		return &Hacker{Name: "Not Found!"}, nil
	}

	return &h, nil
}

var hackers = map[string]Hacker{
	"PaperNaffin":    Hacker{Name: "Hacker"},
	"Snow White":     Hacker{Name: "Coder Chic"},
	"FocalRecursion": Hacker{Name: "Bank Hacker"},
	"Jsldub":         Hacker{Name: "Linux Hacker"},
	"BoarderMimi":    Hacker{Name: "Master Hacker"},
	"Rube5":          Hacker{Name: "All Around Hacker"},
	"Queen":          Hacker{Name: "Info Hacker"},
	"Injection":      Hacker{Name: "Database Hacker"},
	"Alpha Cipher":   Hacker{Name: "Soul Hacker"},
	"SandwichMan":    Hacker{Name: "All Around Hacker"}}

func mainRunRpcServer() {
	fmt.Println("our first grpc server ")
	listener, err := net.Listen("tcp", ":10000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ch := Server{}
	RegisterHackerServiceServer(grpcServer, &ch)

	ch2 := ServerComputer{}
	RegisterComputerServiceServer(grpcServer, &ch2)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
