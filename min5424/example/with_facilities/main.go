package main

import (
	"io"
	"os"

	"github.com/JustAnotherOrganization/l5424"
	"github.com/JustAnotherOrganization/l5424/min5424"
)

// logFunc demonstrates passing a facility level logger into a function using the l5424 full interface.
func logFunc(log l5424.MinLog, w io.Writer) {
	log.Println(w, l5424.NoticeLvl, "This is a full implementation of the min5424 logger demonstrating the facility levels found in l5424")
	log.Println(w, l5424.WarnLvl, "One application can log to multiple facility levels")
}

func logFunc2(log l5424.MinLog, w io.Writer) {
	log.Println(w, l5424.AlertLvl, "This way data can more easily be seperated out by more than severity")
	log.Println(w, l5424.ErrorLvl, "Please read the full RFC5424 for more information on message priority using facility levels in combination with severity levels")
}

func main() {
	log1, err := min5424.New(l5424.FacilityKernel.String(), l5424.EmergencyLvl.String())
	if err != nil {
		panic(err)
	}

	logFunc(log1, os.Stdout)

	log2, err := min5424.New(l5424.FacilityAlert.String(), l5424.EmergencyLvl.String())
	if err != nil {
		panic(err)
	}

	logFunc2(log2, os.Stdout)
}
