package main

import (
	"log"
	"net"

	"grpc/plugins/notes"
	"grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	notes := &notes.Notes{}
	proto.RegisterNoteServiceServer(s, notes)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
