package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/microhq/location-api/handler"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.location"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			new(handler.Location),
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
