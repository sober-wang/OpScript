package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type atAdd struct {
	atMoblites []string `json:"atMoblies"`
	isAtAll    bool     `json:"isAtAll"`
}

// 构建 DingTalk Robot 信息结构体
type jsonMsg struct {
	msgType string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
	At      *atAdd            `json:"at"`
}

func dropErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func creatJSON(msg string) string {

	//var msgAll = []jsonMsg{
	//	msgType: "text",
	//	Text: {
	//		"content": "This is massage from Golang wold",
	//	},
	//}

	m := &jsonMsg{
		"text",
		map[string]string{
			//	"content": "This is massage from Golang wold",
			"content": msg,
		},
		&atAdd{
			atMoblites: []string{
				"181********",
			},
			isAtAll: false,
		},
	}

	//	data, err := json.MarshalIndent(m, "", "    ")
	data, err := json.Marshal(m)
	dropErr(err)
	fmt.Println(m.At.atMoblites)

	return string(data)

}

func main() {
	sendMsg := os.Args[1]
	fmt.Printf("This message's tyep is %T,Message is %s\n", sendMsg)

	dingTalkMsg := creatJSON(sendMsg)
	fmt.Println(dingTalkMsg)

}
