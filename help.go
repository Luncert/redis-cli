package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	cliHelpCommand = iota
	cliHelpGroup
)

type helpEntry struct {
	helpType int
	argc     int
	argv     []string
	full     string
	org      *commandHelp
}

var helpEntries []helpEntry

// command groups
const (
	cgGeneric int = iota
	cgString
	cgList
	cgSet
	cgSortedSet
	cgHash
	cgPubSub
	cgTransactions
	cgConnection
	cgServer
	cgScripting
	cgHyperLogLog
	cgInvalid
)

var commandGroups = []string{
	cgGeneric:      "generic",
	cgString:       "string",
	cgList:         "list",
	cgSet:          "set",
	cgSortedSet:    "sorted_set",
	cgHash:         "hash",
	cgPubSub:       "pubsub",
	cgTransactions: "transactions",
	cgConnection:   "connection",
	cgServer:       "server",
	cgScripting:    "scripting",
	cgHyperLogLog:  "hyperloglog",
}

type commandHelp struct {
	name    string
	params  string
	summary string
	group   int
	since   string
}

var commandHelps []commandHelp
var commandHelpFilePath = filepath.Join("data", "command_help.json")

func initHelp() {
	initCommandHelps()

	commandsLen := len(commandHelps)
	groupsLen := cgHyperLogLog - cgGeneric + 1

	for i := 0; i < groupsLen; i++ {
		argv := []string{fmt.Sprintf("@%s", commandGroups[i])}
		tmp := helpEntry{
			helpType: cliHelpGroup,
			argc:     1,
			argv:     argv,
			full:     argv[0],
			org:      nil,
		}
		helpEntries = append(helpEntries, tmp)
	}

	for i := 0; i < commandsLen; i++ {
		argv := strings.Split(commandHelps[i].name, " ")
		tmp := helpEntry{
			helpType: cliHelpCommand,
			argc:     len(argv),
			argv:     argv,
			full:     commandHelps[i].name,
			org:      &commandHelps[i],
		}
		helpEntries = append(helpEntries, tmp)
	}
}

// load command_help.json
func initCommandHelps() {
	data, err := ioutil.ReadFile(commandHelpFilePath)
	if err != nil {
		// TODO: use the log library in Redishadow
		panic(err)
	}

	var cmdHelps []interface{}
	err = json.Unmarshal(data, &cmdHelps)
	if err != nil {
		panic(err)
	}

	for _, item := range cmdHelps {
		cmdHelpRaw := item.(map[string]interface{})
		cmdHelp := commandHelp{
			name:    cmdHelpRaw["name"].(string),
			params:  cmdHelpRaw["params"].(string),
			summary: cmdHelpRaw["summary"].(string),
			group:   int(cmdHelpRaw["group"].(float64)),
			since:   cmdHelpRaw["since"].(string),
		}
		commandHelps = append(commandHelps, cmdHelp)
	}
}

// outputHelp outputs all command help, filtering by group or command name
func outputHelp(argc int, argv ...string) {
	if argc == 0 {
		outputGenericHelp()
	} else {
		if argc < 0 {
			panic(errors.New("invalid argument"))
		}

		group := cgInvalid
		if argv[0][0] == '@' {
			cmdHelpsLen := len(commandHelps)
			for i := 0; i < cmdHelpsLen; i++ {
				if stringCaseCompare(argv[1][1:], commandGroups[i]) {
					group = i
					break
				}
			}
		}

		for i := 0; i < len(helpEntries); i++ {
			entry := helpEntries[i]
			if entry.helpType != cliHelpCommand {
				continue
			}

			help := entry.org
			if group == cgInvalid {
				if argc == entry.argc {
					j := 0
					for ; j < argc; j++ {
						if stringCaseCompare(argv[j], entry.argv[j]) {
							break
						}
					}
					if j == argc {
						outputCommandHelp(help, true)
					}
				}
			} else if group == help.group {
				outputCommandHelp(help, false)
			}
		}
		fmt.Printf("\r\n")
	}
}

func stringCaseCompare(s, t string) (equal bool) {
	if len(s) != len(t) {
		equal = false
	} else {
		equal = true
		var sr, tr uint8
		for i := 0; i < len(s); i++ {
			sr = charToLower(s[i])
			tr = charToLower(t[i])
			if sr != tr {
				equal = false
				break
			}
		}
	}
	return
}

func charToLower(c uint8) uint8 {
	if c >= 'a' && c <= 'z' {
		c -= 32
	}
	return c
}

func outputCommandHelp(help *commandHelp, printGroup bool) {
	fmt.Printf("\r\n  \x1b[1m%s\x1b[0m \x1b[90m%s\x1b[0m\r\n", help.name, help.params)
	fmt.Printf("  \x1b[33msummary:\x1b[0m %s\r\n", help.summary)
	fmt.Printf("  \x1b[33msince:\x1b[0m %s\r\n", help.since)
	if printGroup {
		fmt.Printf("  \x1b[33mgroup:\x1b[0m %s\r\n", commandGroups[help.group])
	}
}

func outputGenericHelp() {
	version := cliVersion()
	fmt.Printf("redis-cli %s\r\n"+
		"Type: \"help @<group>\" to get a list of commands in <group>\r\n"+
		"      \"help <command>\" for help on <command>\r\n"+
		"      \"help <tab>\" to get a list of possible help topics\r\n"+
		"      \"quit\" to exit\r\n", version)
}
