package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var userSubCmdAlias = map[string]string{
	"q":  "query",
	"rg": "register",
	"d":  "delete",
}

var supportedUserSubCmds = map[string]func(string, string){
	"query":    UserInfo,
	"register": CreateUser,
	"delete":   DeleteUser,
	"update":   UpdateUserInfo,
}

type RequestBody struct {
	Name     string `json:"user"`
	Password string `json:"password"`
	Flag     string `json:"flag"`
}

func User(subCmd string) {
	//add alias support
	if vCmd, vOk := userSubCmdAlias[subCmd]; vOk {
		subCmd = vCmd
	}

	//parse and exec sub cmd
	if subCmdFunc, ok := supportedUserSubCmds[subCmd]; ok {
		subCmdFunc("user", subCmd)
	} else {
		fmt.Println("Unknown cmd", fmt.Sprintf("`%s`", subCmd), "for user")
	}
}

//////////////////////////////////////////////////////////////////////////////
func UserInfo(cmd, subCmd string) {
	if len(os.Args) < 3 {
		fmt.Println("Invalid params.")
		return
	}
	userName := strings.TrimSpace(os.Args[3])
	if userName == "" {
		fmt.Println("Invalid params.")
		return
	}

	url := fmt.Sprintf("%s/pili/user/%s/status", BASE_URL, userName)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("query user failur , %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

// func Login(cmd, subCmd string) {
// 	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
// 	flagSet.Usage = func() {
// 		CmdHelp(cmd, subCmd)
// 	}

// 	var user string
// 	var password string

// 	flagSet.StringVar(&user, "u", "", "user name")
// 	flagSet.StringVar(&password, "p", "", "password")

// 	flagSet.Parse(os.Args[3:])
// 	if user == "" || password == "" {
// 		fmt.Println("Invalid params.")
// 		return
// 	}
// 	user = strings.TrimSpace(user)
// 	password = strings.TrimSpace(password)

// 	url := fmt.Sprintf("%s/rtc/login", BASE_URL)
// 	var reqBody RequestBody
// 	if user != "" {
// 		reqBody.Name = user
// 	}
// 	if password != "" {
// 		reqBody.Password = password
// 	}
// 	req, _ := json.Marshal(reqBody)
// 	reqbody := bytes.NewBuffer([]byte(req))
// 	// fmt.Println(string(req))
// 	request, _ := http.NewRequest("POST", url, reqbody)
// 	request.Header.Set("Content-Type", "application/json")
// 	response, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		fmt.Printf("login failur, %s\n", err.Error())
// 		return
// 	}
// 	defer response.Body.Close()
// 	body, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(body))
// }

func CreateUser(cmd, subCmd string) {
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var user string
	var room string
	var password string
	var flag string

	flagSet.StringVar(&user, "u", "", "user name")
	flagSet.StringVar(&room, "r", "", "room name")
	flagSet.StringVar(&password, "p", "", "password")
	flagSet.StringVar(&flag, "f", "", "user's type")

	flagSet.Parse(os.Args[3:])
	if user == "" || room == "" || password == "" {
		fmt.Println("Invalid params.")
		return
	}
	user = strings.TrimSpace(user)
	room = strings.TrimSpace(room)
	password = strings.TrimSpace(password)
	flag = strings.TrimSpace(flag)

	url := fmt.Sprintf("%s/pili/register/user/%s/pwd/%s/room/%s/", BASE_URL, user, password, room)
	if flag != "" {
		url += flag
	}
	fmt.Printf("url : %s\n", url)
	request, _ := http.NewRequest("POST", url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("register user failur, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func DeleteUser(cmd, subCmd string) {
	if len(os.Args) < 3 {
		fmt.Println("Invalid params.")
		return
	}
	userName := strings.TrimSpace(os.Args[3])
	if userName == "" {
		fmt.Println("Invalid params.")
		return
	}

	url := fmt.Sprintf("%s/pili/user/%s/delete", BASE_URL, userName)
	request, _ := http.NewRequest("POST", url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("delete user failur , %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func UpdateUserInfo(cmd, subCmd string) {
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var user string
	var password string
	var flag string

	flagSet.StringVar(&user, "u", "", "user name")
	flagSet.StringVar(&password, "p", "", "password")
	flagSet.StringVar(&flag, "f", "", "user's type")

	flagSet.Parse(os.Args[3:])
	if user == "" {
		fmt.Println("Invalid params.")
		return
	}
	user = strings.TrimSpace(user)
	password = strings.TrimSpace(password)
	flag = strings.TrimSpace(flag)

	url := fmt.Sprintf("%s/pili/user/%s/update", BASE_URL, user)
	var reqBody RequestBody
	if password != "" {
		reqBody.Password = password
	}
	if flag != "" {
		reqBody.Flag = flag
	}
	req, _ := json.Marshal(reqBody)
	reqbody := bytes.NewBuffer([]byte(req))
	// fmt.Println(string(req))
	request, _ := http.NewRequest("POST", url, reqbody)
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("update info failur, %s\n", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
