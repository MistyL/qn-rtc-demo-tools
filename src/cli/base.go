package cli

import (
	"fmt"
)

const VERISON = "1.0.1"
const BASE_URL = "https://api.qiniudemo.com"

var roomCmdOrder = []string{"query", "create", "delete", "token"}

// var streamCmdOrder = []string{"create", "rtmp-live", "hls-play"}
var userCmdOrder = []string{"query", "register", "update", "delete", "login"}

var helpInfo = map[string]map[string]string{
	"room": map[string]string{
		"query":  "qlink room query <RoomName>",
		"create": "qlink room create -r <RoomName> -u <UserName>",
		"delete": "qlink room delete <RoomName>",
		"token":  "qlink room token -r <RoomName> -u <UserName>",
	},
	"user": map[string]string{
		"query":    "qlink user query <UserName>",
		"register": "qlink user register -u <UserName> -p <Password> -r <RoomName> [-f <Flag>]",
		"update":   "qlink user update -u <UserName> [-p <Password> ] [-f <Flag>]",
		"delete":   "qlink user delete <UserName>",
		"login":    "qlink user login -u <UserName> -p <Password>",
	},
}

func Version() {
	fmt.Println("QLink", VERISON)
}

func Help() {
	fmt.Println("QLink", VERISON)
	fmt.Println()
	fmt.Println("Commands for user:")
	for _, cmd := range userCmdOrder {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, helpInfo["user"][cmd]))
	}

	fmt.Println()
	fmt.Println("Commands for room:")
	for _, cmd := range roomCmdOrder {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, helpInfo["room"][cmd]))
	}

	fmt.Println()
	// fmt.Println("Commands for stream:")
	// for _, cmd := range streamCmdOrder {
	// 	fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, helpInfo["stream"][cmd]))
	// }
	// fmt.Println()
}

func CmdHelp(cmd string, subCmd string) {
	fmt.Println("Usage:", helpInfo[cmd][subCmd])
}

type CliFunc func(cmd string, params ...string)
