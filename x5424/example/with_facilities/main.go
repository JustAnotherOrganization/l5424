package main

import (
	"github.com/JustAnotherOrganization/l5424"
	"github.com/JustAnotherOrganization/l5424/x5424"
)

// logFunc demonstrates passing a facility level logger into a function using the l5424 full interface.
func logFunc(log l5424.Log) {
	log.Noticeln("This is a full implementation of the x5424 logger demonstrating the facility levels found in l5424")
	log.Warnln("One application can log to multiple facility levels")
}

func logFunc2(log l5424.Log) {
	log.Alertln("This way data can more easily be seperated out by more than severity")
	log.Errorln("Please read the full RFC5424 for more information on message priority using facility levels in combination with severity levels")
}

func main() {
	logger := x5424.New()
	logFunc(logger.NewEntry(l5424.FacilityKernel))
	logFunc2(logger.NewEntry(l5424.FacilityAlert))
}
