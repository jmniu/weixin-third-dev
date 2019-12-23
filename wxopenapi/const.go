package wxopenapi

const (
	COMPONENT_VERIFY_TICKET = "component_verify_ticket"
	COMPONENT_ACCESS_TOKEN  = "component_access_token"
	PRE_AUTH_CODE           = "pre_auth_code"
	AUTHORIZED              = "authorized"
	UNAUTHORIZED            = "unauthorized"
	UPDATEAUTHORIZED        = "updateauthorized"
)
const (
	FMT_URL_PRE_AUTH_CODE       = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%v&pre_auth_code=%v&redirect_uri=%v&auth_type=%v"
	FMT_URL_LOCAL_PRE_AUTH_CODE = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%v&pre_auth_code=%v&redirect_uri=%v&auth_type=%v&biz_appid=xxxx#wechat_redirect"
)

const (
	AUTH_TYPE_GZH = 1 //1则商户点击链接后，手机端仅展示公众号
	AUTH_TYPE_XCX = 2 //2表示仅展示小程序
	AUTH_TYPE_ALL = 3 //3表示公众号和小程序都展示
)

const (
	// 第三方平台component_access_token是第三方平台的下文中接口的调用凭据，也叫做令牌（component_access_token）。
	// 每个令牌是存在有效期（2小时）的，且令牌的调用不是无限制的，请第三方平台做好令牌的管理，在令牌快过期时（比如1小时50分）再进行刷新。
	URL_COMPONENT_ACCESS_TOKEN = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	// 该API用于获取预授权码。预授权码用于公众号或小程序授权时的第三方平台方安全验证。
	URL_PRE_AUTH_CODE = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%v"
	// 该API用于使用授权码换取授权公众号或小程序的授权信息，并换取authorizer_access_token和authorizer_refresh_token。
	// 授权码的获取，需要在用户在第三方平台授权页中完成授权流程后，在回调URI中通过URL参数提供给第三方平台方。
	// 请注意，由于现在公众号或小程序可以自定义选择部分权限授权给第三方平台，
	// 因此第三方平台开发者需要通过该接口来获取公众号或小程序具体授权了哪些权限，而不是简单地认为自己声明的权限就是公众号或小程序授权的权限。
	URL_AUTHORIZER_ACCESS = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%v"
	// 该API用于在授权方令牌（authorizer_access_token）失效时，可用刷新令牌（authorizer_refresh_token）获取新的令牌。
	// 请注意，此处token是2小时刷新一次，开发者需要自行进行token的缓存，避免token的获取次数达到每日的限定额度。
	// 缓存方法可以参考：http://mp.weixin.qq.com/wiki/2/88b2bf1265a707c031e51f26ca5e6512.html
	// 当换取authorizer_refresh_token后建议保存。
	URL_AUTHORIZER_ACCESS_REFRESH = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%v"
	// 该API用于获取授权方的基本信息，包括头像、昵称、帐号类型、认证类型、微信号、原始ID和二维码图片URL。
	// 需要特别记录授权方的帐号类型，在消息及事件推送时，对于不具备客服接口的公众号，需要在5秒内立即响应；
	// 而若有客服接口，则可以选择暂时不响应，而选择后续通过客服接口来发送消息触达粉丝。
	//（1）公众号获取方法如下：
	URL_AUTHORIZER_INFO = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%v"
	//（2）小程序获取方法如下
	URL_AUTHORIZER_INFO_MINI = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%v"
	// 该API用于获取授权方的公众号或小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开关。
	// 注意，获取各项选项设置信息，需要有授权方的授权，详见权限集说明。
	URL_AUTHORIZER_OPTION = "https://api.weixin.qq.com/cgi-bin/component/ api_get_authorizer_option?component_access_token=%v"
	// 该API用于设置授权方的公众号或小程序的选项信息，如：地理位置上报，语音识别开关，多客服开关。
	// 注意，设置各项选项设置信息，需要有授权方的授权，详见权限集说明。
	URL_SET_AUTHORIZER_OPTION = "https://api.weixin.qq.com/cgi-bin/component/ api_set_authorizer_option?component_access_token=%v"
)
