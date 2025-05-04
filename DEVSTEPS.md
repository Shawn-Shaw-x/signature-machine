## 第一步：搭建环境
- `go mod init` 一个空环境
- `git init` 托管到 `github`，新建 `.gitignore` 文件用于排除 `github` 提交文件
- 新建 `.env` 文件用于环境变量读取
- 新建 `bin/compile.sh` 用于管理 `protobuf` 的生成命令

## 第二步：搭建控制台应用框架
- 建立文件夹 `cmd/signature`,并新建 `main.go` 文件，这是程序的主入口
- 相同文件夹下建立 `cli.go`，使用 `urfave/cle/v2` 作为命令行应用的框架
- 新建 `Makefile` 文件用于管理 `shell` 命令(`build/test/lint/proto/clean`)
- 测试 `version` 命令是否正常运行(`./signature version`)

## 第三步：搭建 RPC 框架
- 编写 `proto` 文件，定义好接口、输入输出消息
- `proto` 生成 `go` 代码，放入 `protobuf` 包内
- 读取配置，包括 `Grpc` 的 `host`、`port` 和 `LevelDB` 的路径
- 打开 `LevelDB` 连接
- 启动 `GRPC` 服务，注册 `proto` 定义好的路由进去

## 第四步：实现接口
- 新建 `ssm` 包，负责封装 `ecdsa` 和 `eddsa` 的秘钥生成、消息签名

### handle.go 中实现接口
- 获取加密支持方式（`ecdsa`、`eddsa`）：
    返回是否支持该方式的地址生成和签名
- 批量导出公钥：
    根据请求的方式（`ecdsa`、`eddsa`）的不同，分别调用 `ecdsa` 和 `eddsa` 的秘钥对生成方法，
    生成出对应秘钥对。
    将私钥和公钥保存到 `leveldb` 中，将公钥返回给调用方
- 交易签名：
    请求方提供加密方式、`messageHash`、公钥，请求签名。签名机内部调用 `ecdsa` 或 `eddsa` 的
    签名方法将 `messageHash` 签名，返回 `signature`

