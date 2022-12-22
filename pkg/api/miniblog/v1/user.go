// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package v1

// LoginRequest 指定了 `POST /login` 接口的请求参数.
type LoginRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

// LoginResponse 指定了 `POST /login` 接口的返回参数.
type LoginResponse struct {
	Token string `json:"token"`
}

// ChangePasswordRequest 指定了 `POST /v1/users/{name}/change-password` 接口的请求参数.
type ChangePasswordRequest struct {
	// 旧密码.
	OldPassword string `json:"oldPassword" valid:"required,stringlength(6|18)"`

	// 新密码.
	NewPassword string `json:"newPassword" valid:"required,stringlength(6|18)"`
}

// CreateUserRequest 指定了 `POST /v1/users` 接口的请求参数.
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,stringlength(11|11)"`
}
