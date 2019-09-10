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
	fmt.Printf("\r\n  \x1b[1m%s\x1b[0m \x1b[90m%s\x1b[0m\r\n", "help->name", "help->params")
	fmt.Printf("  \x1b[33msummary:\x1b[0m %s\r\n", "help->summary")
	fmt.Printf("  \x1b[33msince:\x1b[0m %s\r\n", "help-since")
	config := initConfig()
	fmt.Println(config)
}

func initConfig() *config {
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
	return config
}

func slaveMode() {

}

func redisGitSHA1() string {
	return ""
}

func redisGitDirty() string {
	return ""
}
