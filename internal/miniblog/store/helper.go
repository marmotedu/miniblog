// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package store

const defaultLimitValue = 20

// defaultLimit 设置默认查询记录数.
func defaultLimit(limit int) int {
	if limit == 0 {
		limit = defaultLimitValue
	}

	return limit
}
