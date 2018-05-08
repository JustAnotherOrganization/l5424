package main

import (
	"github.com/JustAnotherOrganization/l5424/x5424"
)

func main() {
	log := x5424.New()
	log.Noticeln("This is a simple implementation of the x5424 logger")
	log.Debugln("By default debug messages are disabled because the logger is initialized at the notice level (so you won't see this)")
	log.Infoln("You also won't see this")
	log.Noticeln("Notice and above will be written to stdout")
	log.Alertln("All message types with a level below notice (warn, error, etc) are written to stderr instead of stdout")
	log.Warnln("That means this message won't be seen if you are only watching stdout")
	log.Errorln("You can change the writers for both sets of log entries")
	log.Emergencyln("Please explore the complete l5424 interface and the full x5424 docs for more")
}
