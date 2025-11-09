package main

import (
	"errors"

	logger "github.com/rtfmkiesel/kisslog"
)

func main() {
	if err := logger.InitDefault("myapp"); err != nil {
		panic(err)
	}

	log := logger.New("main")

	log.Info("info hello from the main func")
	log.Warning("warning hello from the main func")
	log.Error(errors.New("error hello from the main func"))
	log.Error("error hello from the %s func", "main")
	//log.Fatal(errors.New("fatal hello from the main func"))
	//log.Fatal("fatal hello from the %s func", "main")

	log.Debug("hidden debug hello from the main func")
	logger.FlagDebug = true
	log.Debug("debug hello from the main func")

	sub()
}

func sub() {
	log := logger.New("sub")
	log.Info("info hello from the sub func")
}
