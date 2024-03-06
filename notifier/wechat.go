package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zuoyangs/go-alertmanager-wechatrobot-webhook/model"
	"github.com/zuoyangs/go-alertmanager-wechatrobot-webhook/template"
)

// Send send message to wechat
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := template.TemplateToMarkdown(notification)

	if err != nil {
		log.Printf("template.TemplateToMarkdown 出错:%s", err.Error())

	}
	log.Printf("markdown:%+v", markdown)

	data, err := json.Marshal(markdown)

	if err != nil {
		log.Printf("json.Marshal 出错:%s", err.Error())
	}

	var wechatRobotURL string

	if robotURL != "" {
		wechatRobotURL = robotURL
	} else {
		wechatRobotURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + defaultRobot
	}

	req, err := http.NewRequest(
		"POST",
		wechatRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		log.Printf("http.NewRequest 出错:%s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("client.Do 出错:%s", err.Error())
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
