package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type FSMessage struct {
	Userid string `json:"userid"`
	Message  string `json:"message"`
}

func PostToFS(userid, message, Fsurl, logsign string) string {
	open := beego.AppConfig.String("open-feishu")
	if open != "1" {
		logs.Info(logsign, "[feishu]", "飞书接口未配置未开启状态,请先配置open-feishu为1")
		return "飞书接口未配置未开启状态,请先配置open-feishu为1"
	}
	RTstring := ""
	if strings.Contains(Fsurl, "/v2/") {
		RTstring = PostToFeiShuv2(userid, message, Fsurl, logsign)
	} else {
		RTstring = PostToFeiShu("20788189213@chatroom", message, Fsurl, logsign)  //这里直接吧userid用字符串20788189213@chatroom代替了
		//RTstring = PostToFeiShu("20788189213@chatroom", message, Fsurl, logsign)  //这里直接吧userid用字符串20788189213@chatroom代替了
	}
	return RTstring
}

func PostToFeiShu(userid, message, Fsurl, logsign string) string {
	u := FSMessage{Userid: userid, Message: message}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[feishu]", b)
	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(Fsurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[feishu]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[feishu]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishu", message, "").Add(1)
	//model.AlertToCounter.WithLabelValues("feishu", userid, "").Add(0)
	logs.Info(logsign, "[feishu]", string(result))
	//fmt.Println(result)
	return string(result)

}

type Conf struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type Te struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Element struct {
	Tag      string    `json:"tag"`
	Message     Te        `json:"message"`
	Content  string    `json:"content"`
	Elements []Element `json:"elements"`
}

type Userids struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Headers struct {
	Userid    Userids `json:"userid"`
	Template string `json:"template"`
}

type Cards struct {
	Config   Conf      `json:"config"`
	Elements []Element `json:"elements"`
	Header   Headers   `json:"header"`
}

type FSMessagev2 struct {
	MsgType string `json:"msg_type"`
	Email   string `json:"email"`
	Card    Cards  `json:"card"`
}

func PostToFeiShuv2(userid, message, Fsurl, logsign string) string {
	var color string
	if strings.Count(message, "resolved") > 0 && strings.Count(message, "firing") > 0 {
		color = "orange"
	} else if strings.Count(message, "resolved") > 0 {
		color = "green"
	} else {
		color = "red"
	}
	u := FSMessagev2{
		MsgType: "interactive",
		Email:   "zhuchance@qq.com",
		Card: Cards{
			Config: Conf{
				WideScreenMode: true,
				EnableForward:  true,
			},
			Header: Headers{
				Userid: Userids{
					Content: userid,
					Tag:     "plain_message",
				},
				Template: color,
			},
			Elements: []Element{
				Element{
					Tag: "div",
					Message: Te{
						Content: message,
						Tag:     "lark_md",
					},
				},
				{
					Tag: "hr",
				},
				{
					Tag: "note",
					Elements: []Element{
						{
							Content: "PrometheusAlert    ",
							Tag:     "lark_md",
						},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[feishuv2]", b)
	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(Fsurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[feishuv2]", userid+": "+err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[feishuv2]", userid+": "+err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishuv2", message, "").Add(1)
	logs.Info(logsign, "[feishuv2]", userid+": "+string(result))
	return string(result)
}
