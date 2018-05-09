package l5424 // import "github.com/JustAnotherOrganization/l5424"

type (
	// Logger is based on the standard logger proposals discussed in detail, linked below
	// https://docs.google.com/document/d/1oTjtY49y8iSxmM9YBaz2NrZIlaXtQsq3nQMd-E0HtwM/edit#
	Logger interface {
		// Log is a flexible log function described in the standard logger proposals.
		Log(...interface{}) error
	}
)

type (
	// FacilityLvl values are not normative however are defined in the RFC. They are provided for convenience.
	FacilityLvl uint
	// SeverityLvl values are not normative however are defined in the RFC. They are provided for convenience.
	SeverityLvl uint
)

const (
	// FacilityKernel should be used for messages originating from the kernel.
	FacilityKernel FacilityLvl = iota
	// FacilityUser should be used for messages originating at the user level.
	FacilityUser
	// FacilityMail should be used for messages originating at a mail system.
	FacilityMail
	// FacilitySystemD should be used for messages originating from within a system daemon.
	FacilitySystemD
	// FacilitySecurity should be used for messages related to security or authorization.
	FacilitySecurity
	// FacilitySyslogD should be used for messages originating from syslogd.
	FacilitySyslogD
	// FacilityPrinter should be used for messages originating from a printer system.
	FacilityPrinter
	// FacilityNews should be used for messages originating from a network news system.
	FacilityNews
	// FacilityUUCP should be used for messages originating from a UUCP system.
	FacilityUUCP
	// FacilityClock should be used for messages originating from a clock daemon.
	FacilityClock
	// FacilityAuthorization should be used for messages related to security or authotization.
	// The RFC makes little distinction between FacilityAuthorization and FacilitySecurity, the
	// difference here is in name only.
	FacilityAuthorization
	// FacilityFTP should be used for messages originating from a FTP daemon.
	FacilityFTP
	// FacilityNTP should be used for messages originating from a NTP daemon.
	FacilityNTP
	// FacilityAudit should be used for informational messages.
	FacilityAudit
	// FacilityAlert should be used for alert messages.
	FacilityAlert
	// FacilityClockD should be used for messages originating from a clock daemon.
	// The RFC makes little distinction between FacilityClockD and FacilityClockD, the
	// difference here is in name only.
	FacilityClockD
	// FacilityLocal0 should be used for messages originating from a local system.
	FacilityLocal0
	// FacilityLocal1 should be used for messages originating from a local system.
	FacilityLocal1
	// FacilityLocal2 should be used for messages originating from a local system.
	FacilityLocal2
	// FacilityLocal3 should be used for messages originating from a local system.
	FacilityLocal3
	// FacilityLocal4 should be used for messages originating from a local system.
	FacilityLocal4
	// FacilityLocal5 should be used for messages originating from a local system.
	FacilityLocal5
	// FacilityLocal6 should be used for messages originating from a local system.
	FacilityLocal6
	// FacilityLocal7 should be used for messages originating from a local system.
	FacilityLocal7
)

const (
	// EmergencyLvl system is unusable
	EmergencyLvl SeverityLvl = iota
	// AlertLvl action must be taken immediately
	AlertLvl
	// CritLvl critical conditions
	CritLvl
	// ErrorLvl error conditions
	ErrorLvl
	// WarnLvl warning conditions
	WarnLvl
	// NoticeLvl normal but signifcant condition
	NoticeLvl
	// InfoLvl informational messages
	InfoLvl
	// DebugLvl debug-level messages
	DebugLvl
)
