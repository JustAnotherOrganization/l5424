package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/JustAnotherOrganization/l5424"
)

// Static type checking.
var _ l5424.Log = &Logger{}

const (
	// DefaultFormat is the default logger format.
	DefaultFormat = `%s : %s`
)

var (
	logFormat       = DefaultFormat
	emergencyFormat = fmt.Sprintf(logFormat, `Emergency`, `%s`)
	alertFormat     = fmt.Sprintf(logFormat, `Alert`, `%s`)
	criticalFormat  = fmt.Sprintf(logFormat, `Critical`, `%s`)
	errorFormat     = fmt.Sprintf(logFormat, `Error`, `%s`)
	warnFormat      = fmt.Sprintf(logFormat, `Warning`, `%s`)
	noticeFormat    = fmt.Sprintf(logFormat, `Notice`, `%s`)
	infoFormat      = fmt.Sprintf(logFormat, `Info`, `%s`)
	debugFormat     = fmt.Sprintf(logFormat, `Debug`, `%s`)
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

// Emergency writes an emergency value to the logger.
func (l *Logger) Emergency(v ...interface{}) {
	l.write(l.cout, fmt.Sprintf(emergencyFormat, fmt.Sprint(v...)))
}

// Emergencyln writes an emergency value to the logger followed by a new line.
func (l *Logger) Emergencyln(v ...interface{}) {
	l.write(l.cout, fmt.Sprintln(fmt.Sprintf(emergencyFormat, fmt.Sprint(v...))))
}

// Emergencyf writes an emergency value to the logger using the provided format string.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.write(l.cout, fmt.Sprintf(emergencyFormat, fmt.Sprintf(format, v...)))
}

// Alert writes an alert value to the logger.
func (l *Logger) Alert(v ...interface{}) {
	if l.severity >= l5424.AlertLvl {
		l.write(l.cout, fmt.Sprintf(alertFormat, fmt.Sprint(v...)))
	}
}

// Alertln writes an alert value to the logger followed by a new line.
func (l *Logger) Alertln(v ...interface{}) {
	if l.severity >= l5424.AlertLvl {
		l.write(l.cout, fmt.Sprintln(fmt.Sprintf(alertFormat, fmt.Sprint(v...))))
	}
}

// Alertf writes an alert value to the logger using the provided format string.
func (l *Logger) Alertf(format string, v ...interface{}) {
	if l.severity >= l5424.AlertLvl {
		l.write(l.cout, fmt.Sprintf(alertFormat, fmt.Sprintf(format, v...)))
	}
}

// Critical writes a critical value to the logger.
func (l *Logger) Critical(v ...interface{}) {
	if l.severity >= l5424.CritLvl {
		l.write(l.err, fmt.Sprintf(criticalFormat, fmt.Sprint(v...)))
	}
}

// Criticalln writes a critical value to the logger followed by a new line.
func (l *Logger) Criticalln(v ...interface{}) {
	if l.severity >= l5424.CritLvl {
		l.write(l.err, fmt.Sprintln(fmt.Sprintf(criticalFormat, fmt.Sprint(v...))))
	}
}

// Criticalf write a critical value to the logger using the provided format string.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.severity >= l5424.CritLvl {
		l.write(l.err, fmt.Sprintf(criticalFormat, fmt.Sprintf(format, v...)))
	}
}

// Error writes an error value to the logger.
func (l *Logger) Error(v ...interface{}) {
	if l.severity >= l5424.ErrorLvl {
		l.write(l.err, fmt.Sprintf(errorFormat, fmt.Sprint(v...)))
	}
}

// Errorln writes an error value to the logger follower by a new line.
func (l *Logger) Errorln(v ...interface{}) {
	if l.severity >= l5424.ErrorLvl {
		l.write(l.err, fmt.Sprintln(fmt.Sprintf(errorFormat, fmt.Sprint(v...))))
	}
}

// Errorf writes an error value to the logger using the provided format string.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.severity >= l5424.ErrorLvl {
		l.write(l.err, fmt.Sprintf(errorFormat, fmt.Sprintf(format, v...)))
	}
}

// Warn writes a warning value to the logger.
func (l *Logger) Warn(v ...interface{}) {
	if l.severity >= l5424.WarnLvl {
		l.write(l.err, fmt.Sprintf(warnFormat, fmt.Sprint(v...)))
	}
}

// Warnln writes a warning value to the logger followed by a new line.
func (l *Logger) Warnln(v ...interface{}) {
	if l.severity >= l5424.WarnLvl {
		l.write(l.err, fmt.Sprintln(fmt.Sprintf(warnFormat, fmt.Sprint(v...))))
	}
}

// Warnf writes a warning value to the logger using the provided format string.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.severity >= l5424.WarnLvl {
		l.write(l.err, fmt.Sprintf(warnFormat, fmt.Sprintf(format, v...)))
	}
}

// Notice writes a notice value to the logger.
func (l *Logger) Notice(v ...interface{}) {
	if l.severity >= l5424.NoticeLvl {
		l.write(l.out, fmt.Sprintf(noticeFormat, fmt.Sprint(v...)))
	}
}

// Noticeln writes a notice value to the logger followed by a new line.
func (l *Logger) Noticeln(v ...interface{}) {
	if l.severity >= l5424.NoticeLvl {
		l.write(l.out, fmt.Sprintln(fmt.Sprintf(noticeFormat, fmt.Sprint(v...))))
	}
}

// Noticef writes a notice value to the logger using the provided format string.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.severity >= l5424.NoticeLvl {
		l.write(l.out, fmt.Sprintf(noticeFormat, fmt.Sprintf(format, v...)))
	}
}

// Info writes an info value to the logger.
func (l *Logger) Info(v ...interface{}) {
	if l.severity >= l5424.InfoLvl {
		l.write(l.out, fmt.Sprintf(infoFormat, fmt.Sprint(v...)))
	}
}

// Infoln writes an info value to the logger followed by a new line.
func (l *Logger) Infoln(v ...interface{}) {
	if l.severity >= l5424.InfoLvl {
		l.write(l.out, fmt.Sprintln(fmt.Sprintf(infoFormat, fmt.Sprint(v...))))
	}
}

// Infof writes an info value to the logger using the provided format string.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.severity >= l5424.InfoLvl {
		l.write(l.out, fmt.Sprintf(infoFormat, fmt.Sprintf(format, v...)))
	}
}

// Debug writes a debug value to the logger.
func (l *Logger) Debug(v ...interface{}) {
	if l.severity >= l5424.DebugLvl {
		l.write(l.out, fmt.Sprintf(infoFormat, fmt.Sprint(v...)))
	}
}

// Debugln writes a debug value to the logger followed by a new line.
func (l *Logger) Debugln(v ...interface{}) {
	if l.severity >= l5424.DebugLvl {
		l.write(l.out, fmt.Sprintln(fmt.Sprintf(infoFormat, fmt.Sprint(v...))))
	}
}

// Debugf writes a debug value to the logger using the provided format string.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.severity >= l5424.DebugLvl {
		l.write(l.out, fmt.Sprintf(infoFormat, fmt.Sprintf(format, v...)))
	}
}
