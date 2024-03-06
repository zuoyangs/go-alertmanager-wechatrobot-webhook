package template

import (
	"bytes"
	"fmt"

	"github.com/zuoyangs/go-alertmanager-wechatrobot-webhook/model"
)

// Template 将告警信息转换为markdown格式
func TemplateToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 当前状态:%s \n", status))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		buffer.WriteString(fmt.Sprintf("\n>告警级别: %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>告警类型: %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>故障主机: %s\n", labels["instance"]))

		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n>告警主题: %s\n", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>告警详情: %s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 触发时间: %s\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
