package util

import (
	"fmt"
	"testing"
)

func TestGenericSign(t *testing.T) {
	method := "GET"
	url := "http://localhost/mpay/get_balance_m"
	params := []string{
		`pfkey=CA641BC173479B8C0B35BC84873B3DB9`,
		`ts=1340880299`,
		`userip=112.90.139.30`,
		`zoneid=1`,
		`format=json&openid=00000000000000000000000014BDF6E4`,
		`openkey=AB43BF3DC5C3C79D358CC5318E41CF59`,
		`pf=myapp_m_qq-00000000-android-00000000-ysdk`,
		`appid=15499`,
	}

	sign, err := GenericSign(method, url, params)
	if err != nil {
		fmt.Println(err)
	}

	if `GET&%2Fv3%2Fr%2Fmpay%2Fget_balance_m&appid%3D15499%26format%3Djson%26openid%3D00000000000000000000000014BDF6E4%26openkey%3DAB43BF3DC5C3C79D358CC5318E41CF59%26pf%3Dmyapp_m_qq-00000000-android-00000000-ysdk%26pfkey%3DCA641BC173479B8C0B35BC84873B3DB9%26ts%3D1340880299%26userip%3D112.90.139.30%26zoneid%3D1` == sign {
		fmt.Println("pass")
	} else {
		fmt.Println("no")
	}
}
