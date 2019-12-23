package wxopenapi

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"wx-third-dev/pkg/logger"
)

const (
	URL_MESSAGE_MASS_SENDALL  = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=%v"
	URL_MESSAGE_TEMPLATE_SEND = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%v"
)

type MessageMassSend struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagID   int  `json:"tag_id"`
	} `json:"filter"`
	MPNews struct {
		MediaID string `json:"media_id"`
	} `json:"mpnews"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
	MsgType           string `json:"msgtype"`
	SendIgnoreReprint int    `json:"send_ignore_reprint"`
}

//图文消息的发送
func (this *WxOpen) MessageMassSendAllOfMPNews(accessToken string, mediaID string) bool {
	if len(accessToken) == 0 {
		logger.Warnf("MessageMassSendAllOfMPNews empty accessToken")
		return false
	}

	var postBody = MessageMassSend{
		Filter: struct {
			IsToAll bool `json:"is_to_all"`
			TagID   int  `json:"tag_id"`
		}{
			IsToAll: true,
		},
		MPNews: struct {
			MediaID string `json:"media_id"`
		}{
			MediaID: mediaID,
		},
		Text: struct {
			Content string `json:"content"`
		}{},
		Voice: struct {
			MediaID string `json:"media_id"`
		}{},
		MsgType:           "mpnews",
		SendIgnoreReprint: 0,
	}
	url := fmt.Sprintf(URL_MESSAGE_MASS_SENDALL, accessToken)
	req := gorequest.New()
	_, resBody, _ := req.Post(url).SendStruct(postBody).End()
	logger.Infof("MessageMassSendAllOfMPNews url[%s] reqBody[%v] resBody[%v]", url, postBody, resBody)

	var ret struct {
		ErrCode int `json:"errcode"`
	}
	_ = json.Unmarshal([]byte(resBody), &ret)
	if ret.ErrCode > 0 {
		return false
	} else {
		return true
	}
}

type MessageTemplate struct {
	ToUser      string `json:"touser"`
	TemplateID  string `json:"template_id"`
	URL         string `json:"url"`
	MiniProgram struct {
		AppID    string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
	Data struct {
		First struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"first"`
		Keyword1 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword1"`
		Keyword2 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword2"`
		Keyword3 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword3"`
		Keyword4 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword4"`
		Keyword5 struct {
			Value string `json:"value"`
			Color string
		} `json:"keyword5"`
		Remark struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"remark"`
	} `json:"data"`
}

//模板消息发送
func (this *WxOpen) MessageTemplateSend(accessToken string, msg MessageTemplate) bool {
	if len(accessToken) == 0 {
		logger.Warnf("MessageTemplateSend empty accessToken")
		return false
	}

	url := fmt.Sprintf(URL_MESSAGE_TEMPLATE_SEND, accessToken)
	reqBody, _ := json.Marshal(msg)
	req := gorequest.New()
	_, resBody, _ := req.Post(url).SendString(string(reqBody)).End()
	logger.Infof("MessageTemplateSend url[%s] reqBody[%v] resBody[%v]", url, string(reqBody), resBody)

	var ret struct {
		ErrCode int `json:"errcode"`
	}
	_ = json.Unmarshal([]byte(resBody), &ret)
	if ret.ErrCode > 0 {
		return false
	} else {
		return true
	}
}
