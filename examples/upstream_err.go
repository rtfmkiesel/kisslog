package main

import (
	logger "github.com/rtfmkiesel/kisslog"
)

func main() {
	if err := logger.InitDefault("myapp"); err != nil {
		panic(err)
	}

	log := logger.New("main")
	log.Info("info hello from the main func")

	if err := sub(); err != nil {
		log.Error(err)
	}
}

func sub() error {
	log := logger.New("sub")
	return log.NewError("error from sub")
}
