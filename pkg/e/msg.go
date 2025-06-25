package e

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "failed",
	ERROR_BIND:     "接口负载绑定数据结构体失败",
	ERROR_VALID:    "验证接口参数失败",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_TOKEN: "生成token失败",
	ERROR_AUTH_VALUE: "获取token中携带的数据失败",

	SUCCESS_USER_CREATE: "用户创建成功",
	SUCCESS_USER_UPDATE: "更新用户数据成功",
	ERROR_USER_AUTH:     "用户名或密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
