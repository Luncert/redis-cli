package main

import (
	"encoding/json"
	"io/ioutil"
)

const (
	cliHelpCommand = iota
	cliHelpGroup
)

type helpEntry struct {
	helpType int
	argc     int
	argv     string
	full     string
	org      *commandHelp
}

var helpEntries []helpEntry

// command groups
const (
	cg_generic int = iota
	cg_string
	cg_list
	cg_set
	cg_sorted_set
	cg_hash
	cg_pubsub
	cg_transactions
	cg_connection
	cg_server
	cg_scripting
	cg_hyperloglog
)

type commandHelp struct {
	name    string
	params  string
	summary string
	group   int
	since   string
}

var commandHelps []commandHelp

func init() {
	data, err := ioutil.ReadFile("./command_help.json")
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

func cliVersion() string {
	return ""
}
