package wx_third_dev

import (
	"github.com/jmniu/wx-third-dev/wxopenapi"
	"time"
)

func InitWxOpen(componentToken string,
	componentAesKey string,
	componentAppID string,
	componentAppSecret string,
	componentVerifyTicket string,
	componentAccessToken string) {

	wxOpen := wxopenapi.GWxOpen
	wxOpen.Init(componentToken, componentAesKey, componentAppID, componentAppSecret)
	wxOpen.SetInfo(wxopenapi.COMPONENT_VERIFY_TICKET, componentVerifyTicket, time.Now().Unix(), 600)
	wxOpen.SetInfo(wxopenapi.COMPONENT_ACCESS_TOKEN, componentAccessToken, time.Now().Unix(), 3600)
}
