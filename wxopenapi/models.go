package wxopenapi

type EncryptMsgQuery struct {
	MsgSignature string `form:"msg_signature"`
	Timestamp    string `form:"timestamp"`
	Nonce        string `form:"nonce"`
}

type EncryptMsgBody struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
}

type OpenToken struct {
	Typ       string
	CreatedAt int64
	ExpiredAt int64
	ExpiredIn int
	Info      string
}

//1.com ticket
//4.auth code
type NotifyAuth struct {
	AppId                        string
	CreateTime                   string
	InfoType                     string
	ComponentVerifyTicket        string
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

//2.com access token
type ReqAccessToken struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}
type RspAccessToken struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

//3.com preauthcode
type ReqPreAuthCode struct {
	ComponentAppid string `json:"component_appid"`
}
type RspPreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

//5.auth access token
//5-1
type ReqAuthAccessToken struct {
	ComponentAppid    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}
type FuncscopeCategory struct {
	Id int `json:"id"`
}
type FuncInfo struct {
	FuncCat FuncscopeCategory `json:"funcscope_category"`
}
type AuthAccessTokenInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfos              []FuncInfo `json:"func_info"`
}
type RspAuthAccessToken struct {
	AuthorizationInfo AuthAccessTokenInfo `json:"authorization_info"`
}
type AuthAccessToken struct {
	ComponentAppid         string
	AuthorizerAppid        string
	AuthorizerAccessToken  string
	ExpiresIn              int
	AuthorizerRefreshToken string
	FuncInfos              []FuncInfo
}

//5-2
type ReqUpdateAuthAccessToken struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
type RspUpdateAuthAccessToken struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type UpdateAuthAccessToken struct {
	ComponentAppid         string
	AuthorizerAppid        string
	AuthorizerAccessToken  string
	ExpiresIn              int
	AuthorizerRefreshToken string
}

//6.auth info
type ReqAuthInfo struct {
	ComponentAppid  string `json:"component_appid"`
	AuthorizerAppid string `json:"authorizer_appid"`
}
type TypeInfo struct {
	Id int `json:"id"`
}
type BusiInfo struct {
	OpenPay   int `json:"open_pay"`
	OpenShake int `json:"open_shake"`
	OpenScan  int `json:"open_scan"`
	OpenCard  int `json:"open_card"`
	OpenStore int `json:"open_store"`
}
type RspAuthorizerInfo struct {
	NickName        string   `json:"nick_name"`
	HeadImg         string   `json:"head_img"`
	ServiceTypeInfo TypeInfo `json:"service_type_info"`
	VerifyTypeInfo  TypeInfo `json:"verify_type_info"`
	UserName        string   `json:"user_name"`
	Alias           string   `json:"alias"`
	QrcodeUrl       string   `json:"qrcode_url"`
	BusinessInfo    BusiInfo `json:"business_info"`
	Idc             int      `json:"idc"`
	PrincipalName   string   `json:"principal_name"`
	Signature       string   `json:"signature"`
}
type RspAuthorizationInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfos              []FuncInfo `json:"func_info"`
}
type RspAuthInfo struct {
	ComponentAppid    string
	AuthorizerInfo    RspAuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo RspAuthorizationInfo `json:"authorization_info"`
}

//7.material
type ReqMaterial struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}
type RspNewsItem struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	Url                string `json:"url"`
	ThumbUrl           string `json:"thumb_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}
type RspNewsContent struct {
	NewsItem   []RspNewsItem `json:"news_item"`
	CreateTime int           `json:"create_time"`
	UpdateTime int           `json:"update_time"`
}
type RspMaterialItem struct {
	MediaId    string         `json:"media_id"`
	Name       string         `json:"name"`
	UpdateTime int            `json:"update_time"`
	Url        string         `json:"url"`
	Content    RspNewsContent `json:"content"`
}
type RspMaterial struct {
	Item       []RspMaterialItem `json:"item"`
	TotalCount int               `json:"total_count"`
	ItemCount  int               `json:"item_count"`
}

//新增图文素材
type ReqMaterialPicArticle struct {
	Articles []ReqMaterialPicArticleItem `json:"articles"`
}

//新增图文素材项
type ReqMaterialPicArticleItem struct {
	Title              string `json:"title"`
	ThumbMediaID       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}
