package wxopenapi

import (
	"encoding/json"
	"fmt"
	"github.com/jmniu/weixin-third-dev/pkg/logger"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_MESSAGE_MASS_SENDALL  = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=%v"
	URL_MESSAGE_TEMPLATE_SEND = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%v"
	URL_MESSAGE_TEMPLATE_LIST = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%v"
	URL_MESSAGE_TEMPLATE_ADD  = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%v"
	URL_MESSAGE_TEMPLATE_DEL  = "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%v"
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

type MessageTemplateListItem struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type MessageTemplateListReturn struct {
	TemplateList []MessageTemplateListItem `json:"template_list"`
}

func (this *WxOpen) MessageTemplateList(accessToken string) MessageTemplateListReturn {
	url := fmt.Sprintf(URL_MESSAGE_TEMPLATE_LIST, accessToken)
	req := gorequest.New()
	_, resBody, _ := req.Get(url).End()
	logger.Infof("MessageTemplateList url[%s] resBody[%v]", url, resBody)

	var ret MessageTemplateListReturn
	json.Unmarshal([]byte(resBody), &ret)

	return ret
}

type MessageTemplateAddReturn struct {
	ErrCode    int    `json:"err_code"`
	ErrMsg     string `json:"err_msg"`
	TemplateID string `json:"template_id"`
}

func (this *WxOpen) MessageTemplateAdd(accessToken string, shortID string) MessageTemplateAddReturn {
	url := fmt.Sprintf(URL_MESSAGE_TEMPLATE_ADD, accessToken)
	reqBody, _ := json.Marshal(map[string]string{
		"template_id_short": shortID,
	})
	req := gorequest.New()
	_, resBody, _ := req.Post(url).SendString(string(reqBody)).End()
	logger.Infof("MessageTemplateAdd url[%s] reqBody[%v] resBody[%v]", url, string(reqBody), resBody)

	var ret MessageTemplateAddReturn
	_ = json.Unmarshal([]byte(resBody), &ret)
	return ret
}

func (this *WxOpen) MessageTemplateDelete(accessToken string, templateID string) bool {
	url := fmt.Sprintf(URL_MESSAGE_TEMPLATE_DEL, accessToken)
	reqBody, _ := json.Marshal(map[string]string{
		"template_id": templateID,
	})
	req := gorequest.New()
	_, resBody, _ := req.Post(url).SendString(string(reqBody)).End()
	logger.Infof("MessageTemplateDelete url[%s] reqBody[%v] resBody[%v]", url, string(reqBody), resBody)

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
