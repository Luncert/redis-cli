package main

import (
	"strings"
	"testing"
)

var keys = []string{"QQA", "QQAS", "QQASD", "QQBSD", "QQBS", "QQB"}

func Test_helpSearchTree(t *testing.T) {
	h := NewCommandHelp()

	for _, key := range keys {
		c := &CommandHelpEntry{name: key}
		if err := h.insert(c); err != nil {
			t.Error(err)
		}
	}
	for _, key := range keys {
		node := h.root
		for _, r := range key {
			if node = node.matchChildren(uint8(r)); node == nil {
				t.Error("invalid result")
				break
			}
		}
	}
	for _, key := range []string{"qqa", "qqb"} {
		if entries, ok := h.Search(key); !ok || len(entries) != 3 {
			t.Error("invalid result")
		} else {
			for i, entry := range entries {
				if len(entry.name) != i+3 {
					t.Error("invalid result")
				}
			}
		}
	}
}

func Test_splitParams(t *testing.T) {
	testParams := []string{"key", "min", "max", "[WITHSCORES]", "[LIMIT offset count]"}
	raw := strings.Join(testParams, " ")
	params := splitParams(raw)
	for idx, param := range params {
		if param != testParams[idx] {
			t.Errorf("expect: %s, got: %s", testParams[idx], param)
		}
	}
}
