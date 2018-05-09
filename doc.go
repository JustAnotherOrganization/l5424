// Package l5424 is a collection of RFC5424 compliant logging interfaces and loggers.
//
// The point of this repo is to show several possible ways of implementing RFC5424 compliant loggers.
//
// The Go Language is extremely fragmented in how it addresses logging. The goal of this project is to
// try to start discussion regarding some form of standardized logging in Go.
//
// During the course of this project it's come to my attention that a #logging channel exists in the current
// Gopher's slack for the purpose of standardizing logging in Go (which was a partial goal for this project,
// yay!).
//
// Having read through the different proposals outlined in [Proposal: Add Logging interface to standard library](https://docs.google.com/document/d/1oTjtY49y8iSxmM9YBaz2NrZIlaXtQsq3nQMd-E0HtwM/edit#) I've come to the conclusion that my original implementations of RFC5424 loggers being based purely on the RFC were naive (at best) and have abandoned my original implementations. They can be found in the [legacy](https://github.com/JustAnotherOrganization/l5424/tree/legacy) branch of this repo but will not be directly maintained (in the very unlikely event someone chooses to use them I will welcome any bug fixes (if there are bugs) but that is all that will be accepted for the legacy loggers).
//
// Only one interface is provided and that is the one described in the above mentioned proposals.
//
// If you're reading this you are encouraged to play with the interface and loggers as well as read
// the above proposals and the existing RFC5424 standard.
package l5424 // import "github.com/JustAnotherOrganization/l5424"
