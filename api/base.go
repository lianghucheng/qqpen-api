package api

type BaseCookie struct {
	SessionID string
	SessionType string
}

type BaseRequestParam struct {
	Offerid string
	Openid string
	Openkey string
	Ts int64
	Pf string //平台-注册渠道-版本-安装渠道-业务自定义(自定义)，最大100字符。
	Zoneid string //游戏服务器大区id,游戏不分大区则默认zoneId ="1",String类型。如果应用选择支持角色，则角色ID接在分区ID号后用"_"连接。
	Pfkey string //由平台来源和openkey根据规则生成的一个密钥串.自研和第三方登录的应用不校验，可以传递为pfKey="pfKey".非自研强校验,pfKey="58FCB2258B0BF818008382BD025E8022"（来自平台）
	Accounttype string //账户类型，分为基础货币和安全货币，默认为基础货币
	Billno string //订单号.数字和字符不限，不能包 含特殊字符如& = | % ^   等即可（有效期48小时）
}
