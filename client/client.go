package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()


	c := calendarpb.NewCalendarClient(cc)

}