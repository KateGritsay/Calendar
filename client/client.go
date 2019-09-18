package main

import (
	"context"
	"fmt"
	calendarpb "github.com/KateGritsay/Calendar/pkg/calendar"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()


	c := calendarpb.NewCalendarClient(cc)
	uuid := createEvent(c)
	event := getEvent(c, uuid)

	fmt.Printf("Created event: %+v\n", event)

	event = updateEvent(c, event)

	fmt.Printf("Updated event description: %+v\n", event)

}
func createEvent(client calendarpb.CalendarClient) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
 	createdAt := time.Now()
	res, err := client.CreateEvent(ctx, &calendarpb.CreateEventReq{
		Event: &calendarpb.Event{
			Title: "Tested",
			Description: "Description",
			CreatedAt: &timestamp.Timestamp{Seconds: createdAt.Unix(), Nanos: 0},
			UserId:1,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return res.GetUUID()
}

func getEvent(client calendarpb.CalendarClient, uuid int64) *calendarpb.Event {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := client.GetEvent(ctx, &calendarpb.GetEventReq{UUID: uuid})

	if err != nil {
		log.Fatal(err)
	}

	return res.GetEvent()
}

func updateEvent(client calendarpb.CalendarClient, event *calendarpb.Event) *calendarpb.Event {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := client.UpdateEvent(ctx, &calendarpb.UpdateEventReq{
		Event: &calendarpb.Event{
			UUID:        event.UUID,
			Title:       event.Title,
			Description: "New description",
			CreatedAt: event.CreatedAt,
		},
	})

	if err != nil {
		log.Fatal(event)
	}

	return res.GetEvent()
}