package e

var MsgFlags = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "fail",
	INVALID_PARAMS:   "请求参数错误",
	ACCESS_FORBIDDEN: "您没有权限访问该信息",

	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_PASSWD_NOT_MATCH: "用户名与密码不一致",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_USER_EXIST: "用户已存在",
	ERROR_USER_TYPE:  "用户角色错误",

	ERROR_CLASS_NOT_EXIST: "班级不存在",
	ERROR_CLASS_EXIST:     "班级已存在",

	ERROR_COURSE_NOT_EXIST: "课程不存在",
	ERROR_COURSE_EXIST:     "课程已存在",

	ERROR_STUDENT_COURSE_NOT_EXIST: "选课记录不存在",
	ERROR_STUDENT_COURSE_EXIST:     "选课记录已存在",

	ERROR_TEACHER_NOT_EXIST: "教师不存在",

	ERROR_STUDENT_NOT_EXIST: "学生不存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
