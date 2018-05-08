package l5424 // import "github.com/JustAnotherOrganization/l5424"
import "io"

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
	// Disabled disables logging
	Disabled SeverityLvl = 1000
)

type (
	// Log is a RFC5424 log interface.
	Log interface {
		// Emergency writes an emergency value to the logger.
		Emergency(v ...interface{})
		// Emergencyln writes an emergency value to the logger followed by a new line.
		Emergencyln(v ...interface{})
		// Emergencyf writes an emergency value to the logger using the provided format string.
		Emergencyf(format string, v ...interface{})
		// Alert writes an alert value to the logger.
		Alert(v ...interface{})
		// Alertln writes an alert value to the logger followed by a new line.
		Alertln(v ...interface{})
		// Alertf writes an alert value to the logger using the provided format string.
		Alertf(format string, v ...interface{})
		// Critical writes a critical value to the logger.
		Critical(v ...interface{})
		// Criticalln writes a critical value to the logger followed by a new line.
		Criticalln(v ...interface{})
		// Criticalf write a critical value to the logger using the provided format string.
		Criticalf(format string, v ...interface{})
		// Error writes an error value to the logger.
		Error(v ...interface{})
		// Errorln writes an error value to the logger follower by a new line.
		Errorln(v ...interface{})
		// Errorf writes an error value to the logger using the provided format string.
		Errorf(format string, v ...interface{})
		// Warn writes a warning value to the logger.
		Warn(v ...interface{})
		// Warnln writes a warning value to the logger followed by a new line.
		Warnln(v ...interface{})
		// Warnf writes a warning value to the logger using the provided format string.
		Warnf(format string, v ...interface{})
		// Notice writes a notice value to the logger.
		Notice(v ...interface{})
		// Noticeln writes a notice value to the logger followed by a new line.
		Noticeln(v ...interface{})
		// Noticef writes a notice value to the logger using the provided format string.
		Noticef(format string, v ...interface{})
		// Info writes an info value to the logger.
		Info(v ...interface{})
		// Infoln writes an info value to the logger followed by a new line.
		Infoln(v ...interface{})
		// Infof writes an info value to the logger using the provided format string.
		Infof(format string, v ...interface{})
		// Debug writes a debug value to the logger.
		Debug(v ...interface{})
		// Debugln writes a debug value to the logger followed by a new line.
		Debugln(v ...interface{})
		// Debugf writes a debug value to the logger using the provided format string.
		Debugf(format string, v ...interface{})
		// Print writes to the logger using the provided severity level on the provided writer.
		Print(w io.Writer, lvl SeverityLvl, v ...interface{})
		// Println writes to the logger using the provided severity level, followed by a new line on the provided writer.
		Println(w io.Writer, lvl SeverityLvl, v ...interface{})
		// Printf writes to the logger using the provided severity level and format on the provided writer.
		Printf(w io.Writer, lvl SeverityLvl, format string, v ...interface{})
	}

	// MinLog is a minimal RFC5424 log interface.
	MinLog interface {
		// Print writes to the logger using the provided severity level on the provided writer.
		Print(w io.Writer, lvl SeverityLvl, v ...interface{})
		// Println writes to the logger using the provided severity level, followed by a new line on the provided writer.
		Println(w io.Writer, lvl SeverityLvl, v ...interface{})
		// Printf writes to the logger using the provided severity level and format on the provided writer.
		Printf(w io.Writer, lvl SeverityLvl, format string, v ...interface{})
	}
)
