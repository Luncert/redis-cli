package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	RedisGitSha1  string
	RedisGitDirty int64
	RedisBuildID  string
)

var releaseInfoFilePath = filepath.Join("data", "release.info")

func init() {
	redisGitSha1, redisGitDirty, redisBuildID := loadReleaseInfo()

	setReleaseInfo()

	// if no change happens on this repository, then we use the RedisBuildID read from file.
	// else we use new generated RedisBuildID, and invoke storeReleaseInfo to save release information.
	if redisGitSha1 == RedisGitSha1 &&
		redisGitDirty == RedisGitDirty {
		RedisBuildID = redisBuildID
	} else {
		storeReleaseInfo()
	}
}

func loadReleaseInfo() (redisGitSha1 string, redisGitDirty int64, redisBuildID string) {
	if _, err := os.Stat(releaseInfoFilePath); err != nil {
		// file not exists
		return "00000000", 1, "unknown"
	}

	data, err := ioutil.ReadFile(releaseInfoFilePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "=")
		switch items[0] {
		case "RedisGitSha1":
			redisGitSha1 = items[1]
		case "RedisGitDirty":
			redisGitDirty, err = strconv.ParseInt(items[1], 10, 64)
			if err != nil {
				panic(err)
			}
		case "RedisBuildID":
			redisBuildID = items[1]
		}
	}

	return
}

func setReleaseInfo() {
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
	if ret, err := ExecCmd("git", "diff", "--numstat"); err != nil {
		panic(err)
	} else {
		reg := regexp.MustCompile(`(\d+)[ \t]+(\d+)[ \t]+.+`)
		match := reg.FindStringSubmatch(ret)
		var count int64
		for idx, item := range match {
			if idx%3 == 0 {
				continue
			}
			if i, err := strconv.ParseInt(item, 10, 64); err != nil {
				panic(err)
			} else {
				count = count + i
			}
		}
		RedisGitDirty = count
	}

	// set RedisBuildID
	if ret, err := ExecCmd("hostname"); err != nil {
		panic(err)
	} else {
		now := time.Now()
		RedisBuildID = fmt.Sprintf("%s-%s",
			strings.Trim(ret, "\r\n"),
			strconv.FormatInt(now.Unix(), 10))
	}
}

func storeReleaseInfo() {
	buf := bytes.Buffer{}
	buf.WriteString("RedisGitSha1=")
	buf.WriteString(RedisGitSha1)
	buf.WriteRune('\n')
	buf.WriteString("RedisGitDirty=")
	buf.WriteString(strconv.FormatInt(RedisGitDirty, 10))
	buf.WriteRune('\n')
	buf.WriteString("RedisBuildID=")
	buf.WriteString(RedisBuildID)
	buf.WriteRune('\n')

	if err := ioutil.WriteFile(releaseInfoFilePath, buf.Bytes(), 600); err != nil {
		panic(err)
	}
}
