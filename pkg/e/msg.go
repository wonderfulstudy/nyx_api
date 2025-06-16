package e

var MsgFlags = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "failed",
	ERROR_BIND_JSON: "接口负载绑定数据结构体失败",
	INVALID_PARAMS:  "请求参数错误",

	ERROR_AUTH_TOKEN: "生成token失败",

	ERROR_USER_AUTH: "用户名或密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
