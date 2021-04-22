package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
)

const (
	webHookurl = "https://wechat.xiaocaicai.com/send_message"
	//webHookurl = "https://wechat.u2b.com/send_message"
	//alerts="http://172.26.132.116:9090/graph/api/v1/alerts"
	alerts="http://172.26.132.116:9090/api/v1/alerts"
)


type alprometheus struct {
	Status string `json:"status"`
	Data struct {
		Alerts []struct {
			Labels struct {
				Alertname string `json:"alertname"`
				Container string `json:"container"`
				Namespace string `json:"namespace"`
				Pod string `json:"pod"`
				Severity string `json:"severity"`
			} `json:"labels"`
			Annotations struct {
				Message string `json:"message"`
				RunbookURL string `json:"runbook_url"`
			} `json:"annotations"`
			State string `json:"state"`
			ActiveAt time.Time `json:"activeAt"`
			Value string `json:"value"`
		} `json:"alerts"`
	} `json:"data"`
}



func postwxmsg(w http.ResponseWriter, r *http.Request)  {
	//url := "http://172.29.18.4:9090/api/v1/alerts"
	//alerurl = alerts
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, alerts, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	alprometheus1 := alprometheus{}
	jsonErr := json.Unmarshal(body, &alprometheus1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(people1.Status)
	//fmt.Println(alprometheus1.Data.Alerts[0].Labels.Alertname)
	//fmt.Println(alprometheus1.Data.Alerts[0].Labels.Container)
	//fmt.Println(alprometheus1.Data.Alerts[0].Labels.Namespace)
	//fmt.Println(alprometheus1.Data.Alerts[0].Labels.Pod)
	//fmt.Println(alprometheus1.Data.Alerts[1])

	fmt.Println("告警名称：",alprometheus1.Data.Alerts[0].Labels.Alertname)
	fmt.Println("标签：",alprometheus1.Data.Alerts[0].Labels.Container)
	fmt.Println("命名空间：",alprometheus1.Data.Alerts[0].Labels.Namespace)
	fmt.Println("容器组标签：",alprometheus1.Data.Alerts[0].Labels.Pod)
	fmt.Println("状态：",alprometheus1.Data.Alerts[1].State)
	fmt.Println("开始时间：",alprometheus1.Data.Alerts[1].ActiveAt)

	////////////////////

	song := make(map[string]string)
	song["userid"] = "23610603079@chatroom"
	//song["userid"] = "123456778@qq"
	song["message"] = "告警名称：" + alprometheus1.Data.Alerts[0].Labels.Alertname + "\n" + "标签：" + alprometheus1.Data.Alerts[0].Labels.Container + "\n"+"命名空间：" + alprometheus1.Data.Alerts[0].Labels.Namespace + "\n" + "容器组标签：" +alprometheus1.Data.Alerts[0].Labels.Pod + "\n" + "状态：" + alprometheus1.Data.Alerts[1].State + "\n" + "开始时间："
	//song["message"] = "告警名称：" + alprometheus1.Data.Alerts[0].Labels.Alertname + "\n" + "标签：" + alprometheus1.Data.Alerts[0].Labels.Container + "\n" +"命名空间：" + alprometheus1.Data.Alerts[0].Labels.Namespace + "\n" + "容器组标签：" +alprometheus1.Data.Alerts[0].Labels.Pod + "\n" +"状态：" + alprometheus1.Data.Alerts[1].State + "\n" +"开始时间：" + alprometheus1.Data.Alerts[1].ActiveAt
	//song["message"] = alprometheus1.Data.Alerts[0].Labels.Alertname+"\n" + alprometheus1.Data.Alerts[0].Labels.Container+alprometheus1.Data.Alerts[0].Labels.Namespace
	//song["message"] = alprometheus1.Data.Alerts[0].Labels.Alertname+alprometheus1.Data.Alerts[0].Labels.Container+alprometheus1.Data.Alerts[0].Labels.Namespace
	//song["message"] = alprometheus1.Data.Alerts[0].Labels.Alertname+alprometheus1.Data.Alerts[0].Labels.Container
	//song["message"] = alprometheus1.Data.Alerts[0].Labels.Alertname
	bytesData, _ := json.Marshal(song)

	res, err = http.Post(webHookurl,
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	fmt.Println(*str)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	log.Printf("%s\n", body)
	fmt.Fprint(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", postwxmsg)
	http.ListenAndServe(":80", nil)
}

