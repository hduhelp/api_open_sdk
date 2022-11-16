# gateway-notify
geteway的附属notify sdk

## Usage

```go
import "https://github.com/hduhelp/api_open_sdk/notify"

func init() {
    notify.Init("appName","appSign")

    //InitOAuthApp("clientId", "clientSecret")

    //token := "12345678-236c-473a-ba32-1234567890"
    //InitPrivateApp(&token)
}

func main() {
    notice := New()
    notice.NewWeChat().
        SetTmpl("WZsoRBsLvbwRSk_2qr3oTkproxKbwbNjx4iwE7swJ6Y"). // 学校通知
        SetData("first", "通知标题", ""). // 标题
        SetData("keyword1", "通知大学", ""). // 学校
        SetData("keyword2", "通知人", ""). // 通知人
        SetData("keyword3", time.Now().Format("2006-01-02"), ""). // 通知时间
        SetData("keyword4", "通知内容", "") // 通知内容
    notice.NewDingTalk().
        MarkdownMsg(
            fmt.Sprintf("您收到一条来自%s的通知", "通知人"),
            fmt.Sprintf(
            "# %s\n\n"+
            "#### 通知人：%s\n\n"+
            "#### 通知时间：%s\n\n"+
            "通知内容：%s", "通知标题", "通知人", time.Now().Format("2006-01-02"), "通知内容"))
    notice.Receiver("接收者1的学号","接收者2的学号")
    notice.Receiver("接收者3的学号")
    _ = notice.Send()    
}
```