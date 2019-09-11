package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(RedisGitSha1, RedisGitDirty)
}
