## miniblog 项目目录说明


```bash
├── api # Swagger / OpenAPI 文档存放目录
│   └── openapi
│       └── openapi.yaml # OpenAPI 3.0 API 接口文档
├── cmd # main 文件存放目录
│   └── miniblog
│       └── miniblog.go
├── configs # 配置文件存放目录
│   ├── miniblog.sql # 数据库初始化 SQL
│   ├── miniblog.yaml # miniblog 配置文件
│   └── nginx.conf # Nginx 配置
├── docs # 项目文档
│   ├── devel # 开发文档
│   │   ├── en-US # 英文文档
│   │   └── zh-CN # 中文文档
│   │       ├── architecture.md # miniblog 架构介绍
│   │       ├── conversions # 规范文档存放目录
│   │       │   ├── api.md # 接口规范
│   │       │   ├── commit.md # Commit 规范
│   │       │   ├── directory.md # 目录结构规范
│   │       │   ├── error_code.md # 错误码规范
│   │       │   ├── go_code.md # 代码规范
│   │       │   ├── log.md # 日志规范
│   │       │   └── version.md # 版本规范
│   │       └── README.md
│   ├── guide # 用户文档
│   │   ├── en-US # 英文文档
│   │   └── zh-CN # 中文文档
│   │       ├── announcements.md # 动态与公告
│   │       ├── best-practice # 最佳实践
│   │       ├── faq # 常见问题
│   │       ├── installation # 安装指南
│   │       ├── introduction # 产品介绍
│   │       ├── operation-guide # 操作指南
│   │       ├── quickstart # 快速入门
│   │       └── README.md
│   └── images # 项目图片存放目录
├── examples # 示例源码
├── go.mod
├── go.sum
├── init # Systemd Unit 文件保存目录
│   ├── miniblog.service # miniblog systemd unit
├── internal # 内部代码保存目录，这里面的代码不能被外部程序引用
│   ├── miniblog # miniblog 代码实现目录
│   │   ├── biz # biz 层代码
│   │   ├── controller # controller 层代码
│   │   │   └── v1 # API 接口版本
│   │   │       ├── post # 博客相关代码实现
│   │   │       │   ├── create.go # 创建博客
│   │   │       │   ├── delete_collection.go #批量删除博客
│   │   │       │   ├── delete.go # 删除博客
│   │   │       │   ├── get.go # 获取博客详情
│   │   │       │   ├── list.go # 获取博客列表
│   │   │       │   ├── post.go # 博客 Controller 结构定义、创建
│   │   │       │   └── update.go # 更新博客
│   │   │       └── user
│   │   │           ├── change_password.go # 修改用户密码
│   │   │           ├── create.go #创建用户
│   │   │           ├── delete.go # 删除用户
│   │   │           ├── get.go # 获取用户详情
│   │   │           ├── list.go # 获取用户列表
│   │   │           ├── login.go # 用户登录
│   │   │           ├── update.go  # 更新用户
│   │   │           └── user.go # 用户 Controller 结构定义、创建
│   │   ├── helper.go # 工具类代码存放文件
│   │   ├── miniblog.go # miniblog 主业务逻辑实现代码
│   │   ├── router.go # Gin 路由加载代码
│   │   └── store # store 层代码
│   └── pkg # 内部包保存目录
│       ├── core # core 包，用来保存一些核心的函数
│       ├── errno # errno 包，实现了 miniblog 的错误码功能
│       │   ├── code.go # 错误码定义文件
│       │   └── errno.go # errno 包功能函数文件
│       ├── known # 存放项目级的常量定义
│       ├── log # miniblog 自定义 log 包
│       ├── middleware # Gin 中间件包
│       │   ├── authn.go # 认证中间件
│       │   ├── authz.go # 授权中间件
│       │   ├── header.go # 指定 HTTP Response Header
│       │   └── requestid.go # 请求 / 返回头中添加 X-Request-ID
│       └── model # GORM Model
├── LICENSE # 声明代码所遵循的开源协议
├── Makefile # Makefile 文件，一般大型软件系统都是采用 make 来作为编译工具
├── _output # 临时文件存放目录
├── pkg # 可供外部程序直接使用的 Go 包存放目录
│   ├── api # REST API 接口定义存放目录
│   ├── proto # Protobuf 接口定义存放目录
│   ├── auth # auth 包，用来完成认证、授权功能
│   │   ├── authn.go # 认证功能
│   │   └── authz.go # 授权功能
│   ├── db # db 包，用来完成 MySQL 数据库连接
│   ├── token # JWT Token 的签发和解析
│   ├── util # 工具类包存放目录
│   │   └── id # id 包，用来生成唯一短 ID
│   └── version # version 包，用来保存 / 输出版本信息
├── README-en.md # 英文 README
├── README.md # 中文 README
├── scripts # 脚本文件
│   ├── boilerplate.txt # 指定版权头信息
│   ├── coverage.awk # awk 脚本，用来计算覆盖率
│   ├── make-rules # 子 Makefile 保存目录
│   │   ├── common.mk # 存放通用的 Makefile 变量
│   │   ├── generate.mk # 用来生成相关代码
│   │   ├── golang.mk # 用来编译源码
│   │   └── tools.mk # 用来完成工具的安装
│   └── wrktest.sh # wrk 性能测试脚本
└── third_party # 第三方 Go 包存放目录
```
