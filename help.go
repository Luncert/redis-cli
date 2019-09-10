package main

type helpEntry struct {
	helpType int
	argc     int
	argv     string
	full     string
}

type commandHelp struct {
	name    string
	params  string
	summary string
	group   int
	since   string
}

var commandHelps []commandHelp
