package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var roomSubCmdAlias = map[string]string{
	"q": "query",
	"c": "create",
	"d": "delete",
}

var supportedRoomSubCmds = map[string]func(string, string){
	"query":  QueryRoomInfo,
	"create": CreateRoom,
	"delete": DeleteRoom,
	"token":  GetRoomToken,
}

func Room(subCmd string) {
	//add alias support
	if vCmd, vOk := roomSubCmdAlias[subCmd]; vOk {
		subCmd = vCmd
	}

	//parse and exec sub cmd
	if subCmdFunc, ok := supportedRoomSubCmds[subCmd]; ok {
		subCmdFunc("room", subCmd)
	} else {
		fmt.Println("Unknown cmd", fmt.Sprintf("`%s`", subCmd), "for room")
	}
}

//////////////////////////////////////////////////////////////////////////
func QueryRoomInfo(cmd, subCmd string) {
	if len(os.Args) < 3 {
		fmt.Println("Invalid params.")
		return
	}
	roomName := strings.TrimSpace(os.Args[3])
	// fmt.Printf("roomname = %s\n", roomName)
	if roomName == "" {
		fmt.Println("Invalid params.")
		return
	}

	url := fmt.Sprintf("%s/pili/room/%s/status", BASE_URL, roomName)
	// fmt.Printf("query url : %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("query room failur, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func CreateRoom(cmd, subCmd string) {
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var user string
	var room string

	flagSet.StringVar(&user, "u", "", "user name")
	flagSet.StringVar(&room, "r", "", "room name")

	flagSet.Parse(os.Args[3:])
	if user == "" || room == "" {
		fmt.Println("Invalid params.")
		return
	}
	user = strings.TrimSpace(user)
	room = strings.TrimSpace(room)

	url := fmt.Sprintf("%s/pili/room/%s/user/%s/create", BASE_URL, room, user)
	reqeust, _ := http.NewRequest("POST", url, nil)
	response, err := http.DefaultClient.Do(reqeust)
	if err != nil {
		fmt.Printf("create room failur, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func DeleteRoom(cmd, subCmd string) {
	if len(os.Args) < 3 {
		fmt.Println("Invalid params.")
		return
	}
	roomName := os.Args[3]
	if roomName == "" {
		fmt.Println("Invalid params.")
		return
	}
	url := fmt.Sprintf("%s/pili/room/%s/delete", BASE_URL, roomName)
	// fmt.Printf("delete room url : %s\n", url)
	request, _ := http.NewRequest("POST", url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("delete room failur, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func GetRoomToken(cmd, subCmd string) {
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var user string
	var room string

	flagSet.StringVar(&user, "u", "", "user name")
	flagSet.StringVar(&room, "r", "", "room name")

	flagSet.Parse(os.Args[3:])
	if user == "" || room == "" {
		fmt.Println("Invalid params.")
		return
	}
	user = strings.TrimSpace(user)
	room = strings.TrimSpace(room)

	url := fmt.Sprintf("%s/pili/room/%s/user/%s/token/", BASE_URL, room, user)
	request, _ := http.NewRequest("POST", url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("get roomtoken failuer, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
