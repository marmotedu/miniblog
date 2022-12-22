// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package id

import (
	"strings"

	shortid "github.com/jasonsoft/go-short-id"
)

// GenShortID 生成 6 位字符长度的唯一 ID.
func GenShortID() string {
	opt := shortid.Options{
		Number:        4,
		StartWithYear: true,
		EndWithHost:   false,
	}

	return strings.ToLower(shortid.Generate(opt))
}
