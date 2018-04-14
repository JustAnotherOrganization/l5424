package wrapper

import (
	"fmt"
)

const (
	emergencyLvl = iota
	alertLvl
	critLvl
	errorLvl
	warnLvl
	noticeLvl
	infoLvl
	debugLvl
)

type (
	// LogFunc wraps the given log function.
	LogFunc func(msg string)

	// Logger is a dynamic wrapper for l5424.
	Logger struct {
		// EmergencyLogFunc provides emergency level access to the logger.
		// If EmergencyLogFunc is not set emergency level messages will
		// write to the AlertLogFunc.
		EmergencyLogFunc LogFunc
		// AlertLogFunc provides alert level access to the logger. If
		// AlertLogFunc is not set alert level messages will write to the
		// CriticalLogFunc.
		AlertLogFunc LogFunc
		// CriticalLogFunc provides critical level access to the logger. If
		// CriticalLogFunc is not set critical level messages will write
		// to the ErrorLogFunc.
		CriticalLogFunc LogFunc
		// ErrorLogFunc provides error level access to the logger. If
		// ErrorLogFunc is not set error level messages will write to the
		// WarningLogFunc.
		ErrorLogFunc LogFunc
		// WarningLogFunc provides warning level access to the logger. If
		// WarningLogFunc is not set warning level messages will write to the
		// NoticeLogFunc.
		WarningLogFunc LogFunc
		// NoticeLogFunc provides notice level access to the logger. If
		// NoticeLogFunc is not set notice level messages will write to the
		// InfoLogFunc.
		NoticeLogFunc LogFunc
		// InfoLogFunc provides info level access to the logger. If InfoLogFunc
		// is not set info level messages will write to the DebugLogFunc.
		InfoLogFunc LogFunc
		// DebugLogFunc provides debug level access to the logger. If
		// DebugLogFunc is not set debug level messages will write to the
		// non-leveled logger.
		DebugLogFunc LogFunc
		// FallbackLogFunc provides non-leveled access to the logger.
		// FallbackLogFunc is only needed as a fallback if at least
		// DebugLogFunc is not set. If FallbackLogFunc is called and it
		// is not set the logger will panic.
		FallbackLogFunc LogFunc
	}
)

func (l *Logger) log(lvl int, v string) {
	switch lvl {
	case emergencyLvl:
		if l.EmergencyLogFunc != nil {
			l.EmergencyLogFunc(v)
			return
		}
	case alertLvl:
		if l.AlertLogFunc != nil {
			l.AlertLogFunc(v)
			return
		}
	case critLvl:
		if l.CriticalLogFunc != nil {
			l.CriticalLogFunc(v)
			return
		}
	case errorLvl:
		if l.ErrorLogFunc != nil {
			l.ErrorLogFunc(v)
			return
		}
	case warnLvl:
		if l.WarningLogFunc != nil {
			l.WarningLogFunc(v)
			return
		}
	case noticeLvl:
		if l.NoticeLogFunc != nil {
			l.NoticeLogFunc(v)
			return
		}
	case debugLvl:
		if l.DebugLogFunc != nil {
			l.DebugLogFunc(v)
			return
		}
	default:
		// If this is called and not set a panic will result.
		l.FallbackLogFunc(v)
		return
	}

	l.log(lvl+1, v)
}

// Emergency writes an emergency value to the logger.
func (l *Logger) Emergency(v ...interface{}) {
	l.log(emergencyLvl, fmt.Sprint(v...))
}

// Emergencyln writes an emergency value to the logger followed by a
// new line.
func (l *Logger) Emergencyln(v ...interface{}) {
	l.log(emergencyLvl, fmt.Sprintln(v...))
}

// Emergencyf writes an emergency value to the logger using the
// provided format string.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.log(emergencyLvl, fmt.Sprintf(format, v...))
}

// Alert writes an alert value to the logger.
func (l *Logger) Alert(v ...interface{}) {
	l.log(alertLvl, fmt.Sprint(v...))
}

// Alertln writes an alert value to the logger followed by a new line.
func (l *Logger) Alertln(v ...interface{}) {
	l.log(alertLvl, fmt.Sprintln(v...))
}

// Alertf writes an alert value to the logger using the provided format
// string.
func (l *Logger) Alertf(format string, v ...interface{}) {
	l.log(alertLvl, fmt.Sprintf(format, v...))
}

// Critical writes a critical value to the logger.
func (l *Logger) Critical(v ...interface{}) {
	l.log(critLvl, fmt.Sprint(v...))
}

// Criticalln writes a critical value to the logger followed by a new
// line.
func (l *Logger) Criticalln(v ...interface{}) {
	l.log(critLvl, fmt.Sprintln(v...))
}

// Criticalf write a critical value to the logger using the provided
// format string.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.log(critLvl, fmt.Sprintf(format, v...))
}

// Error writes an error value to the logger.
func (l *Logger) Error(v ...interface{}) {
	l.log(errorLvl, fmt.Sprint(v...))
}

// Errorln writes an error value to the logger follower by a new line.
func (l *Logger) Errorln(v ...interface{}) {
	l.log(errorLvl, fmt.Sprintln(v...))
}

// Errorf writes an error value to the logger using the provided format
// string.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(errorLvl, fmt.Sprintf(format, v...))
}

// Warning writes a warning value to the logger.
func (l *Logger) Warning(v ...interface{}) {
	l.log(warnLvl, fmt.Sprint(v...))
}

// Warningln writes a warning value to the logger followed by a new
// line.
func (l *Logger) Warningln(v ...interface{}) {
	l.log(warnLvl, fmt.Sprintln(v...))
}

// Warningf writes a warning value to the logger using the provided
// format string.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.log(warnLvl, fmt.Sprintf(format, v...))
}

// Notice writes a notice value to the logger.
func (l *Logger) Notice(v ...interface{}) {
	l.log(noticeLvl, fmt.Sprint(v...))
}

// Noticeln writes a notice value to the logger followed by a new line.
func (l *Logger) Noticeln(v ...interface{}) {
	l.log(noticeLvl, fmt.Sprintln(v...))
}

// Noticef writes a notice value to the logger using the provided
// format string.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.log(noticeLvl, fmt.Sprintf(format, v...))
}

// Info writes an info value to the logger.
func (l *Logger) Info(v ...interface{}) {
	l.log(infoLvl, fmt.Sprint(v...))
}

// Infoln writes an info value to the logger followed by a new line.
func (l *Logger) Infoln(v ...interface{}) {
	l.log(infoLvl, fmt.Sprintln(v...))
}

// Infof writes an info value to the logger using the provided format
// string.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(infoLvl, fmt.Sprintf(format, v...))
}

// Debug writes a debug value to the logger.
func (l *Logger) Debug(v ...interface{}) {
	l.log(debugLvl, fmt.Sprint(v...))
}

// Debugln writes a debug value to the logger followed by a new line.
func (l *Logger) Debugln(v ...interface{}) {
	l.log(debugLvl, fmt.Sprintln(v...))
}

// Debugf writes a debug value to the logger using the provided format
// string.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(debugLvl, fmt.Sprintf(format, v...))
}
