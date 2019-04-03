package main

/*
	用 Golang 重新实现 HDFS 小文件筛选工具
	写 CSV 目前还没有实现。感觉比较麻烦所以没做
*/

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

const (
	showDir      string = "LISTSTATUS"
	showProperty string = "GETCONTENTSUMMARY"
)

// 等待写入 csv 的数据 对应的 struct
type ResultSave struct {
	DirPath       string `json:"dirPath"`
	FileSize      int64  `json:"fileSize"`
	SpaceConsumed int64  `json:"spaceConsumed"`
	FileCount     int    `json:"fileCount"`
	AverageSize   int    `json:"averageSize"`
}

// HDFS API 目录属性 struct
type FirstSummaryDir struct {
	ContentSummary ContentSummar `json:"ContentSummary"`
}
type TypeQuota struct {
}
type ContentSummar struct {
	DirectoryCount int       `json:"directoryCount"`
	FileCount      int       `json:"fileCount"`
	Length         int64     `json:"length"`
	Quota          int64     `json:"quota"`
	SpaceConsumed  int64     `json:"SpaceConsumed"`
	SpaceQuota     int       `json:"spaceQuota"`
	TypeQuota      TypeQuota `json:"typeQuota"`
}

// HDFS API ls目录 第二层 JSON 内部信息定义
type End struct {
	AccessTime       int    `json:"accessTime"`
	BlockSize        int    `json:"blockSize"`
	ChildrenNum      int    `json:"childrenNum"`
	FileId           int    `json:"fileId"`
	Group            string `json:"group"`
	Length           int    `json:"length"`
	ModificationTime int64  `json:"modificationTime"`
	Owner            string `json:"owner"`
	PathSuffix       string `json:"pathSuffix"`
	Permission       string `json:"permission"`
	Replication      int    `json:"replication"`
	StoragePolicy    int    `json:"storagePolicy`
	Type             string `json:"type"`
}

// 定义 HDFS JSON ls目录 第二层 struct
type FileStatuses struct {
	FileStatus []End `json:"FileStatus`
}

// 定义 HDFS API ls目录 JSON头 struct
type Hel struct {
	Wol FileStatuses `json:"FileStatuses"`
}

// 错误处理函数
func dropErr(e error) {
	if e != nil {
		fmt.Printf("[ ERROR ] This have a Error, it's %s\n ", e)
	}
}

// 列出目录 ,传入map[string]interface{} 返回 请求结果的 []byte
func operationAPI(conf map[string]interface{}, op string, p string) []byte {

	url := fmt.Sprintf(
		"http://%s:%d/webhdfs/v1/%s?user.name=%s&op=%s",
		conf["ip"].(string),
		conf["port"].(int),
		p,
		conf["username"].(string),
		op,
	)
	//fmt.Printf("HDFS API 请求URL: %s\n\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		dropErr(err)
		var b []byte
		return b
	}

	//req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	// 构建客户端
	client := &http.Client{}
	resp, err := client.Do(req)
	// 函数结束前关闭连接
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	dropErr(err)
	return b

}

// 筛选目录 传入 JSON []byte,返回筛选后的目录 []string
func selectDir(respRsult []byte) []string {
	// 使用定义好的 struct 对应传入数据
	var dirListJSON Hel
	// 定义返回值
	var dirList []string
	json.Unmarshal(respRsult, &dirListJSON)

	for _, v := range dirListJSON.Wol.FileStatus {
		dirList = append(dirList, v.PathSuffix)
	}

	return dirList
}

// 统计目录文件大小
func countDir(smyJSON []byte, pathName string) ResultSave {

	var smyStruct FirstSummaryDir
	json.Unmarshal(smyJSON, &smyStruct)

	resultSize := ResultSave{
		FileSize:      smyStruct.ContentSummary.Length / 1024 / 1024,
		SpaceConsumed: smyStruct.ContentSummary.SpaceConsumed / 1024 / 1024,
		FileCount:     smyStruct.ContentSummary.FileCount,
		DirPath:       pathName,
	}

	// 如果文件大小等于0，则文件平均大小为0
	if int(resultSize.FileSize) == 0 {
		resultSize.AverageSize = 0
		return resultSize
	} else {
		// 求出目录下文件大小平均值
		resultSize.AverageSize = int(resultSize.FileSize) / resultSize.FileCount
		return resultSize
	}
}

// 求结果列表
func resulList(conf map[string]interface{}, sDir []string) []ResultSave {

	var resultSlice []ResultSave
	averageSize := conf["averagesize"].(int)
	fatherPath := conf["dirpath"].(string)
	for _, v := range sDir {
		aPath := fmt.Sprintf("%s/%s", fatherPath, v)
		summaryJSON := operationAPI(conf, showProperty, aPath)
		a := countDir(summaryJSON, aPath)

		// 筛选 合规数据，HDFS 中的数据 平均大小必须大于 averageSize 设置的值
		if a.AverageSize > averageSize {
			break
		} else {
			resultSlice = append(resultSlice, a)
		}
	}
	fmt.Println(resultSlice)

	return resultSlice
}

func wFile(conf map[string]interface{}, message []ResultSave) {
	resultFile := conf["resultfile"].(string)
	f, err := os.OpenFile(resultFile, os.O_WRONLY, 0644)
	defer f.Close()
	dropErr(err)

	for _, v := range message {
		fmt.Println(v)
		waitMsg, err := json.Marshal(v)
		dropErr(err)
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt(waitMsg, n)
		dropErr(err)

	}
}

// 主函数
func main() {
	filePath := os.Args[1]

	fmt.Printf("配置文件名称：%s\n", filePath)
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	dropErr(err)
	confMsg := viper.GetStringMap("hdfsapi")

	dirListByte := operationAPI(confMsg, showDir, viper.GetString("hdfsapi.dirpath"))
	sunDir := selectDir(dirListByte)
	r := resulList(confMsg, sunDir)

	wFile(confMsg, r)

}
