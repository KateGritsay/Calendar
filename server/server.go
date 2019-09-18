package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"github.com/KateGritsay/Calendar/pkg/calendarpb"

)

type calendarServer struct{}


	func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)

calendarpb.RegisterCalendarServer(grpcServer, &calendarServer{})
grpcServer.Serve(lis)

}

func (s *calendarServer) CreateEvent (ctx context.Context, req calendarpb.CreateEventReq) (*calendarpb.CreateEventRes, error){


}

