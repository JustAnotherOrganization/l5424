package l5424 // import "github.com/JustAnotherOrganization/l5424"

const (
	// EmergencyLvl system is unusable
	EmergencyLvl = iota
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

type (
	// Log is an rfc5424 log interface.
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
		// Warning writes a warning value to the logger.
		Warning(v ...interface{})
		// Warningln writes a warning value to the logger followed by a new line.
		Warningln(v ...interface{})
		// Warningf writes a warning value to the logger using the provided format string.
		Warningf(format string, v ...interface{})
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
	}
)
