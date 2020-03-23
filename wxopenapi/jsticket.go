package wxopenapi

import (
	"fmt"
	"github.com/jmniu/weixin-third-dev/wxopencrypt"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

const (
	URL_GET_TICKET = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%v&type=jsapi"
)

func (this *WxOpen) GetJsTicket(accessToken string) string {
	var ret struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Ticket    string `json:"ticket"`
		ExpiresIn int    `json:"expires_in"`
	}
	var url = fmt.Sprintf(URL_GET_TICKET, accessToken)
	gorequest.New().Get(url).EndStruct(&ret)
	return ret.Ticket
}

func (this *WxOpen) GetJsTicketSignature(accessToken string, nonce string, timestamp int, url string) string {
	jsTicket := this.GetJsTicket(accessToken)
	_, signature := wxopencrypt.NewWXBizMsgCrypt().ComputeJsTicketSignature(jsTicket, strconv.Itoa(timestamp), nonce, url)
	return signature
}
