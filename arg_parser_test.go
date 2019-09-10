package main

import (
	"testing"
)

type testArgObject struct {
	Name  string `arg:"name" alias:"n"`
	Allow bool   `arg:"allow" alias:"al"`
	Age   int    `arg:"age" alias:"ag"`
}

func TestParser(t *testing.T) {
	args := []string{"-n", "Joy", "-al", "true", "-ag", "10"}
	parser := newArgParser()
	parser.parse(args)
	if v := parser.getArg("name", "n"); v != args[1] {
		t.Errorf("expect: %s, got: %s", args[1], v)
	}
	if v := parser.getArg("allow", "al"); v != args[3] {
		t.Errorf("expect: %s, got: %s", args[3], v)
	}
	if v := parser.getArg("age", "ag"); v != args[5] {
		t.Errorf("expect: %s, got: %s", args[5], v)
	}
}

func TestParseArgs(t *testing.T) {
	tmp := &testArgObject{}
	if err := ParseArgs(tmp); err != nil {
		t.Error(err)
	}
}
