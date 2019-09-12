package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type redisConfig struct {
	hostIP      string
	hostPort    int
	hostSocket  string
	repeat      int
	interval    int
	dbNum       int
	interactive int
	shutdown    int

	monitorMode              int
	pubSubModel              int
	latencyMode              int
	latencyHistory           int
	clusterMode              int
	clusterReissueCommand    int
	slaveMode                int
	pipeMode                 int
	pipeTimeout              int
	getRdbMode               int
	statMode                 int
	scanMode                 int
	intrinsicLatencyMode     int
	intrinsicLatencyDuration int

	pattern     string
	rdbFilename string
	bigKeys     int
	stdinArg    int
	auth        string
	output      int
	mbDelim     string
	eval        string
}

const (
	RedisCliDefaultPipeTimeout = 30
)

func main() {
	config := initConfig()
	initHelp()
}

func initConfig() *redisConfig {
	config := &redisConfig{
		hostIP:      "127.0.0.1",
		hostPort:    6379,
		hostSocket:  "",
		repeat:      1,
		interval:    0,
		dbNum:       0,
		interactive: 0,
		shutdown:    0,

		monitorMode:          0,
		pubSubModel:          0,
		latencyMode:          0,
		clusterMode:          0,
		slaveMode:            0,
		getRdbMode:           0,
		pipeMode:             0,
		pipeTimeout:          RedisCliDefaultPipeTimeout,
		statMode:             0,
		scanMode:             0,
		intrinsicLatencyMode: 0,

		pattern:     "",
		rdbFilename: "",
		bigKeys:     0,
		stdinArg:    0,
		auth:        "",
		mbDelim:     "\n",
		eval:        "",
	}
	// TODO: fake tty supports
	return config
}

func cliVersion() string {
	version := bytes.Buffer{}
	version.WriteString(GetRedisVersion())
	if ret, err := strconv.ParseInt(GetRedisSHA1(), 16, 64); err != nil {
		panic(err)
	} else if ret != 0 {
		version.WriteString(
			fmt.Sprintf("(git:%s", GetRedisSHA1()))
		if ret, err = strconv.ParseInt(GetRedisGitDirty(), 10, 64); err != nil {
			panic(err)
		} else if ret != 0 {
			version.WriteString("-dirty")
		}
		version.WriteRune(')')
	}
	return version.String()
}

//func autoCompletion(buf *bytes.Buffer, lc *lineNoiseCompletions) {
//
//}
