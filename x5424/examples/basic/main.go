package main

import (
	"github.com/justanotherorganization/l5424"
	"github.com/justanotherorganization/l5424/x5424"
)

func main() {
	l := x5424.New(l5424.InfoLvl.String(), nil)
	l.Log(x5424.Severity, l5424.InfoLvl, "Woo we have logging\n")
	l.Log(x5424.Severity, l5424.DebugLvl, "And level filtering works (this should not display)\n")
	l.Log(x5424.Severity, l5424.AlertLvl, x5424.Facility, l5424.FacilityKernel, "Also we have facility levels!\n")
	l.Log(x5424.Facility, l5424.FacilityAlert, "Facility levels can work by themselves too!\n")
}
