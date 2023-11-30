package mq

import (
	"context"
	"testing"
	"time"
)

func TestSendNotify(t *testing.T) {
	notice := New()
	notice.NewWeChat().
		SetTmpl("WZsoRBsLvbwRSk_2qr3oTkproxKbwbNjx4iwE7swJ6Y").   // 学校通知
		SetData("first", "通知标题", "").                             // 标题
		SetData("keyword1", "通知大学", "").                          // 学校
		SetData("keyword2", "通知人", "").                           // 通知人
		SetData("keyword3", time.Now().Format("2006-01-02"), ""). // 通知时间
		SetData("keyword4", "通知内容", "")                           // 通知内容
	notice.Receiver("12345678")

	p, err := NewProducer("sdk_test", "amqpURL") // 创建一个MQ连接负责消息生产
	if err != nil {
		t.Error(err)
	}

	err = p.PublishNotice(context.Background(), notice) // 发送消息
	if err != nil {
		t.Error(err)
	}
}
