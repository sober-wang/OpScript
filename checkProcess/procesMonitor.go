package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func dropErr(e error) {
	if e != nil {
		fmt.Printf("This is error %s\n", e)
	}
}

type AtSet struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// 钉钉消息头
type RobotHead struct {
	MsgType string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
	At      AtSet             `json:"at"`
}

// 钉钉机器人发送的数据
type TestResult struct {
	HostName   string   `json:"HostName"`
	ProcesName []string `json:"procesName"`
	Status     bool     `json:"Status"`
}

// 构建 钉钉机器人 请求体
func creatJSON(msg string, phoneList []string) RobotHead {
	var robotHead RobotHead
	robotHead.MsgType = "text"
	robotHead.Text = make(map[string]string)
	robotHead.Text["content"] = "进程检测结果\n" + msg + "\n进程宕了"
	robotHead.At.AtMobiles = phoneList
	robotHead.At.IsAtAll = false
	log.Print("消失的进程 : ====> %v", robotHead.Text["content"])

	return robotHead
}

// 发送消息的函数
func sendMsg(url string, data RobotHead) {
	//log.Println(data.ProcesName)
	d, err := json.Marshal(data)
	dropErr(err)
	// 构建一个新的请求，bytes.NewBuffer()传入[]byte 数据
	resq, err := http.NewRequest("POST", url, bytes.NewBuffer(d))
	dropErr(err)
	// 设置请求头
	resq.Header.Set("Content-Type", "application/json")

	// 定义客户端接收返回数据，如果不接受则请求不会请求成功
	client := &http.Client{}
	resp, err := client.Do(resq)
	dropErr(err)
	defer resp.Body.Close()

}

// 进程检测函数 通过执行 shell 命令获取相应内容
func execCMD(procesTag string) bool {
	cmdLine := "ps -ef | grep " + procesTag
	cmdOutput, err := exec.Command("sh", "-c", cmdLine).Output()
	dropErr(err)
	//tmpProces := regexp.MustCompile("\n").Split(string(cmdOutput), 4)
	tmpProces := strings.Split(string(cmdOutput), "\n")

	if len(tmpProces) > 3 {
		return true
	} else {
		return false
	}

}

// 生成发送的消息 json 数据
// 检测进程存活
func creatSendMsg(pTag []string) string {
	// 构建 消息结构体，为了转化成 JSON 化后发送
	var resultMsg TestResult
	// 找到 hostname
	tmpHostName, err := exec.Command("hostname").Output()
	dropErr(err)
	resultMsg.HostName = strings.Split(string(tmpHostName), "\n")[0]

	for _, v := range pTag {
		if execCMD(v) == false {
			resultMsg.Status = false
			resultMsg.ProcesName = append(resultMsg.ProcesName, v)
		}
	}
	// 如果挂掉的进程列表 == 0 那么程序就退出
	if len(resultMsg.ProcesName) == 0 {

		os.Exit(1)
	}

	m, err := json.MarshalIndent(resultMsg, "", "    ")
	dropErr(err)

	return string(m)
}

func main() {
	// 传参数 配置文件地址
	filePath := os.Args[1]
	// 给 viper 解析器穿一个文件
	viper.SetConfigFile(filePath)
	// 读配置文件,出错会返回 error
	err := viper.ReadInConfig()
	dropErr(err)
	// 定义一个配置文件的 对应的变量 类型map[string]interface{}
	var confile map[string]interface{}
	viper.Unmarshal(&confile)

	// 构建 发送消息的内容，进程检测结果的 json
	pList := viper.GetStringSlice("procesList")
	msgText := creatSendMsg(pList)

	// 发送消息
	phList := viper.GetStringSlice("dingTalkRobot.atWho")
	rbHead := creatJSON(msgText, phList)
	rbLink := viper.GetStringSlice("dingTalkRobot.robotLink")
	sendMsg(rbLink[0], rbHead)

}
