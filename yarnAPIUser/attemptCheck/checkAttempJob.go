package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/sober-wang/sendingtalk"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

// 定义一个时间类型

// 消息内容
type M struct {
	AppId         string    `json:"applicationId"`
	Name          string    `json:"name"`
	AttemptNumber int       `json:"attemptNumber"`
	StartedTime   time.Time `json:"startedTime"`
}

// 定义错误处理函数
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

// 获取 Yarn 中 app 信息
func getAPI(url string) []byte {
	//	log.Printf("Yarn API addressses is  %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("%v ; Please check your ResourceManager addresses.", err)
	}

	req.Header.Add("Content-Type", "application/json")

	cln := &http.Client{}
	resp, err := cln.Do(req)
	if err != nil {
		log.Printf("%v ; Please check your ResourceManager addresses.", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	dropErr(err)
	_ = resp.Body.Close()
	return data

}

// 从 API 返回的消息中挑选出 应用名称，application id
// 返回挑选后结果，[]M
func getAPPId(d []byte) []M {
	// app 是一个多个JSON 的列表
	a := jsoniter.Get(d, "apps", "app")
	aSize := a.Size()
	var mList []M
	aData := []byte(a.ToString())
	for i := 0; i < aSize; i++ {
		var m M
		m.AppId = jsoniter.Get(aData, i, "id").ToString()
		m.Name = jsoniter.Get(aData, i, "name").ToString()
		timeTmp := jsoniter.Get(aData, i, "startedTime").ToInt64() / 1000
		m.StartedTime = time.Unix(timeTmp, 0)
		mList = append(mList, m)
	}
	return mList

}

// 从 API 中获取 重试次数，并返回挑选结果 []M
func getAttempt(appIdList []M, ip string, port int) []M {
	var m []M
	for _, v := range appIdList {
		// 构造 Yarn 尝试 API 地址
		attemptUrl := fmt.Sprintf("http://%s:%d/ws/v1/cluster/apps/%s/appattempts",
			ip,
			port,
			v.AppId,
		)
		// 请求数据
		data := getAPI(attemptUrl)
		atn := jsoniter.Get(data, "appAttempts", "appAttempt").Size()
		// 尝试次数超过1次，数据记录为 M 结构体，并放入已准备好的 silce 中
		if atn > 1 {
			var s M
			s.Name = v.Name
			s.AppId = v.AppId
			s.AttemptNumber = atn
			s.StartedTime = v.StartedTime
			log.Printf("The [ %v ] ,attempt number is %v\n", s.AppId, s.AttemptNumber)
			m = append(m, s)
		}
	}
	return m

}

func main() {
	// 追加配置文件,kingpin 可以自动给出提示
	confFile := kingpin.Flag("conf", "config file.").String()
	kingpin.Parse()
	viper.SetConfigFile(*confFile)

	err := viper.ReadInConfig()
	dropErr(err)

	// 获取配置
	yarnIP := viper.GetString("yarnapi.ip")
	yarnPort := viper.GetInt("yarnapi.port")
	log.Printf("ResourceManager addresses is  %v", yarnIP)
	log.Printf("ResourceManager prot is %v", yarnPort)
	// Yarn application API 构造
	appIdUrl := fmt.Sprintf(
		"http://%s:%d/ws/v1/cluster/apps",
		yarnIP,
		yarnPort,
	)

	// 请求 Yarn REST API
	applist := getAPI(appIdUrl)
	// 获取 application id 列表
	appTmp := getAPPId(applist)
	// 获取 attempt 次数
	resultMsg := getAttempt(appTmp, yarnIP, yarnPort)
	// 格式化 数据
	resultByte, err := jsoniter.MarshalIndent(resultMsg, "", "  ")
	dropErr(err)

	// 判断是否发送消息到 钉钉
	if len(resultMsg) == 0 {
		os.Exit(1)
	} else {

		phList := viper.GetStringSlice("dingtalk.atwho")
		robotLink := viper.GetString("dingtalk.robotlink")
		alarmMsg := fmt.Sprintf("这些任务出现过重试: %v\n重试任务数量 => [ %v ] 个", string(resultByte), len(resultMsg))
		msgBody := sendingtalk.CreatMsgBody(alarmMsg, phList)
		sendingtalk.SendMsg(robotLink, msgBody)
	}

}
