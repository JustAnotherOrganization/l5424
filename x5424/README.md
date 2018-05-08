# x5424

[![GoDoc][docs_badge]][docs]

---

An RFC5424 compliant logger. This is an example implementation of the l5424 log interface.

x5424 writes to both an out and an err writer (default: os.Stdout and os.Stderr respectively) where values with a severity level of EmergencyLvl  to AlertLvl are written to err only and values of NoticeLvl or above are written to out only.

[docs]: https://godoc.org/github.com/JustAnotherOrganization/l5424/x5424
[docs_badge]: https://godoc.org/github.com/JustAnotherOrganization/l5424/x5424?status.svg