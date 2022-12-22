// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marmotedu/miniblog/internal/pkg/errno"
)

// ErrResponse 定义了发生错误时的返回消息.
type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 err 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
