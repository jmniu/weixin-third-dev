package wxopenapi

import (
	"encoding/json"
	"fmt"
	"github.com/jmniu/wx-third-dev/pkg/logger"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_GET_USER_LIST = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%v&next_openid=%v"
)

func (this *WxOpen) GetAllUserOpenIDList(accessToken string) []string {
	var openIDList []string
	var nextOpenID string

	for {
		var userGetRes struct {
			Total int `json:"total"`
			Count int `json:"count"`
			Data  struct {
				OpenID []string `json:"openid"`
			} `json:"data"`
			NextOpenID string `json:"next_openid"`
		}
		var url = fmt.Sprintf(URL_GET_USER_LIST, accessToken, nextOpenID)
		req := gorequest.New()
		_, resBody, errs := req.Get(url).End()
		logger.Infof("GetAllUserOpenIDList url[%s] res[%s]", url, resBody)
		if errs != nil {
			break
		}

		json.Unmarshal([]byte(resBody), &userGetRes)

		if userGetRes.Data.OpenID == nil {
			break
		}

		openIDList = append(openIDList, userGetRes.Data.OpenID...)

		nextOpenID = userGetRes.NextOpenID
	}

	return openIDList
}
