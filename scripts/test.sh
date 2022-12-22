#!/usr/bin/env bash

# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# Common utilities, variables and checks for all build scripts.
set -o errexit
set -o nounset
set -o pipefail

# The root of the build/dist directory
PROJ_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source ${PROJ_ROOT}/scripts/lib/logging.sh

INSECURE_SERVER="127.0.0.1:8080"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Retrieve
DCURL="curl -f -s -XDELETE" # Delete

# 注意：使用 root 用户登录系统，否则无法删除指定的用户
mb::test::login()
{
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/login \
    -d'{"username":"root","password":"miniblog1234"}' | grep -Po 'token[" :]+\K[^"]+'
}

mb::test::user()
{
  token="-HAuthorization: Bearer $(mb::test::login)"

  # 1. 如果有 colin、posttest 用户先清空
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/posttest > /dev/null

  # 2. 创建 colin 用户
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users \
    -d'{"username":"colin","password":"miniblog1234","nickname":"belm","email":"nosbelm@qq.com","phone":"1818888xxxx"}' > /dev/null

  # 3. 列出所有用户
  ${RCURL} "${token}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10" > /dev/null

  # 4. 获取 colin 用户的详细信息
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null

  # 5. 修改 colin 用户
  ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users/colin \
    -d'{"nickname":"colin","email":"colin_modified@foxmail.com","phone":"1812884xxxx"}' > /dev/null

  # 6. 删除 colin 用户
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null
  mb::log::info "$(echo -e '\033[32mcongratulations, /v1/users test passed!\033[0m')"
}

mb::test::post()
{

  # 1. 创建测试用户 posttest 用户
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users \
    -d'{"username":"posttest","password":"miniblog1234","nickname":"belm","email":"nosbelm@qq.com","phone":"1818888xxxx"}' > /dev/null

  # 2. posttest 登录 miniblog
  tokenStr=`${CCURL} "${Header}" http://${INSECURE_SERVER}/login -d'{"username":"posttest","password":"miniblog1234"}' | grep -Po 'token[" :]+\K[^"]+'`
  token="-HAuthorization: Bearer ${tokenStr}"

  # 3. 创建一个博客
  postID=`${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts -d'{"title":"installation","content":"installation."}' | grep -Po 'post-[a-z0-9]+'`

  # 4. 列出所有博客
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts > /dev/null

  # 5. 获取所创建博客的信息
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} > /dev/null

  # 6. 修改所创建博客的信息
  ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} -d'{"title":"modified"}' > /dev/null

  # 7. 删除所创建的博客
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} > /dev/null
  mb::log::info "$(echo -e '\033[32mcongratulations, /v1/posts test passed!\033[0m')"
}

# 测试 user 资源 CURD
mb::test::user

# 测试 post 资源 CURD
mb::test::post

mb::log::info "$(echo -e '\033[32mcongratulations, all test passed!\033[0m')"
