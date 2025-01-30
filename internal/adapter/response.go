package adapter

func Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": data,
	}
}

func Failed(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
}
