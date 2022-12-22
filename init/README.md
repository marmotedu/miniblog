# Systemd 配置、安装和启动

- [Systemd 配置、安装和启动](#systemd-配置安装和启动)
	- [1. 前置操作](#前置操作)
	- [2. 创建 miniblog systemd unit 模板文件](#创建-miniblog-systemd-unit-模板文件)
	- [3. 复制 systemd unit 模板文件到 sysmted 配置目录](#复制-systemd-unit-模板文件到-sysmted-配置目录)
	- [4. 启动 systemd 服务](#启动-systemd-服务)

## 1. 前置操作

1. 创建需要的目录 

```bash
sudo mkdir -p /data/miniblog /opt/miniblog/bin /etc/miniblog /var/log/miniblog
```

2. 编译构建 `miniblog` 二进制文件

```bash
make build # 编译源码生成 miniblog 二进制文件
```

3. 将 `miniblog` 可执行文件安装在 `bin` 目录下

```bash
sudo cp _output/platforms/linux/amd64/miniblog /opt/miniblog/bin # 安装二进制文件
```

4. 安装 `miniblog` 配置文件

```bash
sed 's/.\/_output/\/etc\/miniblog/g' configs/miniblog.yaml > miniblog.sed.yaml # 替换 CA 文件路径
sudo cp miniblog.sed.yaml /etc/miniblog/ # 安装配置文件
```

5. 安装 CA 文件

```bash
make ca # 创建 CA 文件
sudo cp -a _output/cert/ /etc/miniblog/ # 将 CA 文件复制到 miniblog 配置文件目录
```

## 2. 创建 miniblog systemd unit 模板文件

执行如下 shell 脚本生成 `miniblog.service.template`

```bash
cat > miniblog.service.template <<EOF
[Unit]
Description=APIServer for blog platform.
Documentation=https://github.com/marmotedu/miniblog/blob/master/init/README.md

[Service]
WorkingDirectory=/data/miniblog
ExecStartPre=/usr/bin/mkdir -p /data/miniblog
ExecStartPre=/usr/bin/mkdir -p /var/log/miniblog
ExecStart=/opt/miniblog/bin/miniblog --config=/etc/miniblog/miniblog.yaml
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
```

## 3. 复制 systemd unit 模板文件到 sysmted 配置目录

```bash
sudo cp miniblog.service.template /etc/systemd/system/miniblog.service
```

## 4. 启动 systemd 服务

```bash
sudo systemctl daemon-reload && systemctl enable miniblog && systemctl restart miniblog
```
