package main

import "strings"

var (
	RedisGitSha1  = "00000000"
	RedisGitDirty = 0
)

func init() {
	// set RedisGitSha1
	if ret, err := ExecCmd("git", "show-ref", "--head", "--hash=8"); err != nil {
		panic(err)
	} else {
		items := strings.Split(ret, "\n")
		if len(items) > 0 {
			RedisGitSha1 = items[0]
		}
	}

	// set RedisGitDirty
	if ret, err := ExecCmd("git", "diff", "--no-ext-diff"); err != nil {
		panic(err)
	} else {
		items := strings.Split(ret, "\n")
		RedisGitDirty = len(items)
	}
}
