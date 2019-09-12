package main

//
//const (
//	lineNoiseDefaultHistoryMaxLen = 100
//	lineNoiseMaxLine = 4096
//)
//
//var (
//	unsupportedTerms = []string{"dumb", "cons25", "emacs", ""}
//
//	rawMode = false
//	mlMode = false
//	atExitRegistered = false
//	historyMaxLen = lineNoiseDefaultHistoryMaxLen
//	historyLen = 0
//	history []string
//)
//
//type LineNoiseCompletions struct {
//	len int
//	cvec []string
//}
//
//type LineNoiseCompletionCallback func(string, *LineNoiseCompletions)
//
//var completionCallback LineNoiseCompletionCallback = nil
//
//// lineNoiseState represents the state during line editing.
//// We pass this state to functions implementing specific editing functionalities.
//type lineNoiseState struct {
//	ifd int
//	ofd int
//	buf string
//	prompt string
//	pLen int
//	pos int
//	oldPos int
//	len int
//	cols int
//	maxRows int
//	historyIndex int
//}
//
//// key actions
//const (
//	keyNull = iota
//	keyCtrlA
//	keyCtrlB
//	keyCtrlC
//	keyCtrlD
//	keyCtrlE
//	keyCtrlF
//	keyCtrlH = 8
//	keyTab = 9
//	keyCtrlK = 11
//	keyCtrlL = 12
//	keyEnter = 13
//	keyCtrlN = 14
//	keyCtrlP = 16
//	keyCtrlT = 20
//	keyCtrlU = 21
//	keyCtrlW = 23
//	keyEsc = 27
//	keyBackspace = 127
//)
//
//func lineNoiseSetMultiLine(ml bool) {
//	mlMode = ml
//}
//
//func isSupportedTerm()
