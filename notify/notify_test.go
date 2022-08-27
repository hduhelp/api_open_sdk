package notify

import (
	"fmt"
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	Init("appName", "appSign")

	notice := New()
	notice.NewWeChat().
		SetTmpl("WZsoRBsLvbwRSk_2qr3oTkproxKbwbNjx4iwE7swJ6Y").   // 学校通知
		SetData("first", "通知标题", "").                             // 标题
		SetData("keyword1", "通知大学", "").                          // 学校
		SetData("keyword2", "通知人", "").                           // 通知人
		SetData("keyword3", time.Now().Format("2006-01-02"), ""). // 通知时间
		SetData("keyword4", "通知内容", "")                           // 通知内容
	notice.NewDingTalk().
		MarkdownMsg(
			fmt.Sprintf("您收到一条来自%s的通知", "通知人"),
			fmt.Sprintf(
				"# %s\n\n"+
					"#### 通知人：%s\n\n"+
					"#### 通知时间：%s\n\n"+
					"通知内容：%s", "通知标题", "通知人", time.Now().Format("2006-01-02"), "通知内容"))
	notice.Receiver("20322230", "19035405")
	err := notice.Send()
	if err != nil {
		t.Error(err)
	}
}
