package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	""
)

type calendarServer struct{}

	func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

grpcServer := grpc.NewServer()

calendarpb.RegisterCalendarServer(grpcServer, &calendarServer{})
grpcServer.Serve(lis)
}