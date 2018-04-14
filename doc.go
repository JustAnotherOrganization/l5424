// Package l5424 is an RFC5424 log interface. It does not provide a logger directly, only the leveled interface.
//
// This can be used to build RCF5424 compliant logging solutions into an application which are
// unaware of their background logger.
//
// l5424 doesn't include any dependencies and is small enough that if you intend on writing your
// own logger (instead of using the provided wrapper) you can simply copy/paste the main interface
// into your project.
package l5424
