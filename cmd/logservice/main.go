package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.Registration{
		ServiceName: "Log service",
		ServiceUrl:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), host, port, reg, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
