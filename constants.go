package main

import "errors"

const youdaoName = "有道"
const baiduName = "百度"

const baidubURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
const youdaoURL = "https://openapi.youdao.com/api"

const configfile = ".dict_config"

var tableTitle = []interface{}{"来源", "原文", "译文", "音标", "详情"}

var baiduErrCodeMessage = map[string]error{
	"52001": errors.New("请求超时, 请稍后重试"),
	"52002": errors.New("系统错误, 请稍后重试"),
	"52003": errors.New("未授权用户, 检查您的 appid 是否正确，或者服务是否开通"),
	"54000": errors.New("必填参数为空, 检查是否少传参数"),
	"54001": errors.New("签名错误, 请检查您的签名生成方法"),
	"54003": errors.New("访问频率受限, 请降低您的调用频率"),
	"54004": errors.New("账户余额不足, 请前往管理控制台为账户充值"),
	"54005": errors.New("长query请求频繁, 请降低长query的发送频率，3s后再试"),
	"58000": errors.New("客户端IP非法"),
}

// format by jscode ->
// text.split('\n').map(x => x.split("\t")).map(s => `"${s[0]}": errors.New("${s[1]}")`).join(",\n")
var youdaoErrCodeMessage = map[string]error{
	"101": errors.New("缺少必填的参数"),
	"102": errors.New("不支持的语言类型"),
	"103": errors.New("翻译文本过长"),
	"104": errors.New("不支持的API类型"),
	"105": errors.New("不支持的签名类型"),
	"106": errors.New("不支持的响应类型"),
	"107": errors.New("不支持的传输加密类型"),
	"108": errors.New("应用ID无效，注册账号，登录后台创建应用和实例并完成绑定，可获得应用ID和应用密钥等信息"),
	"109": errors.New("batchLog格式不正确"),
	"110": errors.New("无相关服务的有效实例"),
	"111": errors.New("开发者账号无效"),
	"112": errors.New("请求服务无效"),
	"113": errors.New("q不能为空"),
	"114": errors.New("不支持的图片传输方式"),
	"201": errors.New("解密失败，可能为DES,BASE64,URLDecode的错误"),
	"202": errors.New("签名检验失败"),
	"203": errors.New("访问IP地址不在可访问IP列表"),
	"205": errors.New("请求的接口与应用的平台类型不一致，如有疑问请参考入门指南"),
	"206": errors.New("因为时间戳无效导致签名校验失败"),
	"207": errors.New("重放请求"),
	"301": errors.New("辞典查询失败"),
	"302": errors.New("翻译查询失败"),
	"303": errors.New("服务端的其它异常"),
	"304": errors.New("会话闲置太久超时"),
	"401": errors.New("账户已经欠费停"),
	"402": errors.New("offlinesdk不可用"),
	"411": errors.New("访问频率受限,请稍后访问"),
	"412": errors.New("长请求过于频繁，请稍后访问"),
}

var errFileIsRequired = errors.New("You need a config file specified by -f flag")
