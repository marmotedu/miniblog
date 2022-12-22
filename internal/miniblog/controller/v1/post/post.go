// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package post

import (
	"github.com/marmotedu/miniblog/internal/miniblog/biz"
	"github.com/marmotedu/miniblog/internal/miniblog/store"
)

// PostController 是 post 模块在 Controller 层的实现，用来处理博客模块的请求.
type PostController struct {
	b biz.IBiz
}

// New 创建一个 post controller.
func New(ds store.IStore) *PostController {
	return &PostController{b: biz.NewBiz(ds)}
}
