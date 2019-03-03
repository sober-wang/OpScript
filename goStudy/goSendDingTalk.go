package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AtAdd struct {
	AtMoblites []string `json:"atMoblies"`
	IsAtAll    bool     `json:"isAtAll"`
}

// 构建 DingTalk Robot 信息结构体
type JSONMsg struct {
	MsgType string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
	At      *AtAdd            `json:"at"`
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

	m := &JSONMsg{
		"text",
		map[string]string{
			//	"content": "This is massage from Golang wold",
			"content": msg,
		},
		&AtAdd{
			AtMoblites: []string{
				"181********",
			},
			IsAtAll: false,
		},
	}

	//	data, err := json.MarshalIndent(m, "", "    ")
	data, err := json.Marshal(m)
	dropErr(err)
	fmt.Println(m.At.AtMoblites)

	return string(data)

}

func main() {
	sendMsg := os.Args[1]
	fmt.Printf("This message's tyep is %T,Message is %s\n", sendMsg)

	dingTalkMsg := creatJSON(sendMsg)
	fmt.Println(dingTalkMsg)

}
