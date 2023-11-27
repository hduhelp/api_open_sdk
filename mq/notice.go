package mq

import (
	"context"
	"encoding/json"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type NoticeTemplate struct {
	ctx      context.Context
	StaffID  []string            `json:"StaffID"`
	WeChat   *WeChatNoticeBody   `json:"template,omitempty"`
	DingTalk *DingTalkNoticeBody `json:"dingtalk,omitempty"`
}

func New() *NoticeTemplate {
	return &NoticeTemplate{}
}

func (t *NoticeTemplate) WithContext(ctx context.Context) *NoticeTemplate {
	t.ctx = ctx
	return t
}

func (t *NoticeTemplate) Receiver(staffId ...string) *NoticeTemplate {
	t.StaffID = append(t.StaffID, staffId...)
	return t
}

func (t *NoticeTemplate) NewWeChat() *WeChatNoticeBody {
	t.WeChat = &WeChatNoticeBody{}
	return t.WeChat
}

func (t *NoticeTemplate) NewDingTalk() *DingTalkNoticeBody {
	t.DingTalk = &DingTalkNoticeBody{}
	return t.DingTalk
}

type WeChatNoticeBody struct {
	Data       map[string]*message.TemplateDataItem `json:"data"`
	TemplateID string                               `json:"template_id"`
	ToUser     string                               `json:"touser"`
	URL        string                               `json:"url"`
}

func (t *WeChatNoticeBody) SetTmpl(templateId string) *WeChatNoticeBody {
	t.TemplateID = templateId
	return t
}

func (t *WeChatNoticeBody) SetData(key, value, color string) *WeChatNoticeBody {
	if t.Data == nil {
		t.Data = make(map[string]*message.TemplateDataItem)
	}
	t.Data[key] = &message.TemplateDataItem{Value: value, Color: color}
	return t
}

func (t *WeChatNoticeBody) Url(url string) *WeChatNoticeBody {
	t.URL = url
	return t
}

type DingTalkNoticeBody struct {
	MsgType  string                     `json:"msgtype"`
	Text     *DingTalkMsgPlainTextInput `json:"text,omitempty"`
	Link     *DingTalkMsgLinkInput      `json:"link,omitempty"`
	Markdown *DingTalkMsgMarkdownInput  `json:"markdown,omitempty"`
}

func (t *DingTalkNoticeBody) PlainText(content string) {
	t.MsgType = "text"
	t.Text = &DingTalkMsgPlainTextInput{
		Content: content,
	}
}

func (t *DingTalkNoticeBody) LinkMsg(messageUrl, picUrl, title, text string) {
	t.MsgType = "link"
	t.Link = &DingTalkMsgLinkInput{
		MessageUrl: messageUrl,
		PicUrl:     picUrl,
		Title:      title,
		Text:       text,
	}
}

func (t *DingTalkNoticeBody) MarkdownMsg(title, text string) {
	t.MsgType = "markdown"
	t.Markdown = &DingTalkMsgMarkdownInput{
		Title: title,
		Text:  text,
	}
}

type DingTalkMsgPlainTextInput struct {
	Content string `json:"content"`
}

type DingTalkMsgLinkInput struct {
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

type DingTalkMsgMarkdownInput struct {
	Title string `json:"title"` // 这个是点进推送列表前显示的简短的消息（进入后不可见）
	Text  string `json:"text"`  // 这个是进入后显示的详细的消息
}

type UniNotifyInput struct {
	Title    string
	SubTitle string
	Content  string
	Extra    string
}

type WechatBody struct {
	StaffId string `json:"staffId"`
	*WeChatNoticeBody
	MessageId string `json:"messageId"`
}

func MarshalWechatBody(staffId string, msg *WeChatNoticeBody) (data []byte, err error) {
	body := &WechatBody{
		StaffId:          staffId,
		WeChatNoticeBody: msg,
	}
	if data, err = json.Marshal(body); err != nil {
		return nil, err
	}
	return data, nil
}

func UnmarshalWechatBody(data []byte) (*WechatBody, error) {
	body := &WechatBody{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

type DingTalkBody struct {
	StaffId string `json:"staffId"`
	*DingTalkNoticeBody
	MessageId string `json:"messageId"`
}

func MarshalDingTalkBody(staffId string, msg *DingTalkNoticeBody) (data []byte, err error) {
	body := &DingTalkBody{
		StaffId:            staffId,
		DingTalkNoticeBody: msg,
	}
	if data, err = json.Marshal(body); err != nil {
		return nil, err
	}
	return data, nil
}

func UnmarshalDingTalkBody(data []byte) (*DingTalkNoticeBody, error) {
	body := &DingTalkNoticeBody{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
