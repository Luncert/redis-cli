package main

import "fmt"

type config struct {
	hostIP      string
	hostPort    int
	hostSocket  string
	repeat      int
	interval    int
	dbNum       int
	interactive int
	shutdown    int
	monitorMode bool
	pubSubModel bool
	latencyMode bool

	auth string
	eval string
	// output int
}

func main() {
	config := &config{
		hostIP:      "127.0.0.1",
		hostPort:    6379,
		hostSocket:  "",
		repeat:      1,
		interval:    0,
		dbNum:       0,
		interactive: 0,
		shutdown:    0,
		monitorMode: false,
		pubSubModel: false,
		latencyMode: false,
		// ...
		auth: "",
		eval: "",
	}
	fmt.Println(config)
}

func usage() {

}

func slaveMode() {

}

func redisGitSHA1() string {
	return ""
}

func redisGitDirty() string {
	return ""
}
