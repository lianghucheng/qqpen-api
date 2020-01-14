package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"qqopen-api/api"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("",handleBuyGoodsM)
	
	err := http.ListenAndServe(":18080", router)
	if err != nil {
		panic(err)
	}
	fmt.Println("服务启动成功！")
}

func handleBuyGoodsM(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	cookies := r.Cookies()
	reqParam := &api.BuyGoodsMReqParams{
		BaseRequestParam: api.BaseRequestParam{
			Offerid:     "",
			Openid:      "",
			Openkey:     "",
			Ts:          0,
			Pf:          "",
			Zoneid:      "",
			Pfkey:       "",
			Accounttype: "",
			Billno:      "",
		},
		Payitem:          "",
		Goodsmeta:        "",
		GoodsUrl:         "",
		Sig:              "",
		Amt:              0,
		MaxNum:           0,
		Appmode:          0,
		AppMetadata:      "",
		Userip:           "",
		Format:           "",
		OutTradeNo:       "",
		SpInfo:           "",
	}


	respData ,err := reqParam.DoRequest(api.BuyGoodsMCookie{
		BaseCookie: api.BaseCookie{},
		OrgLoc:     "",
		Appip:      "",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(method, cookies, respData)
}
