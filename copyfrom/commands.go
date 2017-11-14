package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"bufio"
	"io"
	"encoding/json"
)

const (
	CHECK    = true
)

func NewSyncCommand() cli.Command {
	return cli.Command{
		Name:  "copy",
		Usage: "copy local file to mysql",
		Action: func(c *cli.Context) error {
			copy(c.String("f"),c.String("d"),c.String("t"), c.String("a"))
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "d", Usage: "mysql databases"},
			cli.StringFlag{Name: "t", Usage: "mysql table Name"},
			cli.StringFlag{Name: "f", Usage: "import file to mysql "},
			cli.StringFlag{Name: "a", Usage: "mysql address"},
		},
	}
}


func copy(filtPath string ,database string,table string ,addr string ){

	readLineJson(filtPath , func(jsonstring string) {
		var f interface{}
		b:= json.Unmarshal([]byte(jsonstring),f)
		fmt.Println(b)
	})

}


func readLineJson(filePath string,f func(data string)){
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		f(string(a)) //CallBack 回调
	}
}