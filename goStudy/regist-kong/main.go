package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

// createService 在 Kong 上注册服务
func createService(svcURL, svcName string) {
	log.Printf("I'll registy address %v", svcURL)
	data := url.Values{
		"name": {svcName},
		"url":  {svcURL},
	}

	resp, err := http.PostForm("http://192.168.56.102:8001/services", data)
	if err != nil {
		log.Println("http.PostForm() =>", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() =>", err)
		return
	}
	fmt.Println(string(body))
}

// createRoutes 在 Kong 已注册的服务上注册路由
func createRoutes(svcName string) {
	path := fmt.Sprintf("/%v", svcName)
	data := url.Values{
		"name":    {svcName},
		"paths[]": {path},
	}

	routesURL := fmt.Sprintf("http://192.168.56.102:8001/services/%v/routes", svcName)
	resp, err := http.PostForm(routesURL, data)
	if err != nil {
		log.Println("http.PostForm() =>", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() =>", err)
		return
	}
	fmt.Println(string(body))
}

// mySvc 创建一个服务
func mySvc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, You through Kong Gateway access my http service")
}

func getMyIP() string {
	var myIP string
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, add := range adds {
		if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				myIP = ipnet.IP.String()
			}
		}
	}
	log.Printf("My IPv4 address is %v\n", myIP)
	return myIP
}

func main() {

	kongSvcName := "hello" // Kong service 名称
	kongSvcPort := "9527"  // 服务端口
	//kongSvcIP := getMyIP()
	kongSvcIP := "192.168.56.101" // 填写你的服务 IP
	ipPort := fmt.Sprintf("%v:%v", kongSvcIP, kongSvcPort)
	log.Println("IP and Port => ", ipPort)

	fulURL := fmt.Sprintf("http://%v", ipPort)
	createService(fulURL, kongSvcName)
	createRoutes(kongSvcName)

	http.HandleFunc("/"+kongSvcName, mySvc)

	log.Fatal(http.ListenAndServe(ipPort, nil))

}
