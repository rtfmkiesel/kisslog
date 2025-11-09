package main

import (
	logger "github.com/rtfmkiesel/kisslog"
)

func main() {
	config := &logger.Config{
		Base:    "myapp",
		TimeStr: "2006-01-02",
		Delim:   ';',
	}

	if err := logger.Init(config); err != nil {
		panic(err)
	}

	log := logger.New("main")
	log.Info("info hello from the main func")
}
