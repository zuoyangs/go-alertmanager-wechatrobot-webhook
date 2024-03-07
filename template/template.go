package template

import (
	"bytes"
	"fmt"
	"log"

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
		buffer.WriteString(fmt.Sprintln("\n><font color=\"red\">【MySQL服务告警】"))
		buffer.WriteString(fmt.Sprintf("\n>服务名称: %s\n", "假设我是服务名称"))
		buffer.WriteString(fmt.Sprintf("\n>实例信息: %s\n", "假设我是实例信息"))
		buffer.WriteString(fmt.Sprintf("\n>通知策略: %s\n", "假设我是通知策略"))
		buffer.WriteString(fmt.Sprintf("\n>Slave: %s\n", "true"))
		buffer.WriteString(fmt.Sprintf("\n>上次报警时间 %s\n", "假设我上次报警时间"))
		buffer.WriteString(fmt.Sprintf("\n>当日告警次数 %s\n", "假设我是第2次"))
		buffer.WriteString(fmt.Sprintf("\n>告警级别: %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>告警类型: %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>故障主机: %s\n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n>当前1分钟平均cpu使用率 %s\n", "假设我是50%"))
		buffer.WriteString(fmt.Sprintf("\n>当前1分钟平均内存使用率: %s\n", "假设我是50%"))
		buffer.WriteString(fmt.Sprintf("\n>当前1分钟磁盘/data空间使用率: %s\n", "假设我是50%"))
		buffer.WriteString(fmt.Sprintf("\n>通知人: %s\n", "假设我是运维/开发值班人员"))
		buffer.WriteString(fmt.Sprintf("\n>认领告警: %s\n", "假设我是认领告警按钮"))
		buffer.WriteString(fmt.Sprintf("\n>解决告警: %s\n", "假设我是解决告警按钮"))
		buffer.WriteString(fmt.Sprintf("\n>关注告警: %s\n", "假设我是关注告警中按钮"))
		buffer.WriteString(fmt.Sprintf("\n>推送告警: %s\n", "假设我是推送告警按钮"))
		buffer.WriteString(fmt.Sprintf("\n>屏蔽告警: %s\n", "假设我是带周期的屏蔽告警按钮"))
		buffer.WriteString(fmt.Sprintf("\n>未解决: %s\n", "假设我是未解决告警按钮"))
		buffer.WriteString(fmt.Sprintf("\n>从可观测平台查询搜索与该mysql关联服务: %s\n", "假设我是相关服务"))
		buffer.WriteString(fmt.Sprintf("\n>当月/当季度服务等级目标SLO/SLA: %s\n", "假设我是4个9"))
		buffer.WriteString(fmt.Sprintf("\n>报警间隔: %s\n", "假设我是30min"))

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

	log.Println(buffer.String())

	return
}
