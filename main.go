package main

import (
	"fmt"
	"github.com/saromanov/automatic/automatic"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app       = kingpin.New("automatic", "A command-line chat application.")
	debug     = app.Flag("debug", "Enable debug mode.").Bool()
	deploy    = app.Command("deploy", "Run deploy environment")
	test      = app.Command("test", "Run test environment")
	exec      = app.Command("exec", "Register a new user.")
	newDeploy = deploy.Arg("new", "Nickname for user.").String()
)

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case deploy.FullCommand():
		cfg := automatic.Automatic{}
		err := cfg.LoadConfig("structure.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfg.Process(automatic.Deploy)
	case test.FullCommand():
		cfg := automatic.Automatic{}
		err := cfg.LoadConfig("structure.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfg.Process(automatic.Test)
	}
}
