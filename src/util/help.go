package util

import (
	"fmt"
	"os"
)

var version = "v1.0.0"

var optionDocs = map[string]string{
	"-d": "Show debug message",
	"-v": "Show version",
	"-h": "Show help",
}

var cmds = []string{
	"insert",
	"query",
	"delete",
	"update",
	"list",
}

var cmdDocs = map[string][]string{
	"insert": []string{"qlink insert <name> <password> <room>", "Set name and password and room"},
	"query":  []string{"qlink query [-n <name>][-r <room>]", "Query user by name or room"},
	"delete": []string{"qlink delete <name> ", "Delete user by name"},
	"update": []string{"qlink update <password>", "Update user's password"},
	"list":   []string{"qlink list", "List all the users"},
}

func Version() {
	fmt.Println("qshell", version)
}

func ShowHelp(cmd string, params ...string) {
	if len(params) == 0 {
		fmt.Println(CmdList())
	} else {
		CmdHelps(params...)
	}
}

func CmdHelps(cmds ...string) {
	defer os.Exit(1)
	if len(cmds) == 0 {
		fmt.Println(CmdList())
	} else {
		for _, cmd := range cmds {
			CmdHelp(cmd)
		}
	}
}

func CmdHelp(cmd string) {
	docStr := fmt.Sprintf("Unknow cmd `%s`", cmd)
	if cmdDoc, ok := cmdDocs[cmd]; ok {
		docStr = fmt.Sprintf("Usage: %s\r\n  %s\r\n", cmdDoc[0], cmdDoc[1])
	}
	fmt.Println(docStr)
}

func CmdList() string {
	helpAll := fmt.Sprintf("QShell %s\r\n\r\n", version)
	helpAll += "Options:\r\n"
	for k, v := range optionDocs {
		helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", k, v)
	}
	helpAll += "\r\n"
	helpAll += "Commands:\r\n"
	for _, cmd := range cmds {
		if help, ok := cmdDocs[cmd]; ok {
			cmdDesc := help[1]
			helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", cmd, cmdDesc)
		}
	}
	return helpAll
}
