/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-go.
 */

package entity

// Request is a request struct.
type Request struct {
	Method       string `json:"method" description:"API接口名称"`
	AppKey       string `json:"app_key" description:"联盟分配给应用的appkey，可在应用查看中获取appkey"`
	AccessToken  string `json:"access_token,omitempty" description:"授权令牌 根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传"`
	Format       string `json:"format" description:"响应格式，暂时只支持json"`
	Timestamp    string `json:"timestamp" description:"timestamp 时间戳，格式为yyyy-MM-dd   HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟"`
	Version      string `json:"v" description:"1.0 API协议版本，请根据API具体版本号传入此参数，一般为1.0"`
	SignMethod   string `json:"sign_method" description:"签名方式，签名的摘要算法，暂时只支持md5"`
	Sign         string `json:"sign" description:"签名 API输入参数签名结果"`
	BuyParamJSON string `json:"360buy_param_json" description:"360buyParamJson API入参报文的json字符串，正常情况直接将JSON字符串作为参数值传入API"`
}
