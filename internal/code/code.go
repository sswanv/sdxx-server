package code

import "github.com/dobyte/due/v2/codes"

var (
	OK               = codes.OK               // 成功
	Canceled         = codes.Canceled         // 已取消
	Unknown          = codes.Unknown          // 未知错误
	InvalidArgument  = codes.InvalidArgument  // 无效参数
	DeadlineExceeded = codes.DeadlineExceeded // 操作超时
	NotFound         = codes.NotFound         // 未找到资源
	InternalError    = codes.InternalError    // 服务器内部错误
	Unauthorized     = codes.Unauthorized     // 未授权
	IllegalInvoke    = codes.IllegalInvoke    // 非法调用
	IllegalRequest   = codes.IllegalRequest   // 非法请求
	TooManyRequests  = codes.TooManyRequests  // 请求过于频繁
)
