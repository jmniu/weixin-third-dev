微信公众号第三方开发

注：
1. 本类库不关心token的刷新时间以及时机
2. 为防止token过期，每次InitWxOpen的时候传入最新的token

用法：
```go
package helloworld
var (
    componentToken          = "xxxxxxx"
    componentAppID          = "yyyyyyy"
    componentAppSecret      = "xxxx.xxx.xxx"
    componentVerifyTicket   = GetVerifyTicket()         //服务商票据，自己实现
    componentAccessToken    = GetComponentAccessToken() //服务商token，自己实现
    accessToken             = GetAccessToken()          //公众号token，自己实现
)

wxOpen := InitWxOpen(componentToken,
	componentAesKey,
	componentAppID,
	componentAppSecret,
	componentVerifyTicket,
	componentAccessToken)
ret := wxOpen.SetMenu(accessToken, menu)
fmt.Println(ret)
```