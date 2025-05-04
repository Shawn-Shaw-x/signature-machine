## 第一步：搭建环境
- go mod init 一个空环境
- git init 托管到 github，新建 .gitignore 文件用于排除 github 提交文件
- 新建空的 .env 文件用于环境变量读取

## 第二步：搭建控制台应用框架
- 建立文件夹 cmd/signature,并新建 main.go 文件，这是程序的主入口
- 相同文件夹下建立 cli.go，使用 urfave/cle/v2 作为命令行应用的框架
- 新建 Makefile 文件用于管理 shell 命令
- 测试 version 命令是否正常运行

## 第三步：搭建 RPC 框架
- 编写 proto 文件，定义好接口、输入输出消息
- proto 生成 go 代码
- 读取配置，包括 Grpc 的 host、port 和 LevelDB 的路径
- 打开 LevelDB 连接
- 启动 GRPC 服务，注册 proto 定义好的路由进去

## 第四步：实现接口

