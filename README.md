# l5424

[![GoDoc][docs_badge]][docs]

---

An RFC5424 logger.

l5424 writes to both an out and an err writer (default: os.Stdout and os.Stderr respectively) where values with a severity level of EmergencyLvl are written to both while values between WarnLvl and AlertLvl are written to err only and values of NoticeLvl or above are written to out only.

[docs]: https://godoc.org/github.com/JustAnotherOrganization/l5424
[docs_badge]: https://godoc.org/github.com/JustAnotherOrganization/l5424?status.svg