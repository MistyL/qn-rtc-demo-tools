package main

import (
	"cli"
	"fmt"
	"github.com/qiniu/log"
	"os"
	"runtime"
)

var supportedCmds = map[string]func(string){
	"room": cli.Room,
	// "stream": cli.Stream,
	"user": cli.User,
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//parse command
	log.SetOutputLevel(log.Lerror)
	logFp, openErr := os.OpenFile("link-anchor.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if openErr != nil {
		return
	}
	log.SetOutput(logFp)
	if len(os.Args) <= 1 {
		fmt.Println("Use help or help [cmd1 [cmd2 [cmd3 ...]]] to see supported commands.")
		return
	}

	if len(os.Args) > 2 {
		cmd := os.Args[1]
		subCmd := os.Args[2]
		if cliFunc, ok := supportedCmds[cmd]; ok {
			cliFunc(subCmd)
		} else {
			fmt.Println("Unknown cmd", fmt.Sprintf("`%s`", cmd))
		}
	} else {
		if len(os.Args) > 1 {
			//parse flags, show help or version
			option := os.Args[1]
			switch option {
			case "-v":
				cli.Version()
			case "-h":
				cli.Help()
			default:
				fmt.Println("Unknow option", fmt.Sprintf("`%s`", option))
			}
		} else {
			fmt.Println("Use -h to see supported commands.")
		}
	}
}
