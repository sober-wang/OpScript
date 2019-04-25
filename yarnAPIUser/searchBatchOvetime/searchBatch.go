package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"

	"github.com/sober-wang/sendingtalk"

	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

// 消息内容
type M struct {
	ApplicationId string `json:"ApplicationId"`
	Name          string `json:"Name"`
}

// 定义错误处理函数
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

// 获取 Yarn 中 app 信息
func getAppID(ip string, port int) []byte {
	url := fmt.Sprintf(
		"http://%s:%d/ws/v1/cluster/apps",
		ip,
		port,
	)
	log.Printf("Yarn API addressses is  %s\n", url)

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

// 挑选出 超出10分中的任务 id 并返回
func tenMin(data []byte, i, rt int) string {
	// 获取队列信息
	que := jsoniter.Get(data, i, "queue").ToString()
	// 获取 运行状态
	state := jsoniter.Get(data, i, "state").ToString()
	// 获取运行时常
	elpTime := jsoniter.Get(data, i, "elapsedTime").ToInt()
	//	startTime := jsoniter.Get(data, i, "startTime").ToInt()
	//	nowTime := int(time.Now().UnixNano() / 1e6)
	//	elpTime := nowTime - startTime

	// 筛选出异常 application
	if que == "batch" && state == "RUNNING" && elpTime > rt {
		log.Printf("被命中 => %v", jsoniter.Get(data, i, "id").ToString())
		return jsoniter.Get(data, i, "id").ToString()
	} else {
		log.Printf("被丢弃 => AppID:[ %v ] State:[ %v ] Queue:[ %v ]", jsoniter.Get(data, i, "id").ToString(), state, que)
		return ""
	}
}

// 过滤函数遍历 app 信息并调用 tenMin() 筛选
func filterData(d []byte, rt int) []M {
	a := jsoniter.Get(d, "apps", "app")
	aSize := a.Size()
	aData := []byte(a.ToString())
	var m []M
	for i := 0; i < aSize; i++ {
		appId := tenMin(aData, i, rt)
		if len(appId) != 0 {
			var sglM M
			sglM.ApplicationId = appId
			sglM.Name = jsoniter.Get(aData, i, "name").ToString()
			m = append(m, sglM)
			log.Printf("%v", sglM)
		}
	}
	// 判断报警信息 切片长度，如果为 0 就退出
	if len(m) == 0 {
		log.Println("没有批处理任务超过10分钟")
		os.Exit(1)
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

	applicationMsg := getAppID(yarnIP, yarnPort)
	runT := viper.GetInt("yarnapi.runtime") * 60 * 1000
	waitM := filterData(applicationMsg, runT)
	mB, err := jsoniter.MarshalIndent(waitM, "", "  ")
	dropErr(err)
	log.Printf("%v", string(mB))

	// 发送钉钉消息
	phList := viper.GetStringSlice("dingtalk.atwho")
	robotLink := viper.GetString("dingtalk.robotlink")
	alarmMsg := fmt.Sprintf("这些批处理任务运行时常超过 %v 分钟: %v", viper.GetInt("yarnapi.runtime"), string(mB))
	msgBody := sendingtalk.CreatMsgBody(alarmMsg, phList)
	sendingtalk.SendMsg(robotLink, msgBody)
}
