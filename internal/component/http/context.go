package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/dobyte/due/v2/codes"
	"github.com/gofiber/fiber/v3"
)

type Error struct {
	Code    int    `json:"code"`    // 响应码
	Message string `json:"message"` // 响应消息
}

type Context interface {
	fiber.Ctx
	// CTX 获取fiber.Ctx
	CTX() fiber.Ctx
	// Proxy 获取代理API
	Proxy() *Proxy
	// Failure 失败响应
	Failure(rst any) error
	// Success 成功响应
	Success(data ...any) error
	// StdRequest 获取标准请求（net/http）
	StdRequest() *http.Request
}

type context struct {
	fiber.Ctx
	proxy *Proxy
}

// CTX 获取fiber.Ctx
func (c *context) CTX() fiber.Ctx {
	return c.Ctx
}

// Proxy 代理API
func (c *context) Proxy() *Proxy {
	return c.proxy
}

// Failure 失败响应
func (c *context) Failure(rst any) error {
	var err Error
	switch v := rst.(type) {
	case error:
		code := codes.Convert(v)
		err = Error{
			Code:    code.Code(),
			Message: code.Message(),
		}
		return c.Status(StatusBadRequest).JSON(&Error{Code: code.Code(), Message: code.Message()})
	case *codes.Code:
		err = Error{
			Code:    v.Code(),
			Message: v.Message(),
		}
	default:
		err = Error{
			Code:    codes.Unknown.Code(),
			Message: codes.Unknown.Message(),
		}
	}
	return c.Status(StatusBadRequest).JSON(err)
}

// Success 成功响应
func (c *context) Success(data ...any) error {
	if len(data) > 0 {
		return c.JSON(data[0])
	} else {
		return c.JSON(struct{}{})
	}
}

// StdRequest 获取标准请求（net/http）
func (c *context) StdRequest() *http.Request {
	req := c.Request()

	std := &http.Request{}
	std.Method = c.Method()
	std.URL, _ = url.Parse(req.URI().String())
	std.Proto = c.Protocol()
	std.ProtoMajor, std.ProtoMinor, _ = http.ParseHTTPVersion(std.Proto)
	std.Header = c.GetReqHeaders()
	std.Host = c.Host()
	std.ContentLength = int64(len(c.Body()))
	std.RemoteAddr = c.RequestCtx().RemoteAddr().String()
	std.RequestURI = string(req.RequestURI())

	if req.Body() != nil {
		std.Body = io.NopCloser(bytes.NewReader(req.Body()))
	}

	return std
}
