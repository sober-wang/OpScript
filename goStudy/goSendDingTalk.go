package main

import (
	"encoding/json"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// 构建 是否 @ 人的结构体
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

// 错误处理函数
func dropErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type ConfMsg map[string]interface{}

// 处理配置文件 Yaml
func readConf(f string) ConfMsg {

	var c ConfMsg

	fileMsg, err := ioutil.ReadFile(f)
	dropErr(err)

	err = yaml.Unmarshal(fileMsg, &c)
	dropErr(err)

	return c

}

// 创建DingTalk 机器人接收信息
func creatJSON(msg string, atList []string) string {

	m := &JSONMsg{
		"text",
		map[string]string{
			//	"content": "This is massage from Golang wold",
			"content": msg,
		},
		&AtAdd{
			AtMoblites: atList,
			//			AtMoblites: []string{
			//				"181********",
			//			},
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

	filePath := os.Args[2]
	confMsg := readConf(filePath)
	robotConf, ok := confMsg["dingTalk"].(map[interface{}]interface{})
	fmt.Println(robotConf["robotLink"])
	fmt.Printf("%s %T\n", ok, robotConf["atWho"])

	phoneNumber := robotConf["atWho"]
	fmt.Printf("%T\n", phoneNumber)

	//	dingTalkMsg := creatJSON(sendMsg)
	//	fmt.Println(dingTalkMsg)

}
