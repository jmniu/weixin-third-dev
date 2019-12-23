package wxopenapi

import (
	"encoding/json"
	"fmt"
	"github.com/jmniu/wx-third-dev/pkg/logger"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_GET_WEBAUTH_ACCESS_TOKEN = "https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=%s&code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
)

type UserAuthInfo struct {
	AccessToken  string
	ExpiresIN    int
	RefreshToken string
	OpenID       string
	Scope        string
}

func (this *WxOpen) GetUserAuthInfo(appID string, code string, componentAppID string) *UserAuthInfo {
	var componentAccessToken = this.GetInfo(COMPONENT_ACCESS_TOKEN).Info
	var url = fmt.Sprintf(URL_GET_WEBAUTH_ACCESS_TOKEN, appID, code, componentAppID, componentAccessToken)
	request := gorequest.New()
	_, body, errs := request.Get(url).End()
	logger.Infof("GetAccessTokenByCode reqURL[%s] resBody[%s]", url, body)
	if errs != nil {
		logger.Warnf("GetAccessTokenByCode fail.")
		return nil
	}

	var authInfo UserAuthInfo
	json.Unmarshal([]byte(body), &authInfo)

	return &authInfo
}
