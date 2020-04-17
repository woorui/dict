package main

const baidubURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
const youdaoURL = "https://openapi.youdao.com/api"

var errCodeMessage = map[string]string{
	"52001": "请求超时, 请稍后重试",
	"52002": "系统错误, 请稍后重试",
	"52003": "未授权用户, 检查您的 appid 是否正确，或者服务是否开通",
	"54000": "必填参数为空, 检查是否少传参数",
	"54001": "签名错误, 请检查您的签名生成方法",
	"54003": "访问频率受限, 请降低您的调用频率",
	"54004": "账户余额不足, 请前往管理控制台为账户充值",
	"54005": "长query请求频繁, 请降低长query的发送频率，3s后再试",
	"58000": "客户端IP非法",
}
