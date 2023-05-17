package utils

import (
	"fmt"
	"github.com/avct/uasurfer"
	"testing"
)

func TestParseUA(t *testing.T) {
	ua := "Mozilla/5.0 (Linux; U; Android 13; zh-CN; PEMT00 Build/TP1A.220905.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 UWS/3.22.1.249 Mobile Safari/537.36 AliApp(DingTalk/7.0.0.11) com.alibaba.android.rimet/27811579 Channel/263200 language/zh-CN abi/64 UT4Aplus/0.2.25 colorScheme/dark"
	fmt.Println(_parseUA(ua))
}

func _parseUA(s string) string {
	ua := uasurfer.Parse(s)
	//return ua.DeviceType.String()
	var browser, os, deviceType string
	switch ua.Browser.Name {
	case uasurfer.BrowserChrome:
		browser = "Chrome"
	case uasurfer.BrowserIE:
		browser = "Edge"
	default:
		browser = ua.Browser.Name.StringTrimPrefix()
	}

	os = ua.OS.Name.StringTrimPrefix()

	switch ua.DeviceType {
	case uasurfer.DeviceComputer:
		deviceType = "PC端"
	case uasurfer.DevicePhone:
		deviceType = "手机端"
	default:
		deviceType = ua.DeviceType.StringTrimPrefix()
	}
	return fmt.Sprintf("%s （%s-%s）", browser, os, deviceType)
}
