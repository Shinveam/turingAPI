package turing

import (
	"github.com/bitly/go-simplejson"
	"github.com/imroc/req"
	"net"
)

//接口地址
const URL = "http://openapi.tuling123.com/openapi/api/v2"

//发送的参数内容
type TuringParams struct {
	ReqType    int         `json:"repType"`    //输入类型，0-文本（默认），1-图片，2-视频，无需包含
	Perception *Perception `json:"perception"` //输入信息，必须包含
	UserInfo   *UserInfo   `json:"userInfo"`   //用户参数，必须包含
}

//参数组成部分
type Perception struct {
	InputText  *InputText  `json:"inputText"`  //文本信息，无需包含
	InputImage *InputImage `json:"inputImage"` //图片信息，无需包含
	InputMedia *InputMedia `json:"inputMedia"` //音频信息，无需包含
	SelfInfo   *SelfInfo   `json:"selfInfo"`   //客户端属性，无需包含
}

//perception的组成部分
type InputText struct {
	Text string `json:"text"` //直接输入文本，取值范围1-128字符，必须包含
}
type InputImage struct {
	Url string `json:"url"` //图片地址，必须包含
}
type InputMedia struct {
	Url string `json:"url"` //音频地址，必须包含
}
type SelfInfo struct {
	Location *Location `json:"location"` //地理位置信息，无需包含
}

type Location struct {
	City     string `json:"city"`     //所在城市，必须包含
	Province string `json:"province"` //所在省份，无需包含
	Street   string `json:"street"`   //所在街道，无需包含
}

//userInfo组成部分
type UserInfo struct {
	ApiKey     string `json:"apiKey"`     //机器人标识，32位，必须包含
	UserId     string `json:"userId"`     //用户唯一标识，长度小于等于32位，必须包含
	GroupId    string `json:"groupId"`    //群聊唯一标识，长度小于等于64位，无需包含
	UserIdName string `json:"userIdName"` //群内用户昵称，长度小于等于64位，无需包含
}

//消息接受参数
type TuringResponse struct {
	Intent  *intent   `json:"intent"`  //请求意图，必须包含
	Results []Results `json:"results"` //输出结果集， 无需包含
}

//intent组成部分
type intent struct {
	Code       int    `json:"code"`       //输出功能code，必须包含
	IntentName string `json:"intent"`     //意图名称，无需包含
	ActionName string `json:"actionName"` //意图动作名称。无需包含
}

//results组成部分
type Results struct {
	ResultType string            `json:"resultType"` //返回类型，必须包含
	Values     map[string]string `json:"values"`     //返回值（返回多种数据类型值），必须包含  -->学术不精，此处数值类型定义不太清
	GroupType  int               `json:"groupType"`  //组编号，0为独立输出，大于0时包含同组相关内容，必须包含
}

type TuringParam func(params *TuringParams)

func ReqType(req_type int) TuringParam {
	if req_type != 0 && req_type != 1 && req_type != 2 {
		req_type = 0
	}
	return func(params *TuringParams) {
		params.ReqType = req_type
	}
}

//主体部分
func Robots(ApiKey string, req_type TuringParam, contents string) (interface{}, error) {

	var cuid string
	//获取本机MAC地址
	MACAddr, err := net.Interfaces()
	if err != nil {
		cuid = "demo"
	} else {
		for _, itf := range MACAddr {
			if cuid = itf.HardwareAddr.String(); len(cuid) > 0 {
				break
			}
		}
	}

	turingParam := &TuringParams{
		ReqType: 0,

		Perception: &Perception{
			InputText: &InputText{
				Text: contents,
			},
			//InputImage: &InputImage{
			//	Url: contents,
			//},
			//InputMedia: &InputMedia{
			//	Url: contents,
			//},

			//SelfInfo: &SelfInfo{
			//	Location: &Location{
			//		City:     "北京",
			//		Province: "北京",
			//		Street:   "信息路",
			//	},
			//},
		},

		UserInfo: &UserInfo{
			ApiKey: ApiKey,
			UserId: "demo",
		},
	}

	//fmt.Println(turingParam)

	rt := req_type
	rt(turingParam)

	header := req.Header{
		"Content-Type": "application/json",
	}
	//fmt.Println(req.BodyJSON(turingParam))
	resp, err := req.Post(URL, header, req.BodyJSON(turingParam))
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp.String())

	//处理接受的json数据--simplejson的使用
	ret, err := simplejson.NewJson([]byte(resp.String()))
	if err != nil {
		return nil, err
	}

	retArray, err := ret.Get("results").Array()
	if err != nil {
		return nil, err
	}

	//var values string
	for i, _ := range retArray {
		retMsg := ret.Get("results").GetIndex(i)
		//fmt.Println("retMsg:",retMsg.MustMap())

		values := retMsg.Get("values").MustMap()
		for _, v := range values {
			//fmt.Println(values[k])
			//fmt.Println(reflect.TypeOf(values[k]))
			//fmt.Println(v)
			return v, err
		}
	}

	return nil, err
}
