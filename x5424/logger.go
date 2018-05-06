package x5424

import (
	"fmt"
	"io"
	"os"

	"github.com/JustAnotherOrganization/l5424"
)

// Static type checking.
var (
	_ l5424.Log    = &Logger{}
	_ l5424.MinLog = &Logger{}
)

const (
	// DefaultFormat is the default logger format.
	DefaultFormat = `%s : %s`
)

var (
	// TODO: allow for overriding the default formats.
	logFormat = DefaultFormat
)

type (
	// Logger implements an l5424 logger.
	Logger struct {
		out  io.Writer
		err  io.Writer
		cout io.Writer

		severity l5424.SeverityLvl
	}
)

// New returns a new Logger with a default severity level of l5424.NoticeLvl.
func New() *Logger {
	l := Logger{
		out:      os.Stdout,
		err:      os.Stderr,
		severity: l5424.NoticeLvl,
	}
	go l.setCombinedOut()
	return &l
}

func (l *Logger) setCombinedOut() {
	l.cout = io.MultiWriter(l.out, l.err)
}

// SetOut sets the out writer for the Logger (default os.Stdout).
func (l *Logger) SetOut(w io.Writer) {
	l.out = w
	go l.setCombinedOut()
}

// AddWriterToOut adds a writer to the out multiwriter for the Logger.
func (l *Logger) AddWriterToOut(w io.Writer) {
	l.out = io.MultiWriter(l.out, w)
	go l.setCombinedOut()
}

// SetErr sets the err writer for the Logger (default os.Stderr).
func (l *Logger) SetErr(w io.Writer) {
	l.err = w
	go l.setCombinedOut()
}

// AddWriterToErr adds a writer to the err multiwriter for the Logger.
func (l *Logger) AddWriterToErr(w io.Writer) {
	l.err = io.MultiWriter(l.err, w)
	go l.setCombinedOut()
}

// SetLevel sets the logger severity level.
func (l *Logger) SetLevel(lvl l5424.SeverityLvl) {
	l.severity = lvl
}

func (l *Logger) write(w io.Writer, v string) {
	fmt.Fprint(w, v)
}

// Print writes to the logger using the provided severity level.
func (l *Logger) Print(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	l.write(w, fmt.Sprintf(logFormat, lvl.String(), fmt.Sprint(v...)))
}

// Println writes to the logger using the provided severity level, followed by a new line.
func (l *Logger) Println(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	l.write(w, fmt.Sprintln(fmt.Sprintf(logFormat, lvl.String(), fmt.Sprint(v...))))
}

// Printf writes to the logger using the provided severity level and format.
func (l *Logger) Printf(w io.Writer, lvl l5424.SeverityLvl, format string, v ...interface{}) {
	l.write(w, fmt.Sprintf(logFormat, lvl.String(), fmt.Sprintf(format, v...)))
}

// Emergency writes an emergency value to the logger.
func (l *Logger) Emergency(v ...interface{}) {
	l.Print(l.cout, l5424.EmergencyLvl, v...)
}

// Emergencyln writes an emergency value to the logger followed by a new line.
func (l *Logger) Emergencyln(v ...interface{}) {
	l.Println(l.cout, l5424.EmergencyLvl, v...)
}

// Emergencyf writes an emergency value to the logger using the provided format string.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.Printf(l.cout, l5424.EmergencyLvl, format, v...)
}

// Alert writes an alert value to the logger.
func (l *Logger) Alert(v ...interface{}) {
	if l.severity <= l5424.AlertLvl {
		l.Print(l.err, l5424.AlertLvl, v...)
	}
}

// Alertln writes an alert value to the logger followed by a new line.
func (l *Logger) Alertln(v ...interface{}) {
	if l.severity <= l5424.AlertLvl {
		l.Println(l.err, l5424.AlertLvl, v...)
	}
}

// Alertf writes an alert value to the logger using the provided format string.
func (l *Logger) Alertf(format string, v ...interface{}) {
	if l.severity <= l5424.AlertLvl {
		l.Printf(l.err, l5424.AlertLvl, format, v...)
	}
}

// Critical writes a critical value to the logger.
func (l *Logger) Critical(v ...interface{}) {
	if l.severity <= l5424.CritLvl {
		l.Print(l.err, l5424.CritLvl, v...)
	}
}

// Criticalln writes a critical value to the logger followed by a new line.
func (l *Logger) Criticalln(v ...interface{}) {
	if l.severity <= l5424.CritLvl {
		l.Println(l.err, l5424.CritLvl, v...)
	}
}

// Criticalf write a critical value to the logger using the provided format string.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.severity <= l5424.CritLvl {
		l.Printf(l.err, l5424.CritLvl, format, v...)
	}
}

// Error writes an error value to the logger.
func (l *Logger) Error(v ...interface{}) {
	if l.severity <= l5424.ErrorLvl {
		l.Print(l.err, l5424.ErrorLvl, v...)
	}
}

// Errorln writes an error value to the logger follower by a new line.
func (l *Logger) Errorln(v ...interface{}) {
	if l.severity <= l5424.ErrorLvl {
		l.Println(l.err, l5424.ErrorLvl, v...)
	}
}

// Errorf writes an error value to the logger using the provided format string.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.severity <= l5424.ErrorLvl {
		l.Printf(l.err, l5424.ErrorLvl, format, v...)
	}
}

// Warn writes a warning value to the logger.
func (l *Logger) Warn(v ...interface{}) {
	if l.severity <= l5424.WarnLvl {
		l.Print(l.err, l5424.WarnLvl, v...)
	}
}

// Warnln writes a warning value to the logger followed by a new line.
func (l *Logger) Warnln(v ...interface{}) {
	if l.severity <= l5424.WarnLvl {
		l.Println(l.err, l5424.WarnLvl, v...)
	}
}

// Warnf writes a warning value to the logger using the provided format string.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.severity <= l5424.WarnLvl {
		l.Printf(l.err, l5424.WarnLvl, format, v...)
	}
}

// Notice writes a notice value to the logger.
func (l *Logger) Notice(v ...interface{}) {
	if l.severity <= l5424.NoticeLvl {
		l.Print(l.out, l5424.NoticeLvl, v...)
	}
}

// Noticeln writes a notice value to the logger followed by a new line.
func (l *Logger) Noticeln(v ...interface{}) {
	if l.severity <= l5424.NoticeLvl {
		l.Println(l.out, l5424.NoticeLvl, v...)
	}
}

// Noticef writes a notice value to the logger using the provided format string.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.severity <= l5424.NoticeLvl {
		l.Printf(l.out, l5424.NoticeLvl, format, v...)
	}
}

// Info writes an info value to the logger.
func (l *Logger) Info(v ...interface{}) {
	if l.severity <= l5424.InfoLvl {
		l.Print(l.out, l5424.InfoLvl, v...)
	}
}

// Infoln writes an info value to the logger followed by a new line.
func (l *Logger) Infoln(v ...interface{}) {
	if l.severity <= l5424.InfoLvl {
		l.Println(l.out, l5424.InfoLvl, v...)
	}
}

// Infof writes an info value to the logger using the provided format string.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.severity <= l5424.InfoLvl {
		l.Printf(l.out, l5424.InfoLvl, format, v...)
	}
}

// Debug writes a debug value to the logger.
func (l *Logger) Debug(v ...interface{}) {
	if l.severity <= l5424.DebugLvl {
		l.Print(l.out, l5424.DebugLvl, v...)
	}
}

// Debugln writes a debug value to the logger followed by a new line.
func (l *Logger) Debugln(v ...interface{}) {
	if l.severity <= l5424.DebugLvl {
		l.Println(l.out, l5424.DebugLvl, v...)
	}
}

// Debugf writes a debug value to the logger using the provided format string.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.severity <= l5424.DebugLvl {
		l.Printf(l.out, l5424.DebugLvl, format, v...)
	}
}
