package x5424

import (
	"fmt"

	"github.com/JustAnotherOrganization/l5424"
)

// Static type checking.
var _ l5424.Log = &Entry{}

const (
	// DefaultEntryFormat is the default entry format.
	DefaultEntryFormat = `%s - %s : %s`
)

var (
	// TODO: allow for overriding the default formats.
	entryFormat          = DefaultEntryFormat
	emergencyEntryFormat = fmt.Sprintf(entryFormat, `%s`, `Emergency`, `%s`)
	alertEntryFormat     = fmt.Sprintf(entryFormat, `%s`, `Alert`, `%s`)
	criticalEntryFormat  = fmt.Sprintf(entryFormat, `%s`, `Critical`, `%s`)
	errorEntryFormat     = fmt.Sprintf(entryFormat, `%s`, `Error`, `%s`)
	warnEntryFormat      = fmt.Sprintf(entryFormat, `%s`, `Warning`, `%s`)
	noticeEntryFormat    = fmt.Sprintf(entryFormat, `%s`, `Notice`, `%s`)
	infoEntryFormat      = fmt.Sprintf(entryFormat, `%s`, `Info`, `%s`)
	debugEntryFormat     = fmt.Sprintf(entryFormat, `%s`, `Debug`, `%s`)
)

type (
	// Entry provides facility value access to Logger. This is a separate type to
	// better support having the zero value facility level as l5424.FacilityKernel.
	Entry struct {
		l *Logger

		facility l5424.FacilityLvl
	}
)

// NewEntry returns a new Entry with a severity level matching the above logger.
func (l *Logger) NewEntry(facilityLvl l5424.FacilityLvl) *Entry {
	return &Entry{
		l:        l,
		facility: facilityLvl,
	}
}

// Emergency writes an emergency value to the logger.
func (e *Entry) Emergency(v ...interface{}) {
	e.l.write(e.l.cout, fmt.Sprintf(emergencyEntryFormat, fmt.Sprint(v...)))
}

// Emergencyln writes an emergency value to the logger followed by a new line.
func (e *Entry) Emergencyln(v ...interface{}) {
	e.l.write(e.l.cout, fmt.Sprintln(fmt.Sprintf(emergencyEntryFormat, fmt.Sprint(v...))))
}

// Emergencyf writes an emergency value to the logger using the provided format string.
func (e *Entry) Emergencyf(format string, v ...interface{}) {
	e.l.write(e.l.cout, fmt.Sprintf(emergencyEntryFormat, fmt.Sprintf(format, v...)))
}

// Alert writes an alert value to the logger.
func (e *Entry) Alert(v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.l.write(e.l.cout, fmt.Sprintf(alertEntryFormat, fmt.Sprint(v...)))
	}
}

// Alertln writes an alert value to the logger followed by a new line.
func (e *Entry) Alertln(v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.l.write(e.l.cout, fmt.Sprintln(fmt.Sprintf(alertEntryFormat, fmt.Sprint(v...))))
	}
}

// Alertf writes an alert value to the logger using the provided format string.
func (e *Entry) Alertf(format string, v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.l.write(e.l.cout, fmt.Sprintf(alertEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Critical writes a critical value to the logger.
func (e *Entry) Critical(v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.l.write(e.l.err, fmt.Sprintf(criticalEntryFormat, fmt.Sprint(v...)))
	}
}

// Criticalln writes a critical value to the logger followed by a new line.
func (e *Entry) Criticalln(v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.l.write(e.l.err, fmt.Sprintln(fmt.Sprintf(criticalEntryFormat, fmt.Sprint(v...))))
	}
}

// Criticalf write a critical value to the logger using the provided format string.
func (e *Entry) Criticalf(format string, v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.l.write(e.l.err, fmt.Sprintf(criticalEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Error writes an error value to the logger.
func (e *Entry) Error(v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.l.write(e.l.err, fmt.Sprintf(errorEntryFormat, fmt.Sprint(v...)))
	}
}

// Errorln writes an error value to the logger follower by a new line.
func (e *Entry) Errorln(v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.l.write(e.l.err, fmt.Sprintln(fmt.Sprintf(errorEntryFormat, fmt.Sprint(v...))))
	}
}

// Errorf writes an error value to the logger using the provided format string.
func (e *Entry) Errorf(format string, v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.l.write(e.l.err, fmt.Sprintf(errorEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Warn writes a warning value to the logger.
func (e *Entry) Warn(v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.l.write(e.l.err, fmt.Sprintf(warnEntryFormat, fmt.Sprint(v...)))
	}
}

// Warnln writes a warning value to the logger followed by a new line.
func (e *Entry) Warnln(v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.l.write(e.l.err, fmt.Sprintln(fmt.Sprintf(warnEntryFormat, fmt.Sprint(v...))))
	}
}

// Warnf writes a warning value to the logger using the provided format string.
func (e *Entry) Warnf(format string, v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.l.write(e.l.err, fmt.Sprintf(warnEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Notice writes a notice value to the logger.
func (e *Entry) Notice(v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.l.write(e.l.out, fmt.Sprintf(noticeEntryFormat, fmt.Sprint(v...)))
	}
}

// Noticeln writes a notice value to the logger followed by a new line.
func (e *Entry) Noticeln(v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.l.write(e.l.out, fmt.Sprintln(fmt.Sprintf(noticeEntryFormat, fmt.Sprint(v...))))
	}
}

// Noticef writes a notice value to the logger using the provided format string.
func (e *Entry) Noticef(format string, v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.l.write(e.l.out, fmt.Sprintf(noticeEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Info writes an info value to the logger.
func (e *Entry) Info(v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.l.write(e.l.out, fmt.Sprintf(infoEntryFormat, fmt.Sprint(v...)))
	}
}

// Infoln writes an info value to the logger followed by a new line.
func (e *Entry) Infoln(v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.l.write(e.l.out, fmt.Sprintln(fmt.Sprintf(infoEntryFormat, fmt.Sprint(v...))))
	}
}

// Infof writes an info value to the logger using the provided format string.
func (e *Entry) Infof(format string, v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.l.write(e.l.out, fmt.Sprintf(infoEntryFormat, fmt.Sprintf(format, v...)))
	}
}

// Debug writes a debug value to the logger.
func (e *Entry) Debug(v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.l.write(e.l.out, fmt.Sprintf(infoFormat, fmt.Sprint(v...)))
	}
}

// Debugln writes a debug value to the logger followed by a new line.
func (e *Entry) Debugln(v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.l.write(e.l.out, fmt.Sprintln(fmt.Sprintf(infoFormat, fmt.Sprint(v...))))
	}
}

// Debugf writes a debug value to the logger using the provided format string.
func (e *Entry) Debugf(format string, v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.l.write(e.l.out, fmt.Sprintf(infoFormat, fmt.Sprintf(format, v...)))
	}
}
