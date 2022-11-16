package notify

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"sync"
)

var notifierPool *sync.Pool

// Init 内部应用认证初始化
func Init(appName, appSign string) {
	notifierPool = &sync.Pool{
		New: func() interface{} {
			return gorequest.New().
				Post("https://api.hduhelp.com/notify?from="+appName).
				Set("sign", appSign)
		},
	}
}

// InitOAuthApp OAuth应用认证初始化 （发送范围仅限授权过该OAuth应用的用户）
func InitOAuthApp(clientId, clientSecret string) {
	notifierPool = &sync.Pool{
		New: func() interface{} {
			return gorequest.New().
				Post("https://api.hduhelp.com/notify").
				Param("client_id", clientId).
				Param("client_secret", clientSecret)
		},
	}
}

// InitPrivateApp 个人token认证初始化 （发送范围仅限token对应用户）
func InitPrivateApp(token *string) {
	notifierPool = &sync.Pool{
		New: func() interface{} {
			return gorequest.New().
				Post("https://api.hduhelp.com/notify").
				Param("auth", *token)
		},
	}
}

type SendResult struct {
	Cache bool `json:"cache"`
	Data  struct {
		Failed  int `json:"failed"`
		Success int `json:"success"`
	} `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

type Template struct {
	StaffID  []string            `json:"StaffID"`
	WeChat   *WeChatNoticeBody   `json:"template,omitempty"`
	DingTalk *DingTalkNoticeBody `json:"dingtalk,omitempty"`
	App      *UniNotifyInput     `json:"body,omitempty"`
}

func New() *Template {
	return &Template{}
}

func (t *Template) Receiver(staffId ...string) *Template {
	t.StaffID = append(t.StaffID, staffId...)
	return t
}

func (t *Template) NewWeChat() *WeChatNoticeBody {
	t.WeChat = &WeChatNoticeBody{
		ToUser: "HDUHELP_TRANSFER_PARAM_OPENID",
	}
	return t.WeChat
}

func (t *Template) NewDingTalk() *DingTalkNoticeBody {
	t.DingTalk = &DingTalkNoticeBody{}
	return t.DingTalk
}

func (t *Template) NewApp(title, subtitle, content, extra string) *UniNotifyInput {
	t.App = &UniNotifyInput{
		Title:    title,
		SubTitle: subtitle,
		Content:  content,
		Extra:    extra,
	}
	return t.App
}

func (t *Template) Send() error {
	result := &SendResult{}

	notifier := notifierPool.Get().(*gorequest.SuperAgent)
	defer notifierPool.Put(notifier)
	_, _, errs := notifier.Send(t).EndStruct(result)

	if errs != nil || result.Msg != "success" || result.Data.Failed != 0 {
		return errors.New("send notify failed")
	}
	return nil
}

type WeChatNoticeBody struct {
	Data        map[string]*TemplateDataItem `json:"data"`
	TemplateID  string                       `json:"template_id"`
	ToUser      string                       `json:"touser"`
	URL         string                       `json:"url"`
	MiniProgram *MiniProgram                 `json:"miniprogram,omitempty"`
}

func (t *WeChatNoticeBody) SetTmpl(templateId string) *WeChatNoticeBody {
	t.TemplateID = templateId
	return t
}

func (t *WeChatNoticeBody) SetData(key, value, color string) *WeChatNoticeBody {
	if t.Data == nil {
		t.Data = make(map[string]*TemplateDataItem)
	}
	t.Data[key] = &TemplateDataItem{Value: value, Color: color}
	return t
}

func (t *WeChatNoticeBody) Url(url string) *WeChatNoticeBody {
	t.URL = url
	return t
}

func (t *WeChatNoticeBody) MiniPro(appid, pagePath string) *WeChatNoticeBody {
	t.MiniProgram = &MiniProgram{
		AppId:    appid,
		PagePath: pagePath,
	}
	return t
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type MiniProgram struct {
	AppId    string `json:"appid"`    // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
	PagePath string `json:"pagepath"` // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
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
