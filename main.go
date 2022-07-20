package main

import (
	"flag"
	"log"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	go p.run()
	return nil
}
func (p *program) run() error {
	logger.Info("I'm Running.")

	<-p.exit // closeされるまで待ち合わせる
	return nil
}
func (p *program) Stop(s service.Service) error {
	logger.Info("I'm Stopping.")
	close(p.exit)
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	nameFlag := flag.String("name", "DummyWindowsService", "Name of service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        *nameFlag,
		DisplayName: "Dummy Windows Service",
		Description: "This service does nothing.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}

		log.Printf("Action(%s) succeeded.\n", *svcFlag)
		return
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
