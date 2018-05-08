package min5424

import (
	"fmt"
	"io"

	"github.com/JustAnotherOrganization/l5424"
)

// Static type checking.
var (
	_ l5424.MinLog = &Logger{}
)

const (
	defaultFormat = "%s : %s"
)

var (
	facilityFormat = fmt.Sprintf("%s - %s", "%s", defaultFormat)
)

type (
	// Logger implements a l5424.MinLogger.
	Logger struct {
		format      string
		severity    l5424.SeverityLvl
		facility    l5424.FacilityLvl
		facDisabled bool
	}
)

func new(facility, severity, format string) (*Logger, error) {
	var logger Logger

	_facility, err := l5424.FacilityLvlString(facility)
	if err != nil {
		logger.facDisabled = true
		if format == "" {
			format = defaultFormat
		}
	} else {
		logger.facility = _facility
		if format == "" {
			format = facilityFormat
		}
	}
	logger.format = format

	_severity, err := l5424.SeverityLvlString(severity)
	if err != nil {
		return nil, err
	}

	logger.severity = _severity

	return &logger, nil
}

// New returns a new Logger:
// facility is the string value of a l5424.FacilityLvl
// severity  is the string value of a l5424.SeverityLvl
func New(facility, severity string) (*Logger, error) {
	return new(facility, severity, "")
}

// NewWithFormat returns a new logger using the provided format,
// see New for more information.
func NewWithFormat(facility, severity, format string) (*Logger, error) {
	return new(facility, severity, format)
}

func (l *Logger) write(w io.Writer, lvl l5424.SeverityLvl, v string) {
	if lvl < l5424.Disabled {
		fmt.Fprint(w, v)
	}
}

// Print writes to the logger using the provided severity level.
func (l *Logger) Print(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	if l.facDisabled {
		l.write(w, lvl, fmt.Sprintf(l.format, lvl.String(), fmt.Sprint(v...)))
	} else {
		l.write(w, lvl, fmt.Sprintf(l.format, l.facility.String(), lvl.String(), fmt.Sprint(v...)))
	}
}

// Println writes to the logger using the provided severity level, followed by a new line.
func (l *Logger) Println(w io.Writer, lvl l5424.SeverityLvl, v ...interface{}) {
	if l.facDisabled {
		l.write(w, lvl, fmt.Sprintln(fmt.Sprintf(l.format, lvl.String(), fmt.Sprint(v...))))
	} else {
		l.write(w, lvl, fmt.Sprintln(fmt.Sprintf(l.format, l.facility.String(), lvl.String(), fmt.Sprint(v...))))
	}
}

// Printf writes to the logger using the provided severity level and format.
func (l *Logger) Printf(w io.Writer, lvl l5424.SeverityLvl, format string, v ...interface{}) {
	if l.facDisabled {
		l.write(w, lvl, fmt.Sprintf(l.format, lvl.String(), fmt.Sprintf(format, v...)))
	} else {
		l.write(w, lvl, fmt.Sprintf(l.format, l.facility.String(), lvl.String(), fmt.Sprintf(format, v...)))
	}
}
