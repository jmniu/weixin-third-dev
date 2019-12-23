package wxopenapi

import (
	"fmt"
	"github.com/wendal/errors"
	"strings"
	"time"
)

const (
	//设置公众号的菜单
	URL_SET_MENU = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%v"
)

func (this *WxOpen) SetMenu(accessToken string, menu string) {
	fmt.Println("SetMenu ", time.Now().Unix())
	if len(accessToken) <= 0 {
		fmt.Println("accessToken is empty")
		return
	}

	//reqstr, _ := json.Marshal(menu)
	reqstr := []byte(menu)
	rsp, err := PostJsonByte(fmt.Sprintf(URL_SET_MENU, accessToken), reqstr)
	if err != nil {
		fmt.Println("请求 SetMenu 失败 err", err.Error())
		return
	}
	if strings.Contains(string(rsp), "errcode") {
		fmt.Println("请求 SetMenu 失败 rsp", string(rsp))
		if strings.Contains(string(rsp), "access_token expired") {
			err = errors.New("口令过期")
		} else {
			err = errors.New(string(rsp))
		}
		return
	}
	fmt.Println("请求 SetMenu 完成")
}
