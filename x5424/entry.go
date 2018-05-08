package x5424 // import "github.com/JustAnotherOrganization/l5424/x5424"

import (
	"fmt"
	"io"

	"github.com/JustAnotherOrganization/l5424"
)

// Static type checking.
var (
	_ l5424.Log    = &Entry{}
	_ l5424.MinLog = &Entry{}
)

var (
	// TODO: allow for overriding the default formats.
	entryFormat = fmt.Sprintf("%s - %s", "%s", logFormat)
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

// Print writes to the logger using the provided severity level.
func (e *Entry) Print(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	e.l.write(w, lvl, fmt.Sprintf(entryFormat, e.facility.String(), lvl.String(), fmt.Sprint(v...)))
}

// Println writes to the logger using the provided severity level, followed by a new line.
func (e *Entry) Println(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	e.l.write(w, lvl, fmt.Sprintln(fmt.Sprintf(entryFormat, e.facility.String(), lvl.String(), fmt.Sprint(v...))))
}

// Printf writes to the logger using the provided severity level and format.
func (e *Entry) Printf(w io.Writer, lvl l5424.SeverityLvl, format string, v ...interface{}) {
	e.l.write(w, lvl, fmt.Sprintf(entryFormat, e.facility.String(), lvl.String(), fmt.Sprintf(format, v...)))
}

// Emergency writes an emergency value to the logger.
func (e *Entry) Emergency(v ...interface{}) {
	e.Print(e.l.err, l5424.EmergencyLvl, v...)
}

// Emergencyln writes an emergency value to the logger followed by a new line.
func (e *Entry) Emergencyln(v ...interface{}) {
	e.Println(e.l.err, l5424.EmergencyLvl, v...)
}

// Emergencyf writes an emergency value to the logger using the provided format string.
func (e *Entry) Emergencyf(format string, v ...interface{}) {
	e.Printf(e.l.err, l5424.EmergencyLvl, format, v...)
}

// Alert writes an alert value to the logger.
func (e *Entry) Alert(v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.Print(e.l.err, l5424.AlertLvl, v...)
	}
}

// Alertln writes an alert value to the logger followed by a new line.
func (e *Entry) Alertln(v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.Println(e.l.err, l5424.AlertLvl, v...)
	}
}

// Alertf writes an alert value to the logger using the provided format string.
func (e *Entry) Alertf(format string, v ...interface{}) {
	if e.l.severity >= l5424.AlertLvl {
		e.Printf(e.l.err, l5424.AlertLvl, format, v...)
	}
}

// Critical writes a critical value to the logger.
func (e *Entry) Critical(v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.Print(e.l.err, l5424.CritLvl, v...)
	}
}

// Criticalln writes a critical value to the logger followed by a new line.
func (e *Entry) Criticalln(v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.Println(e.l.err, l5424.CritLvl, v...)
	}
}

// Criticalf write a critical value to the logger using the provided format string.
func (e *Entry) Criticalf(format string, v ...interface{}) {
	if e.l.severity >= l5424.CritLvl {
		e.Printf(e.l.err, l5424.CritLvl, format, v...)
	}
}

// Error writes an error value to the logger.
func (e *Entry) Error(v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.Print(e.l.err, l5424.ErrorLvl, v...)
	}
}

// Errorln writes an error value to the logger follower by a new line.
func (e *Entry) Errorln(v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.Println(e.l.err, l5424.ErrorLvl, v...)
	}
}

// Errorf writes an error value to the logger using the provided format string.
func (e *Entry) Errorf(format string, v ...interface{}) {
	if e.l.severity >= l5424.ErrorLvl {
		e.Printf(e.l.err, l5424.ErrorLvl, format, v...)
	}
}

// Warn writes a warning value to the logger.
func (e *Entry) Warn(v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.Print(e.l.err, l5424.WarnLvl, v...)
	}
}

// Warnln writes a warning value to the logger followed by a new line.
func (e *Entry) Warnln(v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.Println(e.l.err, l5424.WarnLvl, v...)
	}
}

// Warnf writes a warning value to the logger using the provided format string.
func (e *Entry) Warnf(format string, v ...interface{}) {
	if e.l.severity >= l5424.WarnLvl {
		e.Printf(e.l.err, l5424.WarnLvl, format, v...)
	}
}

// Notice writes a notice value to the logger.
func (e *Entry) Notice(v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.Print(e.l.out, l5424.NoticeLvl, v...)
	}
}

// Noticeln writes a notice value to the logger followed by a new line.
func (e *Entry) Noticeln(v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.Println(e.l.out, l5424.NoticeLvl, v...)
	}
}

// Noticef writes a notice value to the logger using the provided format string.
func (e *Entry) Noticef(format string, v ...interface{}) {
	if e.l.severity >= l5424.NoticeLvl {
		e.Printf(e.l.out, l5424.NoticeLvl, format, v...)
	}
}

// Info writes an info value to the logger.
func (e *Entry) Info(v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.Print(e.l.out, l5424.InfoLvl, v...)
	}
}

// Infoln writes an info value to the logger followed by a new line.
func (e *Entry) Infoln(v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.Println(e.l.out, l5424.InfoLvl, v...)
	}
}

// Infof writes an info value to the logger using the provided format string.
func (e *Entry) Infof(format string, v ...interface{}) {
	if e.l.severity >= l5424.InfoLvl {
		e.Printf(e.l.out, l5424.InfoLvl, format, v...)
	}
}

// Debug writes a debug value to the logger.
func (e *Entry) Debug(v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.Print(e.l.out, l5424.DebugLvl, v...)
	}
}

// Debugln writes a debug value to the logger followed by a new line.
func (e *Entry) Debugln(v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.Println(e.l.out, l5424.DebugLvl, v...)
	}
}

// Debugf writes a debug value to the logger using the provided format string.
func (e *Entry) Debugf(format string, v ...interface{}) {
	if e.l.severity >= l5424.DebugLvl {
		e.Printf(e.l.out, l5424.DebugLvl, format, v...)
	}
}
