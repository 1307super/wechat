package officialaccount

import (
	"github.com/silenceper/wechat/officialaccount"
	offConfig "github.com/silenceper/wechat/officialaccount/config"
	offContext "github.com/silenceper/wechat/officialaccount/context"
	opContext "github.com/silenceper/wechat/openplatform/context"
)

type OfficialAccount struct {
	//授权的公众号的appID
	appID string
	*officialaccount.OfficialAccount
	opContext *opContext.Context
}

//NewOfficialAccount 实例化
//appID :为授权方公众号 APPID，非开放平台第三方平台 APPID
func NewOfficialAccount(opCtx *opContext.Context, appID string) *OfficialAccount {
	officialAccount := officialaccount.NewOfficialAccount(&offConfig.Config{
		AppID:          opCtx.AppID,
		EncodingAESKey: opCtx.EncodingAESKey,
		Token:          opCtx.Token,
		Cache:          opCtx.Cache,
	})
	//设置获取access_token的函数
	officialAccount.GetContext().SetGetAccessTokenFunc(func(offCtx *offContext.Context) (accessToken string, err error) {
		// 获取授权方的access_token
		return opCtx.GetAuthrAccessToken(appID)
	})
	return &OfficialAccount{appID: appID, OfficialAccount: officialAccount}
}
