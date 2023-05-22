package main

import (
	"context"
	"log"
	"os"
	"svc/internal/program"
	"svc/internal/svc"

	"github.com/kardianos/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	svc := svc.New(ctx, cancel)

	svcFlag := len(os.Args) > 1 && os.Args[1] == "service"

	if svcFlag {
		config := &service.Config{
			Name:        "TestService",
			DisplayName: "Test Service",
			Description: "Тестовый сервис программы на Go",
			Arguments:   []string{"service"},
		}

		prg := program.New(svc)

		s, err := service.New(prg, config)
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) > 2 && os.Args[2] == "install" {

			if err := s.Install(); err != nil {
				log.Fatal(err)
			}
			log.Println("Service installed seccesfuly")
			return
		}
		if len(os.Args) > 2 && os.Args[2] == "uninstall" {

			if err := s.Uninstall(); err != nil {
				log.Fatal(err)
			}
			log.Println("Service uninstalled seccesfuly")
			return
		}
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	} else {

		file, err := os.OpenFile("output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		svc.Start()
	}

}
