## 博客相关操作

**注意：**

- 假设你已经注册了 `belm` 用户。如果没有注册，可参考 [用户相关操作](./user.md) 进行注册；
- 确保 `belm` 用户的密码是 `miniblog1234`，否则需要修改下文中的 password。

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

### 创建博客

执行以下 `curl` 命令创建一个博客:

```bash
$ curl -XPOST -H"Content-Type: application/json" -H"Authorization: Bearer $token" -d'{"title":"miniblog installation guide","content":"The installation method is coming."}' http://127.0.0.1:8080/v1/posts
{"postID":"post-22vtll"}
```

输出参数说明：

- `postID`: 博客 ID，唯一代表一篇博客，由后端自动生成，系统唯一。

### 获取博客列表

执行以下 `curl` 命令获取博客列表:

```bash
$ curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/posts
{"totalCount":1,"posts":[{"username":"belm","postID":"post-22vtll","title":"miniblog installation guide","content":"The installation method is coming.","createdAt":"2022-11-20 15:32:58","updatedAt":"2022-11-20 15:32:58"}]}
```

### 获取博客详情

执行以下 `curl` 命令获取博客详情：

```bash
$ curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/posts/post-22vtll
{"username":"belm","postID":"post-22vtll","title":"miniblog installation guide","content":"The installation method is coming","createdAt":"2022-11-20 15:32:58","updatedAt":"2022-11-20 15:32:58"}
```

### 更新博客内容

执行以下 `curl` 命令更新博客内容：

```bash
$ curl -XPUT -H"Content-Type: application/json" -H"Authorization: Bearer $token" -d'{"content":"The installation method is still on the way."}' http://127.0.0.1:8080/v1/posts/post-22vtll
null
```

### 删除博客

执行以下 `curl` 命令删除博客：

```bash
$ curl -XDELETE -H"Authorization: Bearer $token" 'http://127.0.0.1:8080/v1/posts?postID=post-22vtll&postID=post-22yxq5'
null
```
