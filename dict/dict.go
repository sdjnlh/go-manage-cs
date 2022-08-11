package dict

// 微信接口
const (
	WxLogin  = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WxUnid   = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
	WXreqUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
)

const (

	// Ebl小程序
	LibrarayId                     = "wx98bbdb4a7d2c92f3"
	LibrarySecret                  = "bcb29d9aceddd5e9304222627d517769"
	SystemAddress                  = "3005"       //默认端口
	SystemTitle                    = "金宝塑业后台管理系统" //默认标题
	SystemAesKey                   = "dataon$123" //默认密钥
	DefaultPwd                     = "123456"     //用户默认密码
	SystemSuccessCode              = "200"
	SystemFailureCode              = "10002"
	SystemSuccessMsg               = "请求成功"
	SystemFailureMsg               = "请求失败"
	ERROR_AUTH_NO_TOKRN            = "5001" //非法token
	ERROR_AUTH_CHECK_TOKEN_FAIL    = "5002" //token验证失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = "5004" //token过期
)
