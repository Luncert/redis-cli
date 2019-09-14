package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"sort"
)

var commandHelpFilePath = filepath.Join("data", "command_help.json")

type CommandHelp struct {
	root *helpSearchTreeNode
}

func NewCommandHelp() *CommandHelp {
	c := &CommandHelp{
		root: &helpSearchTreeNode{
			key:   0,
			refs:  map[uint8]*helpSearchTreeNode{},
			entry: nil,
		},
	}
	loadCommandHelp(c)
	return c
}

func loadCommandHelp(c *CommandHelp) {
	data, err := ioutil.ReadFile(commandHelpFilePath)
	if err != nil {
		fatalF("Failed to load command_help.json: %v", err)
	}

	var cmdHelps []interface{}
	err = json.Unmarshal(data, &cmdHelps)
	if err != nil {
		panic(err)
	}

	for _, item := range cmdHelps {
		cmdHelpRaw := item.(map[string]interface{})

		// build params

		cmdHelp := &CommandHelpEntry{
			name:    cmdHelpRaw["name"].(string),
			summary: cmdHelpRaw["summary"].(string),
			params:  splitParams(cmdHelpRaw["params"].(string)),
			//group:   int(cmdHelpRaw["group"].(float64)),
			//since:   cmdHelpRaw["since"].(string),
		}
		if err = c.insert(cmdHelp); err != nil {
			fatalF("Failed to build helpSearchTree: %v", err)
		}
	}
}

func splitParams(raw string) (params []string) {
	inBracket := false
	buf := bytes.Buffer{}
	for _, r := range raw {
		switch r {
		case '[':
			inBracket = true
			buf.WriteRune(r)
		case ']':
			inBracket = false
			buf.WriteRune(r)
		case ' ':
			if !inBracket {
				params = append(params, buf.String())
				buf.Reset()
			} else {
				buf.WriteRune(r)
			}
		default:
			buf.WriteRune(r)
		}
	}
	lastParam := buf.String()
	if len(lastParam) > 0 {
		params = append(params, lastParam)
	}
	return
}

func (c *CommandHelp) insert(e *CommandHelpEntry) error {
	i := 0
	sz := len(e.name)

	if sz == 0 {
		return errors.New("invalid CommandHelpEntry, empty key")
	}

	preNode := c.root
	node := c.root
	for i < sz {
		preNode = node
		node = node.matchChildren(e.name[i])
		if node != nil {
			i++
		} else {
			for ; i < sz; i++ {
				node = &helpSearchTreeNode{
					key:   e.name[i],
					refs:  map[uint8]*helpSearchTreeNode{},
					entry: nil,
				}
				preNode.refs[node.key] = node
				preNode = node
			}
			preNode.entry = e
			return nil
		}
	}

	if node.entry != nil {
		return errors.New("invalid CommandHelpEntry, key existed")
	} else {
		node.entry = e
	}
	return nil
}

func (c *CommandHelp) PreciseSearch(key string) (*CommandHelpEntry, bool) {
	node := c.searchNode(key)
	if node == nil || node.entry == nil {
		return nil, false
	}
	return node.entry, true
}

func (c *CommandHelp) Search(key string) ([]*CommandHelpEntry, bool) {
	node := c.searchNode(key)
	if node == nil {
		return nil, false
	}
	entries := node.getDescendants()
	if node.entry != nil {
		entries = append(node.getDescendants(), node.entry)
	}
	// sort
	sort.Sort(CommandHelpEntries(entries))
	return entries, true
}

func (c *CommandHelp) searchNode(key string) *helpSearchTreeNode {
	node := c.root
	for _, r := range key {
		// to lower
		if r >= 97 && r <= 122 {
			r -= 32
		}
		if node = node.matchChildren(uint8(r)); node == nil {
			return nil
		}
	}
	return node
}

type CommandHelpEntry struct {
	name    string
	summary string
	params  []string
}

func (c *CommandHelpEntry) marshall() (ret int64) {
	for i, r := range c.name {
		ret = ret + (int64(r) << uint(i))
	}
	return
}

type CommandHelpEntries []*CommandHelpEntry

func (c CommandHelpEntries) Len() int {
	return len(c)
}

func (c CommandHelpEntries) Less(i, j int) bool {
	return c[i].marshall() < c[j].marshall()
}

func (c CommandHelpEntries) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type helpSearchTreeNode struct {
	key   uint8
	refs  map[uint8]*helpSearchTreeNode
	entry *CommandHelpEntry
}

func (h *helpSearchTreeNode) match(key uint8) bool {
	return h.key == key
}

func (h *helpSearchTreeNode) matchChildren(key uint8) *helpSearchTreeNode {
	for _, ref := range h.refs {
		if ref.match(key) {
			return ref
		}
	}
	return nil
}

func (h *helpSearchTreeNode) getDescendants() (entries []*CommandHelpEntry) {
	for _, node := range h.refs {
		if node.entry != nil {
			entries = append(entries, node.entry)
		}
		entries = append(entries, node.getDescendants()...)
	}
	return
}
