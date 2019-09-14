package main

import (
	"bytes"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

var LivePrefixState struct {
	ServerAddr string
	DbIdx      int
	LivePrefix string
}

var commandHelp = NewCommandHelp()

func fatalF(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	fmt.Print('\n')
	os.Exit(1)
}

func main() {
	LivePrefixState.ServerAddr = "localhost:6379"

	color.Green("redis-cli v%s\n", GetClientVersion())

	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("redis-cli"),
		prompt.OptionPrefix("> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionPrefixTextColor(prompt.DarkGreen),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkBlue),
	)
	p.Run()
}

func changeLivePrefix() (string, bool) {
	buf := bytes.Buffer{}
	if LivePrefixState.ServerAddr != "" {
		buf.WriteString(LivePrefixState.ServerAddr)
	}
	if LivePrefixState.DbIdx != 0 {
		buf.WriteRune('[')
		buf.WriteString(strconv.FormatInt(int64(LivePrefixState.DbIdx), 10))
		buf.WriteRune(']')
	}
	buf.WriteString("> ")
	return buf.String(), true
}

func executor(in string) {
	in = strings.TrimSpace(in)
}

func completer(d prompt.Document) (suggests []prompt.Suggest) {
	line := d.CurrentLine()
	if line != "" {
		items := strings.Split(line, " ")
		if len(items) > 1 {
			if entry, ok := commandHelp.PreciseSearch(items[0]); ok {
				paramIdx := len(items) - 2
				if paramIdx < len(entry.params) {
					suggests = append(suggests, prompt.Suggest{
						Text:        entry.params[paramIdx],
						Description: "",
					})
				}
			}
		} else if entries, ok := commandHelp.Search(items[0]); ok {
			for _, entry := range entries {
				suggests = append(suggests, prompt.Suggest{
					Text:        entry.name,
					Description: entry.summary,
				})
			}
		}
	}
	return
}
