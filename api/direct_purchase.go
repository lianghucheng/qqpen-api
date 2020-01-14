package api

type BuyGoodsMCookie struct {
	BaseCookie
	OrgLoc string//需要填写: /v3/r/mpay/buy_goods_m（注意：如果经过接口机器，需要填写应用签名时使用的URI）
	Appip string
}

type BuyGoodsMReqParams struct {
	BaseRequestParam
	Payitem string //id1*price1*num1;id2*price2*num2)长度必须<512.道具直购物品信息在这里游戏自己定义自己管理.p*num*10 物品总额不能超过2000000000（分）
	Goodsmeta string //格式必须是“name*des”.长度必须<=256字符，必须使用utf8编码。目前goodsmeta超过76个字符后不能添加$ - _ .   ! * ' ( ) , ·等特殊字符.   格式中的“*”不要做编码
	GoodsUrl string //物品的图片url(长度<512字符) 注：参数已废弃直接传空
	Sig string
	Amt int64 //(可选)道具总价格.单位为1Q点）
	MaxNum int //(可选)
	Appmode int //(可选)1表示用户不可以修改物品数量，2 表示用户可以选择购买物品的数量。默  认2（注：批量购买的时候，必须等于1）
	AppMetadata string //（可选）发货时透传给应用。长度必须<=128字符
	Userip string //(可选)
	Format string //(可选)json、jsonp_$func。默认json。如果jsonp
	OutTradeNo string //订单号 只有走订单中心情况下需要传 其他情况不传
	SpInfo string //传code_id,batch_id，道具模式下单后台传抵扣券ID
}

type BuyGoodsMRespData struct{
	Ret int
	Msg string
	Token string
	UrlParams string
}

/**
0: 成功
1: 系统繁忙
2: token已过期
3: token不存在
4: 请求参数错误：（这里填写错误的具体参数）
 */

func (ctx *BuyGoodsMReqParams) DoRequest(cookie BuyGoodsMCookie) (respData BuyGoodsMRespData, err error) {

	return
}
