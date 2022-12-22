## 用户相关操作

### 注册用户

执行以下 `curl` 命令注册一个新用户 `belm`:

```bash
$ curl -XPOST -H"Content-Type: application/json" -d'{"username":"belm","password":"miniblog1234","nickname":"belm","email":"nosbelm@qq.com","phone":"18188888xxx"}' http://127.0.0.1:8080/v1/users
null
```

输出参数说明：

- `username`: 用户名；
- `password`: 用户密码；
- `nickname`: 用户昵称；
- `email`: 用户邮箱地址；
- `phone`: 用户注册电话。

### 用户登录

执行以下 `curl` 命令登录 `belm` 用户:

```bash
$ curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"belm","password":"miniblog1234"}' http://127.0.0.1:8080/login
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYLVVzZXJuYW1lIjoiY29saW4iLCJleHAiOjIwMjg5MjcwNTIsImlhdCI6MTY2ODkyNzA1MiwibmJmIjoxNjY4OTI3MDUyfQ.F5fIj6GaSzAedmu5Wh_ja6Yk2qzi5XF9RauK511tC9A"}
```

登录用户后，后台会返回签发的 token，之后的请求都需要携带 token 进行请求认证，否则会返回以下错误：

```bash
{"code":20004,"message":"Token was invalid."}
```

为了方便使用 token，可以使用以下命令，将 token 保存在 Shell 变量中：

```bash
token=`curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"belm","password":"miniblog1234"}' http://127.0.0.1:8080/login | jq -r .token`
```

### 获取用户列表（仅限 root 用户）

执行以下 `curl` 命令获取用户列表:

```bash
$ curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users
{"totalCount":1,"users":[{"username":"belm","nickname":"belm","email":"nosbelm@qq.com","phone":"nosbelm@qq.com","postCount":0,"createdAt":"2022-11-20 14:19:01","updatedAt":"2022-11-20 14:19:01"}]}
```

### 获取用户详情

执行以下 `curl` 命令获取 `belm` 用户详情:

```bash
$ curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users/belm
{"username":"belm","nickname":"belm","email":"nosbelm@qq.com","phone":"18188888xxx","postCount":0,"createdAt":"2022-11-20 14:19:01","updatedAt":"2022-11-20 14:19:01"}
```

### 更新用户信息

执行以下 `curl` 命令更新 `belm` 用户信息:

```bash
$ curl -XPUT -H"Content-Type: application/json" -H"Authorization: Bearer $token" -d'{"nickname":"belm(modified)"}' http://127.0.0.1:8080/v1/users/belm
null
```


### 修改用户密码

执行以下 `curl` 命令修改用户密码:

```bash
$ curl -XPUT -H"Content-Type: application/json" -d'{"oldPassword":"miniblog1234","newPassword":"miniblog12345"}' http://127.0.0.1:8080/v1/users/belm/change-password
null
```

### 注销用户

执行以下 `curl` 命令删除 `belm` 用户:

```bash
$ curl -XDELETE -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users/belm
null
```
