package main

import (
	"flag"

	logger "github.com/rtfmkiesel/kisslog"
)

func main() {
	flag.BoolVar(&logger.FlagDebug, "debug", false, "Show debug output")
	flag.BoolVar(&logger.FlagTime, "showtime", false, "Print messages with timestamps")
	flag.BoolVar(&logger.FlagColor, "color", false, "Print colored messages")
	flag.Parse()

	if err := logger.InitDefault("myapp"); err != nil {
		panic(err)
	}

	log := logger.New("main")
	log.Info("info hello from the main func")
}
