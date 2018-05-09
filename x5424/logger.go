package x5424

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/JustAnotherOrganization/l5424"
)

// Logger implements the l5424.Logger (an implementation of the proposed standardized logger).
type (
	Logger struct {
		w   io.Writer
		lvl l5424.SeverityLvl
	}

	checkNext struct {
		label string
		prev  interface{}
		f     func(v interface{}) (uint, bool)
	}
)

// New returns a new Logger, level is a l5424 severity level string (unrecognized severity strings result
// in a level of l5424.NoticeLvl). If w is nil os.Stdout is assumed.
func New(level string, w io.Writer) *Logger {
	l := Logger{
		lvl: l5424.NoticeLvl,
		w:   os.Stdout,
	}

	if lvl, err := l5424.SeverityLvlString(level); err == nil {
		l.lvl = lvl
	}

	if w != nil {
		l.w = w
	}

	return &l
}

const (
	// Severity can be passed as a key for a severity level. The accompanied severity level
	// must be a string or a pointer assosciated with a SeverityLvl enum found in l5424.
	Severity = "severity"
	// Facility can be passed as a key for a facility level. The accompanied facility level
	// must be a string or a pointer assosciated with a FacilityLvl enum found in l5424.
	Facility = "facility"
)

// Log is a flexible log method described in the standarized logger proposals.
// v can be key-pairs of alternating key, value (limited to recognized key, currently;
// all other data is appended to the end of the log entry in order).
func (l *Logger) Log(v ...interface{}) error {
	if len(v) == 0 {
		return errors.New("log value cannot be nil")
	}

	// TODO: offer alternate behavior that assumes all data is key/value

	var (
		next                     *checkNext
		facilitySet, severitySet bool
		m                        = make(map[string]uint)
		final                    []interface{}
	)

	for i, _v := range v {
		// Don't bother looking through all the keys if all recognized keys
		// have already been found.
		if facilitySet && severitySet {
			final = append(final, v[i:]...)
			break
		}

		if next != nil {
			if u, ok := next.f(_v); ok {
				// This should not happen...
				if next.label != Severity && next.label != Facility {
					final = append(final, next.prev, _v)
				} else {
					if next.label == Severity {
						// The severity level for this message is higher than the log level, do nothing.
						if l.lvl < l5424.SeverityLvl(u) {
							return nil
						}

						severitySet = true
					}
					if next.label == Facility {
						facilitySet = true
					}

					m[next.label] = u
				}
			}

			next = nil
			continue
		}

		switch _v.(type) {
		case string:
			switch _v.(string) {
			case Severity:
				next = &checkNext{
					label: Severity,
					prev:  v,
					f:     checkSeverity,
				}
			case Facility:
				next = &checkNext{
					label: Facility,
					prev:  v,
					f:     checkFacility,
				}
			default:
				final = append(final, _v)
			}
		}
	}

	if next != nil {
		final = append(final, next.prev)
		next = nil // Kind of not needed now lol.
	}

	return l.log(m, final...)
}

func (l *Logger) log(m map[string]uint, v ...interface{}) error {
	var format []string
	if facility, ok := m[Facility]; ok {
		format = append(format, l5424.FacilityLvl(facility).String())
	}
	if severity, ok := m[Severity]; ok {
		format = append(format, l5424.SeverityLvl(severity).String())
	}

	// TODO: allow for some other format delimeters?
	// json formatting as an option?
	return l._log(strings.Join(format, " - "), v...)
}

func (l *Logger) _log(format string, v ...interface{}) error {
	_, err := fmt.Fprint(l.w, fmt.Sprintf("%s: %s", format, fmt.Sprint(v...)))
	return err
}

func checkSeverity(v interface{}) (uint, bool) {
	switch v.(type) {
	case string:
		lvl, err := l5424.SeverityLvlString(v.(string))
		if err == nil {
			return uint(lvl), true
		}
		return uint(lvl), true
	case l5424.SeverityLvl:
		return uint(v.(l5424.SeverityLvl)), true
	default:
		return 0, false
	}
}

func checkFacility(v interface{}) (uint, bool) {
	switch v.(type) {
	case string:
		lvl, err := l5424.FacilityLvlString(v.(string))
		if err == nil {
			return uint(lvl), true
		}
		return uint(lvl), true
	case l5424.FacilityLvl:
		return uint(v.(l5424.FacilityLvl)), true
	default:
		return 0, false
	}
}
