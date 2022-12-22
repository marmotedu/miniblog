// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/miniblog/internal/pkg/core"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	v1 "github.com/marmotedu/miniblog/pkg/api/miniblog/v1"
)

// ChangePassword 用来修改指定用户的密码.
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password function called")

	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().ChangePassword(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
