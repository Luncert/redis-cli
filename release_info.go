package main

import (
	"bufio"
	"bytes"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	redisVersion = "2.9.11"

	redisGitSha1  string
	redisGitDirty string
	redisBuildID  string
)

var releaseInfoFilePath = filepath.Join("data", "release.info")

func init() {
	gitSha1, gitDirty, buildID := loadReleaseInfo()

	setReleaseInfo()

	// if no change happens on this repository, then we use the RedisBuildID read from file.
	// else we use new generated RedisBuildID, and invoke storeReleaseInfo to save release information.
	if redisGitSha1 == gitSha1 &&
		redisGitDirty == gitDirty {
		redisBuildID = buildID
	} else {
		storeReleaseInfo()
	}
}

func loadReleaseInfo() (redisGitSha1 string, redisGitDirty string, redisBuildID string) {
	if _, err := os.Stat(releaseInfoFilePath); err != nil {
		// file not exists
		return "00000000", "1", "unknown"
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
			redisGitDirty = items[1]
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
			redisGitSha1 = items[0]
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
		redisGitDirty = strconv.FormatInt(count, 10)
	}

	// set RedisBuildID
	if ret, err := ExecCmd("hostname"); err != nil {
		panic(err)
	} else {
		now := time.Now()
		redisBuildID = fmt.Sprintf("%s-%s",
			strings.Trim(ret, "\r\n"),
			strconv.FormatInt(now.Unix(), 10))
	}
}

func storeReleaseInfo() {
	buf := bytes.Buffer{}
	buf.WriteString("RedisGitSha1=")
	buf.WriteString(redisGitSha1)
	buf.WriteRune('\n')
	buf.WriteString("RedisGitDirty=")
	buf.WriteString(redisGitDirty)
	buf.WriteRune('\n')
	buf.WriteString("RedisBuildID=")
	buf.WriteString(redisBuildID)
	buf.WriteRune('\n')

	if err := ioutil.WriteFile(releaseInfoFilePath, buf.Bytes(), 600); err != nil {
		panic(err)
	}
}

func GetRedisVersion() string {
	return redisVersion
}

func GetRedisSHA1() string {
	return redisGitSha1
}

func GetRedisGitDirty() string {
	return redisGitDirty
}

func GetRedisBuildID() uint64 {
	buf := bytes.Buffer{}
	buf.WriteString(redisVersion)
	buf.WriteString(redisBuildID)
	buf.WriteString(redisGitDirty)
	buf.WriteString(redisGitSha1)

	table := crc64.MakeTable(crc64.ISO)
	return crc64.Checksum([]byte(buf.String()), table)
}
