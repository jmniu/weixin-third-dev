package wxopenapi

import (
	"encoding/json"
	"fmt"
	"github.com/jmniu/wx-third-dev/pkg/logger"
	"github.com/parnurzeal/gorequest"
	"github.com/wendal/errors"
	"strings"
	"time"
)

const (
	MATERIAL_TYPE_IMAGE = "image"
	MATERIAL_TYPE_VIDEO = "video"
	MATERIAL_TYPE_VOICE = "voice"
	MATERIAL_TYPE_NEWS  = "news"
)

const (
	URL_MATERIAL_BATCH_GET   = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%v"
	URL_UPLOAD_PICTURE       = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%v"
	URL_MATERIAL_ADD_NEWS    = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%v"     //新增永久图文素材
	URL_MATERIAL_UPDATE_NEWS = "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%v"  //更新永久图文素材
	URL_MATERIAL_ADD         = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%v" //新增其他类型的素材
	URL_MATERIAL_DELETE      = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%v"
)

// 新增永久素材
// @return 素材mediaID
func (this *WxOpen) MaterialAddNews(accessToken string, articles ReqMaterialPicArticle) (mediaID string) {
	if len(accessToken) == 0 {
		logger.Warnf("MaterialAddNews empty accessToken")
		return ""
	}

	var url = fmt.Sprintf(URL_MATERIAL_ADD_NEWS, accessToken)
	var reqBody, _ = json.Marshal(articles)
	req := gorequest.New()
	_, resBody, errs := req.Post(url).SendString(string(reqBody)).End()
	logger.Infof("MaterialAddNews url[%s] reqBody[%s] res[%s]", url, reqBody, resBody)
	if errs != nil {
		logger.Warnf("MaterialAddNews error. errs[%v]", errs)
		return ""
	}

	var ret = struct {
		MediaID string `json:"media_id"`
	}{}
	_ = json.Unmarshal([]byte(resBody), &ret)
	return ret.MediaID

}

//更新永久素材
func (this *WxOpen) MaterialUpdateNews(accessToken string, mediaID string, index int, article ReqMaterialPicArticleItem) bool {
	if len(accessToken) == 0 {
		logger.Warnf("MaterialUpdateNews empty accessToken")
		return false
	}
	var url = fmt.Sprintf(URL_MATERIAL_UPDATE_NEWS, accessToken)
	var reqBody, _ = json.Marshal(article)
	req := gorequest.New()
	_, resBody, errs := req.Post(url).SendString(string(reqBody)).End()
	if errs != nil {
		logger.Warnf("MaterialUpdateNews error. errs[%v]", errs)
		return false
	}

	var ret struct {
		ErrCode int `json:"errcode"`
	}
	json.Unmarshal([]byte(resBody), &ret)
	if ret.ErrCode == 0 {
		return true
	} else {
		return false
	}
}

func (this *WxOpen) MaterialAdd(accessToken string, typ string, url string) (mediaID string, mediaURL string) {
	if len(accessToken) == 0 {
		logger.Warnf("MaterialAdd empty accessToken")
		return "", ""
	}
	//下载url文件到本地
	req := gorequest.New()
	_, body, errs := req.Timeout(time.Second * 10).Get(url).End()
	if errs != nil {
		logger.Warnf("MaterialAdd download error. errs[%v]", errs)
		return "", ""
	}

	var uploadURL = fmt.Sprintf(URL_MATERIAL_ADD, accessToken)
	_, resBody, errs := req.Post(uploadURL).Type("multipart").SendFile([]byte(body), "my.jpeg", "media").End() //todo my.jpeg可能会出现问题
	logger.Infof("MaterialAdd upload to weixin url[%s] res[%s]", uploadURL, resBody)
	if errs != nil {
		logger.Warnf("MaterialAdd upload file error. errs[%v]", errs)
		return "", ""
	}

	var ret = struct {
		MediaID string `json:"media_id"`
		URL     string `json:"url"`
	}{}
	_ = json.Unmarshal([]byte(resBody), &ret)

	mediaID = ret.MediaID
	mediaURL = ret.URL
	return
}

func (this *WxOpen) GetMaterial(access_token string, req ReqMaterial) (rspobj RspMaterial, err error) {
	fmt.Println("GetMaterial ", time.Now().Unix())
	if len(access_token) > 0 {
		fmt.Println("auth_appid is ", access_token)
		reqstr, _ := json.Marshal(req)
		var rsp []byte
		rsp, err = PostJsonByte(fmt.Sprintf(URL_MATERIAL_BATCH_GET, access_token), reqstr)
		if err != nil {
			fmt.Println("请求 Material 失败 err", err.Error())
			return
		}
		if strings.Contains(string(rsp), "errcode") {
			fmt.Println("请求 Material 失败 rsp", string(rsp))
			if strings.Contains(string(rsp), "access_token expired") {
				err = errors.New("口令过期")
			} else {
				err = errors.New(string(rsp))
			}
			return
		}
		// var rspobj RspMaterial
		json.Unmarshal(rsp, &rspobj)
		// rspobj.ComponentAppid = SAPPID
		fmt.Println(rspobj)
		return
	} else {
		fmt.Println("access_token is empty")
	}
	return
}

func (this *WxOpen) UploadPicture(accessToken string, url string) string {
	if len(accessToken) == 0 {
		logger.Warnf("UploadPicture empty accessToken")
		return ""
	}
	//下载url文件到本地
	req := gorequest.New()
	_, body, errs := req.Timeout(time.Second * 10).Get(url).End()
	if errs != nil {
		logger.Warnf("UploadPicture download error. errs[%v]", errs)
		return ""
	}

	var uploadURL = fmt.Sprintf(URL_UPLOAD_PICTURE, accessToken)
	_, body, errs = req.Post(uploadURL).Type("multipart").SendFile([]byte(body), "fw.jpeg", "media").End()
	logger.Infof("UploadPicture upload to weixin url[%s] res[%s]", uploadURL, body)
	if errs != nil {
		logger.Warnf("UploadPicture upload file error. errs[%v]", errs)
		return ""
	}

	var ret = struct {
		URL string `json:"url"`
	}{}
	json.Unmarshal([]byte(body), &ret)

	return ret.URL
}

func (this *WxOpen) DeleteMaterial(accessToken string, mediaID string) bool {
	if len(accessToken) == 0 {
		logger.Warnf("DeleteMaterial empty accessToken")
		return false
	}

	var delURL = fmt.Sprintf(URL_MATERIAL_DELETE, accessToken)
	var postBody = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaID,
	}
	req := gorequest.New()
	_, resBody, errs := req.Post(delURL).SendStruct(postBody).End()
	if errs != nil {
		logger.Warnf("DeleteMaterial error. url[%s] errs[%v]", delURL, errs)
		return false
	}
	var res struct {
		ErrCode int `json:"errcode"`
	}
	json.Unmarshal([]byte(resBody), &res)
	if res.ErrCode == 0 {
		return true
	} else {
		return false
	}
}
