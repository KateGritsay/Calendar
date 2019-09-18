package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"github.com/KateGritsay/Calendar/pkg/calendar"
	"github.com/KateGritsay/Calendar/internal/calendar"
	"time"
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

func mapEventpbToEvent(event *calendarpb.Event) *calendar.Event {
	return &calendar.Event{
		UUID:        event.UUID,
		UserId:      event.UserId,
		Description: event.Description,
		End:         time.Unix(event.End.Seconds, int64(event.End.Nanos)),
		Start:       time.Unix(event.Start.Seconds, int64(event.Start.Nanos)),
		NoticeTime:  event.NoticeTime,
		Title:       event.Title,
	}
}