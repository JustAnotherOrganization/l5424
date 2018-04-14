# l5424

[![GoDoc][docs_badge]][docs]

---

An RFC5424 log interface. l5424 does not provide a logger directly, only the leveled interface.

This can be used to build RCF5424 compliant logging solutions into an application which are unaware of their background logger.

l5424 doesn't include any dependencies and is small enough that if you intend on writing your own logger (instead of using the provided [wrapper](https://github.com/JustAnotherOrganization/l5424/tree/master/dw)) you can simply copy/paste the main interface into your project.

[docs]: https://godoc.org/github.com/JustAnotherOrganization/l5424
[docs_badge]: https://godoc.org/github.com/JustAnotherOrganization/l5424?status.svg