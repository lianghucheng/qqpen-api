package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

func GenericSign(method string, theUrl string, params []string) (sign string, err error) {

	//源串构造步骤
	//第1步：将请求的URI路径进行URL编码（URI不含host，URI示例：/v3/r/mpay/get_balance_m，）。
	//请开发者关注：URL编码注意事项，否则容易导致后面签名不能通过验证。
	parseUrl, err := url.Parse(theUrl)
	if err != nil {
		return
	}
	prefix := "/v3/r"
	uri := prefix + parseUrl.RequestURI()
	theUrlEncode := url.QueryEscape(uri)

	//第2步： 将除“sig”外的所有参数按key进行字典升序排列。
	//注：除非OpenAPI文档中特别标注了某参数不参与签名，否则除sig外的所有参数都要参与签名。
	sort.Strings(params)

	//第3步：将第2步中排序后的参数(key=value)用&拼接起来，并进行URL编码。
	//请开发者关注：URL编码注意事项，否则容易导致后面签名不能通过验证。
	paramsEncode := url.QueryEscape(strings.Join(params,"&"))

	//第4步：将HTTP请求方式（GET或者POST）以及第1步和第3步中的字符串用&拼接起来。
	//注：Java_SDK_V3.0.6仅支持POST方式，如果用GET可能导致一直计算sig不正确。

	encode := method + "&" + theUrlEncode + "&" + paramsEncode

	appkey := ""
	for _,v := range params {
		if strings.Contains(v, "appkey") {
			appkey = strings.Split(v, "=")[1]
			break
		}
	}

	mac := hmac.New(sha1.New, []byte(appkey+"&"))
	mac.Write([]byte(encode))
	sign = url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return
}
