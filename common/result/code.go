// 状态码
package result

// 定义状态
type Codes struct {
	SUCCESS         uint
	FAILED          uint
	NOAUTH          uint
	AUTHFORMATERROR uint
	Message         map[uint]string
}

// APICode 状态码
var APICode = Codes{
	SUCCESS:         200,
	FAILED:          501,
	NOAUTH:          403,
	AUTHFORMATERROR: 405,
}

// 状态信息
func init() {
	APICode.Message = map[uint]string{
		APICode.SUCCESS:         "success",
		APICode.FAILED:          "failed",
		APICode.NOAUTH:          "请求头中auth为空",
		APICode.AUTHFORMATERROR: "请求头中auth格式错误",
	}
}

// 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
