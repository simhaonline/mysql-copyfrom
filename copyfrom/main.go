package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
	"time"
)

const VERSION = "v1"

func main() {
	initProgress()
	progress.Start()
	defer progress.Stop()

	app := cli.NewApp()
	app.Name = "copyfrom"
	app.Usage = "a tool for driving mysql Storage"
	app.Author = "Zhangjianxin"
	app.Email = "zhangjianxinnet@gmail.com"
	app.Version = fmt.Sprintf("%s %s/%s %s", VERSION,
		runtime.GOOS, runtime.GOARCH, runtime.Version())
	app.EnableBashCompletion = true
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		NewSyncCommand(),
	}

	app.Run(os.Args)
}